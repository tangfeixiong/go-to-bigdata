# Help

Go LevelDB
```
[vagrant@localhost go-to-bigdata]$ go run ./examples/syndtr0x2Fgoleveldb/main.go 
Data: [118 97 108 117 101]
index: 1 key: key0 value: these
index: 2 key: key1 value: are
index: 3 key: key2 value: test
index: 4 key: key3 value: data
index: 5 key: key4 value: !
Iter satus: <nil>
index: 5 key: key3 value: data
index: 5 key: key4 value: !
Iter satus: <nil>
```

MySQL client
```
[vagrant@localhost go-to-bigdata]$ go run ./examples/go-sql-driver0x2Fmysql/main.go
```

TCP server
```
[vagrant@localhost rap]$ go run ../../go-to-bigdata/examples/firstrow0x2Ftcp_server/main.go
```


MySQL Handshake
```
[vagrant@localhost rap]$ go run ../../go-to-bigdata/examples/mysql-client-server-protocol/handshake/main.go
```

MySQL ping
```
[vagrant@localhost go-to-bigdata]$ go run ./examples/go-sql-driver0x2Fmysql/main.go -mysql-server 172.17.4.50:2017
panic: Error 1096: No tables used

goroutine 1 [running]:
main.main()
	/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-bigdata/examples/go-sql-driver0x2Fmysql/main.go:45 +0x5f6
exit status 2
[vagrant@localhost go-to-bigdata]$ go run ./examples/go-sql-driver0x2Fmysql/main.go -mysql-server 172.17.4.50:2017
panic: malformed packet

goroutine 1 [running]:
main.main()
	/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-bigdata/examples/go-sql-driver0x2Fmysql/main.go:45 +0x5f6
exit status 2
```

MySQL pong
```
[vagrant@localhost mysql-client-server-protocol]$ go install ./pong
[vagrant@localhost mysql-client-server-protocol]$ pong
Client accepted: 172.17.4.50:49026
Send init packet: [10 53 46 53 46 51 49 0 1 0 0 0 99 105 112 104 101 114 58 32 0 15 162 33 2 0 0 0 21 0 0 0 0 0 0 0 0 0 0 97 98 99 100 101 102 103 104 105 106 107 108 0]
Max packet size: 0
Client collation: utf8_general_ci
Plugin: mysql_native_password
handshaked: &{User:root Passwd:?b
                                 ?<?1?8?5?"?$" Net: Addr: DBName:mysql Params:map[] Collation:utf8_general_ci Loc:UTC MaxAllowedPacket:4194304 TLSConfig: tls:<nil> Timeout:0s ReadTimeout:0s WriteTimeout:0s AllowAllFiles:false AllowCleartextPasswords:false AllowNativePasswords:true AllowOldPasswords:false ClientFoundRows:false ColumnsWithAlias:false InterpolateParams:false MultiStatements:false ParseTime:false RejectReadOnly:false}
command: ping
command code: 22
```