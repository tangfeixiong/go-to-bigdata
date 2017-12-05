package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/firstrow/tcp_server"
	"github.com/pkg/errors"
)

func main() {
	ts := tcp_server.New("0.0.0.0:2017")

	ts.OnNewClient(func(c *tcp_server.Client) {
		// new client connected
		// lets send some message
		fmt.Println("Client accepted:", c.Conn().RemoteAddr().String())
		myc = &myClient{
			tc:   c,
			id:   1,
			salt: []byte("123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		}
		if err := myc.handshake(); err != nil {
			myc.tc.Conn().Close()
			fmt.Println(err)
		}
		fmt.Println("handshaked")
	})
	ts.OnNewMessage(func(c *tcp_server.Client, message string) {
		// new message received
		fmt.Println("message:", message)
		c.Send(message)
	})
	ts.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		// connection with client lost
		fmt.Println("closed, err: ", err)
	})

	ts.Listen()
}

type myClient struct {
	tc         *tcp_server.Client
	id         uint32
	salt       []byte
	capability uint32
	collation  uint8
	user       string
	dbname     string

	sequence uint8
}

var myc *myClient

func (myc *myClient) handshake() error {
	if err := myc.writeInitialHandshake(); err != nil {
		return errors.Wrap(err, "handshake")
	}
	if err := myc.readHandshakeResponse(); err != nil {
		return errors.Wrap(err, "handshake")
	}
	fmt.Printf("client: %+v\n", myc)

	data := make([]byte, 4, 32)
	data = append(data, OKHeader)
	data = append(data, 0, 0)
	if myc.capability&ClientProtocol41 > 0 {
		data = append(data, uint8(ServerStatusAutocommit), uint8(ServerStatusAutocommit>>8))
		data = append(data, 0, 0)
	}

	err := myc.writePacket(data)
	myc.sequence = 0
	if err != nil {
		return errors.Wrap(err, "handshake")
	}

	return nil
}

// Version informations.
var (
	MinProtocolVersion     byte   = 10
	MaxPayloadLen          int    = 1<<24 - 1
	ServerVersion          string = "5.5.31"
	ServerStatusAutocommit uint16 = 0x0002
	DefaultCharset                = "utf8"
	DefaultCollationID            = 33
	BinaryCollationID             = 63
	DefaultCollationName          = "utf8_general_ci"
)

// Header informations.
const (
	OKHeader          byte = 0x00
	ErrHeader         byte = 0xff
	EOFHeader         byte = 0xfe
	LocalInFileHeader byte = 0xfb
)

// Client informations.
const (
	ClientLongPassword uint32 = 1 << iota
	ClientFoundRows
	ClientLongFlag
	ClientConnectWithDB
	ClientNoSchema
	ClientCompress
	ClientODBC
	ClientLocalFiles
	ClientIgnoreSpace
	ClientProtocol41
	ClientInteractive
	ClientSSL
	ClientIgnoreSigpipe
	ClientTransactions
	ClientReserved
	ClientSecureConnection
	ClientMultiStatements
	ClientMultiResults
	ClientPSMultiResults
	ClientPluginAuth
	ClientConnectAtts
	ClientPluginAuthLenencClientData
)

