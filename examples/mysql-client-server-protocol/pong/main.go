package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:2017")
	if err != nil {
		panic("Error starting TCP server.")
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()

		// new client connected
		// lets send some message
		fmt.Println("Client accepted:", conn.RemoteAddr().String())
		mc := &mysqlConn{
			id:               1,
			netConn:          conn,
			cfg:              mysql.NewConfig(),
			maxAllowedPacket: maxPacketSize,
			maxWriteSize:     maxPacketSize - 1,
			sequence:         0x00,
			closech:          make(chan struct{}),
		}

		// Enable TCP Keepalives on TCP connections
		if tc, ok := mc.netConn.(*net.TCPConn); ok {
			if err := tc.SetKeepAlive(true); err != nil {
				// Don't send COM_QUIT before handshake.
				mc.netConn.Close()
				mc.netConn = nil
				//return nil, err
				panic(err)
			}
		}

		// Call startWatcher for context support (From Go 1.8)
		if s, ok := interface{}(mc).(watcher); ok {
			s.startWatcher()
		}

		mc.buf = newBuffer(mc.netConn)

		// Set I/O timeouts
		mc.buf.timeout = mc.cfg.ReadTimeout
		mc.writeTimeout = mc.cfg.WriteTimeout

		initbytes, err := mc.writeInitPacket()
		if err != nil {
			mc.netConn.Close()
			fmt.Println("Failed to write init packets:", err)
			continue
		}
		fmt.Printf("Send init packet: %+v\n", initbytes)

		if err := mc.readAuthPacket(); err != nil {
			mc.netConn.Close()
			fmt.Println("Failed to write auth response:", err)
			continue
		}
		fmt.Printf("handshaked: %+v\n", mc.cfg)

		for test := 0; test < 2; test++ {
			if err := mc.readCommandPacket(); err != nil {
				mc.netConn.Close()
				fmt.Println("Failed to write command response:", err)
			}
		}
		return
	}
}

type mysqlConn struct {
	id         uint32
	salt       []byte
	capability uint32
	collation  uint8
	user       string
	dbname     string

	// github.com/go-sql-driver/mysql/connection.go
	buf              buffer
	netConn          net.Conn
	cfg              *mysql.Config
	maxAllowedPacket int
	maxWriteSize     int
	writeTimeout     time.Duration
	flags            clientFlag
	sequence         uint8

	// for context support (Go 1.8+)
	watching bool
	watcher  chan<- mysqlContext
	closech  chan struct{}
	finished chan<- struct{}
	canceled atomicError // set non-nil if conn is canceled
	closed   atomicBool  // set when conn is closed, before closech is closed
}

/*
  https://github.com/go-sql-driver/mysql/blob/master/driver.go
*/
// watcher interface is used for context support (From Go 1.8)
type watcher interface {
	startWatcher()
}

/******************************************************************************
*                           Initialisation Process                            *
******************************************************************************/

