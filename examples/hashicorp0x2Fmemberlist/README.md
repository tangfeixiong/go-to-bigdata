

Test
```
[vagrant@localhost go-to-bigdata]$ go test -test.v -test.run DefaultLANConfig_protocolVersion github.com/hashicorp/memberlist/
=== RUN   TestDefaultLANConfig_protocolVersion
--- PASS: TestDefaultLANConfig_protocolVersion (0.00s)
PASS
ok  	github.com/hashicorp/memberlist	0.012s
[vagrant@localhost go-to-bigdata]$ go test -test.v -test.run Create_protocolVersion github.com/hashicorp/memberlist/
=== RUN   TestCreate_protocolVersion
--- PASS: TestCreate_protocolVersion (0.01s)
PASS
ok  	github.com/hashicorp/memberlist	0.014s
[vagrant@localhost go-to-bigdata]$ go test -test.v -test.run Create github.com/hashicorp/memberlist/
=== RUN   TestCreate_protocolVersion
--- PASS: TestCreate_protocolVersion (0.00s)
=== RUN   TestCreate_secretKey
--- PASS: TestCreate_secretKey (0.00s)
=== RUN   TestCreate_secretKeyEmpty
--- PASS: TestCreate_secretKeyEmpty (0.00s)
=== RUN   TestCreate_keyringOnly
--- PASS: TestCreate_keyringOnly (0.00s)
=== RUN   TestCreate_keyringAndSecretKey
--- PASS: TestCreate_keyringAndSecretKey (0.00s)
=== RUN   TestCreate_invalidLoggerSettings
--- PASS: TestCreate_invalidLoggerSettings (0.00s)
=== RUN   TestCreate
--- PASS: TestCreate (0.01s)
=== RUN   TestMemberList_CreateShutdown
2017/11/15 23:01:38 [DEBUG] memberlist: Using dynamic bind port 42080
--- PASS: TestMemberList_CreateShutdown (0.00s)
PASS
ok  	github.com/hashicorp/memberlist	0.038s
```