func (myc *myClient) writeInitialHandshake() error {
	var defaultCapability = ClientLongPassword | ClientLongFlag | ClientConnectWithDB |
		ClientProtocol41 | ClientTransactions | ClientSecureConnection | ClientFoundRows

	data := make([]byte, 4, 128)

	// min version 10
	data = append(data, MinProtocolVersion)
	// server version[00]
	data = append(data, ServerVersion...)
	data = append(data, 0)
	// connection id
	data = append(data, byte(myc.id), byte(myc.id>>8), byte(myc.id>>16), byte(myc.id>>24))
	// auth-plugin-data-part-1
	data = append(data, myc.salt[0:8]...)
	// filler [00]
	data = append(data, 0)
	// capability flag lower 2 bytes, using default capability here
	data = append(data, byte(defaultCapability), byte(defaultCapability>>8))
	// charset, utf-8 default
	data = append(data, uint8(DefaultCollationID))
	//status
	data = append(data, uint8(ServerStatusAutocommit), uint8(ServerStatusAutocommit>>8))
	// below 13 byte may not be used
	// capability flag upper 2 bytes, using default capability here
	data = append(data, byte(defaultCapability>>16), byte(defaultCapability>>24))
	// filler [0x15], for wireshark dump, value is 0x15
	data = append(data, 0x15)
	// reserved 10 [00]
	data = append(data, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	// auth-plugin-data-part-2
	data = append(data, myc.salt[8:]...)
	// filler [00]
	data = append(data, 0)
	err := myc.writePacket(data)
	if err != nil {
		return errors.Wrap(err, "writePacket")
	}
	return nil
}

// writePacket writes data that already have header
func (myc *myClient) writePacket(data []byte) error {
	wb := bufio.NewWriterSize(myc.tc.Conn(), 16*1024)
	myc.sequence = 0

	length := len(data) - 4

	for length >= MaxPayloadLen {
		data[0] = 0xff
		data[1] = 0xff
		data[2] = 0xff

		data[3] = myc.sequence

		if n, err := wb.Write(data[:4+MaxPayloadLen]); err != nil {
			return errors.Wrap(err, "write")
		} else if n != (4 + MaxPayloadLen) {
			return errors.New("send")
		} else {
			myc.sequence++
			length -= MaxPayloadLen
			data = data[MaxPayloadLen:]
		}
	}

	data[0] = byte(length)
	data[1] = byte(length >> 8)
	data[2] = byte(length >> 16)
	data[3] = myc.sequence

	var n int
	var err error
	if n, err = wb.Write(data); err != nil {
		return errors.Wrap(err, "write")
	}
	if n != len(data) {
		return errors.New("send")
	}
	if err = wb.Flush(); err != nil {
		return errors.Wrap(err, "flush")
	}
	myc.sequence += 1
	fmt.Println("packets:", myc.sequence)
	return nil
}

func (myc *myClient) readHandshakeResponse() error {
	data, err := myc.readPacket()
	if err != nil {
		return errors.Wrap(err, "read packets")
	}

	pos := 0
	// capability
	myc.capability = binary.LittleEndian.Uint32(data[:4])
	pos += 4
	// skip max packet size
	pos += 4
	// charset, skip, if you want to use another charset, use set names
	myc.collation = data[pos]
	pos++
	// skip reserved 23[00]
	pos += 23
	// user name
	myc.user = string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
	pos += len(myc.user) + 1
	// auth length and auth
	authLen := int(data[pos])
	pos++
	auth := data[pos : pos+authLen]
	pos += authLen
	if myc.capability&ClientConnectWithDB > 0 {
		if len(data[pos:]) > 0 {
			idx := bytes.IndexByte(data[pos:], 0)
			myc.dbname = string(data[pos : pos+idx])
		}
	}

	fmt.Println("auth:", auth)
	// Open session and do auth
	/*
		cc.ctx, err = cc.server.driver.OpenCtx(uint64(cc.connectionID), cc.capability, uint8(cc.collation), cc.dbname)
		if err != nil {
			cc.Close()
			return errors.Trace(err)
		}
		if !cc.server.skipAuth() {
			// Do Auth
			addr := cc.conn.RemoteAddr().String()
			host, _, err1 := net.SplitHostPort(addr)
			if err1 != nil {
				return errors.Trace(mysql.NewErr(mysql.ErrAccessDenied, cc.user, addr, "Yes"))
			}
			user := fmt.Sprintf("%s@%s", cc.user, host)
			if !cc.ctx.Auth(user, auth, cc.salt) {
				return errors.Trace(mysql.NewErr(mysql.ErrAccessDenied, cc.user, host, "Yes"))
			}
		}
	*/
	return nil
}

func (myc *myClient) readPacket() ([]byte, error) {
	rb := bufio.NewReaderSize(myc.tc.Conn(), 16*1024)

	var header [4]byte

	if _, err := io.ReadFull(rb, header[:]); err != nil {
		return nil, errors.Wrap(err, "read header")
	}

	length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
	if length < 1 {
		return nil, fmt.Errorf("invalid payload length %d", length)
	}

	sequence := uint8(header[3])
	if sequence != myc.sequence {
		return nil, fmt.Errorf("invalid sequence %d != %d", sequence, myc.sequence)
	}

	myc.sequence++

	data := make([]byte, length)
	if _, err := io.ReadFull(rb, data); err != nil {
		return nil, errors.Wrap(err, "read centent")
	}
	if length < MaxPayloadLen {
		return data, nil
	}

	var buf []byte
	buf, err := myc.readPacket()
	if err != nil {
		return nil, err
	}
	return append(data, buf...), nil
}