// Handshake Initialization Packet
// http://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::Handshake
func (mc *mysqlConn) writeInitPacket() ([]byte, error) {
	data := make([]byte, 52)

	pos := 0
	// protocol version [1 byte]
	data[0] = uint8(minProtocolVersion)
	pos += 1

	// server version [null terminated string]
	data[1], data[2], data[3], data[4], data[5], data[6], data[7] = '5', '.', '5', '.', '3', '1', 0x00
	pos += 6 + 1

	// connection id [4 bytes]
	data[8], data[9], data[10], data[11] = uint8(1), uint8(0), uint8(0), uint8(0)
	pos += 4

	// first part of the password cipher [8 bytes]
	data[12], data[13], data[14], data[15], data[16], data[17], data[18], data[19] = 'c', 'i', 'p', 'h', 'e', 'r', ':', ' '
	pos += 8
	//cipher := data[pos : pos+8]

	// (filler) always 0x00 [1 byte]
	data[20] = 0x00
	pos += 1

	// capability flags (lower 2 bytes) [2 bytes]
	data[21], data[22] = uint8(clientLongPassword|clientFoundRows|clientLongFlag|clientConnectWithDB),
		uint8((clientProtocol41|clientTransactions|clientSecureConn)>>8)
	pos += 2

	if len(data) > pos {
		// character set [1 byte]
		// status flags [2 bytes]
		// capability flags (upper 2 bytes) [2 bytes]
		// length of auth-plugin-data [1 byte]
		// reserved (all [00]) [10 bytes]
		data[23] = uint8(collations[defaultCollation])
		data[24], data[25] = uint8(statusInAutocommit), uint8(statusInAutocommit>>8)
		data[26], data[27] = 0x00, 0x00
		data[28] = 0x15 // 8 + 12 + 1
		data[29], data[30], data[31], data[32], data[33], data[34], data[35], data[36], data[27], data[38] =
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00
		pos += 1 + 2 + 2 + 1 + 10

		// second part of the password cipher [mininum 13 bytes],
		// where len=MAX(13, length of auth-plugin-data - 8)
		//
		// The web documentation is ambiguous about the length. However,
		// according to mysql-5.7/sql/auth/sql_authentication.cc line 538,
		// the 13th byte is "\0 byte, terminating the second part of
		// a scramble". So the second part of the password cipher is
		// a NULL terminated string that's at least 13 bytes with the
		// last byte being NULL.
		//
		// The official Python library uses the fixed length 12
		// which seems to work but technically could have a hidden bug.
		data[39], data[40], data[41], data[42], data[43], data[44], data[45], data[46], data[47], data[48], data[49], data[50], data[51] =
			'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 0x00
		pos += 12 + 1
		//cipher = append(cipher, data[pos:pos+12]...)

		// TODO: Verify string termination
		// EOF if version (>= 5.5.7 and < 5.5.10) or (>= 5.6.0 and < 5.6.2)
		// \NUL otherwise
		//
		//if data[len(data)-1] == 0 {
		//	return
		//}
		//return ErrMalformPkt

		var b [52 + 4]byte
		copy(b[4:], data)
		err := mc.writePacket(b[:])
		return data, err
		// make a memory safe copy of the cipher slice
		//var b [20]byte
		//copy(b[:], cipher)
		//return b[:], nil
	}

	var b [23 + 4]byte
	copy(b[4:], data)
	err := mc.writePacket(b[:])
	return data, err
	// make a memory safe copy of the cipher slice
	//var b [8]byte
	//copy(b[:], cipher)
	//return b[:], nil
}

// Client Authentication Packet
// http://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::HandshakeResponse
func (mc *mysqlConn) readAuthPacket() error {
	data, err := mc.readPacket()
	if err != nil {
		return errors.Wrap(err, "read packet")
	}

	// ClientFlags [32 bit]
	mc.flags = clientFlag(binary.LittleEndian.Uint32(data[:4]))

	// MaxPacketSize [32 bit] (none)
	var mps uint32 = binary.LittleEndian.Uint32(data[4:8])
	fmt.Println("Max packet size:", mps)

	// Charset [1 byte]
	var found bool
	mc.collation = data[9]
	for k, v := range collations {
		if v == mc.collation {
			found = true
			mc.cfg.Collation = k
			break
		}
	}
	if !found {
		// Note possibility for false negatives:
		// could be triggered  although the collation is valid if the
		// collations map does not contain entries the server supports.
		fmt.Println("Client collation:", mc.cfg.Collation)
		//return errors.New("unknown collation")
	}

	// SSL Connection Request Packet
	// http://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::SSLRequest

	if len(data) > 4+4+1+23 {
		// Filler [23 bytes] (all 0x00)
		//fmt.Println(data[4+4+1+23:])

		// User [null terminated string]
		mc.cfg.User = string(data[4+4+1+23 : 4+4+1+23+bytes.IndexByte(data[4+4+1+23:], 0)])

		// ScrambleBuffer [length encoded integer]
		password_size := int(data[4+4+1+23+len(mc.cfg.User)+1])
		mc.cfg.Passwd = string(data[4+4+1+23+len(mc.cfg.User)+1+1 : 4+4+1+23+len(mc.cfg.User)+1+1+password_size])
		//data[pos] = byte(len(scrambleBuff))
		//pos += 1 + copy(data[pos+1:], scrambleBuff)

		// Databasename [null terminated string]
		pos := 4 + 4 + 1 + 23 + len(mc.cfg.User) + 1 + 1 + password_size
		if len(data) > 4+4+1+23+len(mc.cfg.User)+1+1+password_size {
			idx := bytes.IndexByte(data[pos:], 0)
			mc.cfg.DBName = string(data[pos : pos+idx])
			pos += idx + 1
			//pos += copy(data[pos:], mc.cfg.DBName)
			//data[pos] = 0x00
			//pos++
		}

		// Assume native client during response
		for {
			if len(data) <= pos {
				break
			}
			idx := bytes.IndexByte(data[pos:], 0)
			plugin := string(data[pos : pos+idx])
			pos += idx + 1
			fmt.Println("Plugin:", plugin)
			//pos += copy(data[pos:], "mysql_native_password")
			//data[pos] = 0x00

		}
	}

	// Send Auth response
	b := make([]byte, 11)
	b[4] = iOK
	b[5] = 0x00
	b[6] = 0x00
	b[7], b[8] = uint8(statusInAutocommit), uint8(statusInAutocommit>>8)
	b[9], b[10] = 0x00, 0x00
	return mc.writePacket(b)
}

func (mc *mysqlConn) readCommandPacket() error {
	mc.sequence = 0x00
	data, err := mc.readPacket()
	if err != nil {
		return errors.Wrap(err, "readPacket")
	}

	// Reset Packet Sequence

	// Add command byte
	command := data[0]

	var b []byte
	// Send CMD response
	switch command {
	case comPing:
		fmt.Println("command: ping")
		/*
		   OK
		     https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_basic_ok_packet.html
		*/
		// OK with CLIENT_PROTOCOL_41. 0 affected rows, last-insert-id was 0, AUTOCOMMIT enabled, 0 warnings. No further info.
		var affected_rows uint32 = 0x00000000
		var last_insert_id = 0x00000000
		status_flags := statusInAutocommit
		var warnings uint16 = 0x0000

		b = make([]byte, 11)
		b[4] = iOK
		// Length-Encoded Integer Type
		//   https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_basic_dt_integers.html#sect_protocol_basic_dt_int_le
		b[5] = uint8(affected_rows)
		b[6] = uint8(last_insert_id)
		// if CLIENT_PROTOCOL_41 is enabled
		b[7], b[8] = uint8(status_flags), uint8(status_flags>>8)
		b[9], b[10] = uint8(warnings), uint8(warnings>>8)
		//
	case comStmtPrepare:
		fmt.Println("command: stmtprepare")
		/*
		   ERR
		     https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_basic_err_packet.html
		*/
		var err_code uint16 = 0x0448
		var sql_state_marker byte = '#'
		sql_state := "HY000"
		error_message := "No tables used"

		b = make([]byte, 13, 64)
		b[4] = iERR
		b[5], b[6] = uint8(err_code), uint8(err_code>>8)
		// if CLIENT_PROTOCOL_41 is enabled
		b[7] = sql_state_marker
		b[8], b[9], b[10], b[11], b[12] = sql_state[0], sql_state[1], sql_state[2], sql_state[3], sql_state[4]
		//
		b = append(b, []byte(error_message)...)
	default:
		fmt.Println("command code:", command)
		/*
		   EOF
		     https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_basic_eof_packet.html
		*/
		// A MySQL 4.1 EOF packet with: 0 warnings, AUTOCOMMIT enabled.
		var warnings uint16 = 0x0000
		status_flags := statusInAutocommit

		b = make([]byte, 9)
		b[4] = iEOF
		// if CLIENT_PROTOCOL_41 is enabled
		b[5], b[6] = uint8(warnings), uint8(warnings>>8)
		b[7], b[8] = uint8(status_flags), uint8(status_flags>>8)
		//
	}
	return mc.writePacket(b)
}
