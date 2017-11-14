
Test kvstore
```
[vagrant@localhost go-to-bigdata]$ go test -test.run kvstore ./raft-grpc/pkg/server/
ok  	github.com/tangfeixiong/go-to-bigdata/raft-grpc/pkg/server	0.011s
```

Test ProposeOnCommit 
```
[vagrant@localhost go-to-bigdata]$ go test -test.run ProposeOnCommit ./raft-grpc/pkg/server/
2017-11-13 22:59:14.876315 I | replaying WAL of member 3
2017-11-13 22:59:14.877044 I | replaying WAL of member 1
2017-11-13 22:59:14.877710 I | replaying WAL of member 2
2017-11-13 22:59:14.919139 I | raftexample: create wal error (sync .: invalid argument)
FAIL	github.com/tangfeixiong/go-to-bigdata/raft-grpc/pkg/server	0.435s
```

### Issue

grpc-gateway version
```
[vagrant@localhost go-to-bigdata]$ cp vendor/google.golang.org/genproto/googleapis/api/annotations/* vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api/
[vagrant@localhost go-to-bigdata]$ go test -test.run ProposeOnCommit -test.v ./raft-grpc/pkg/server/
# github.com/tangfeixiong/go-to-bigdata/raft-grpc/pb
raft-grpc/pb/service.pb.gw.go:86:39: not enough arguments in call to runtime.AnnotateContext
	have ("context".Context, *http.Request)
	want ("context".Context, *runtime.ServeMux, *http.Request)
raft-grpc/pb/service.pb.gw.go:88:21: not enough arguments in call to runtime.HTTPError
	have ("context".Context, runtime.Marshaler, http.ResponseWriter, *http.Request, error)
	want ("context".Context, *runtime.ServeMux, runtime.Marshaler, http.ResponseWriter, *http.Request, error)
raft-grpc/pb/service.pb.gw.go:93:21: not enough arguments in call to runtime.HTTPError
	have ("context".Context, runtime.Marshaler, http.ResponseWriter, *http.Request, error)
	want ("context".Context, *runtime.ServeMux, runtime.Marshaler, http.ResponseWriter, *http.Request, error)
raft-grpc/pb/service.pb.gw.go:97:36: not enough arguments in call to forward_RaftReplicaService_Demo_0
	have ("context".Context, runtime.Marshaler, http.ResponseWriter, *http.Request, "github.com/tangfeixiong/go-to-bigdata/vendor/github.com/golang/protobuf/proto".Message, []func("context".Context, http.ResponseWriter, "github.com/tangfeixiong/go-to-bigdata/vendor/github.com/golang/protobuf/proto".Message) error...)
	want ("context".Context, *runtime.ServeMux, runtime.Marshaler, http.ResponseWriter, *http.Request, "github.com/tangfeixiong/go-to-bigdata/vendor/github.com/golang/protobuf/proto".Message, ...func("context".Context, http.ResponseWriter, "github.com/tangfeixiong/go-to-bigdata/vendor/github.com/golang/protobuf/proto".Message) error)
FAIL	github.com/tangfeixiong/go-to-bigdata/raft-grpc/pkg/server [build failed]
```

Manually modify stub to solve
```
[vagrant@localhost go-to-bigdata]$ rm -rf /tmp/raftexample-*
[vagrant@localhost go-to-bigdata]$ go test -test.run ProposeOnCommit -test.v ./raft-grpc/pkg/server/
=== RUN   TestProposeOnCommit
2017-11-14 14:15:22.644505 I | replaying WAL of member 3
2017-11-14 14:15:22.645942 I | replaying WAL of member 1
2017-11-14 14:15:22.651912 I | replaying WAL of member 2
2017-11-14 14:15:22.684597 I | loading WAL at term 0 and index 0
2017-11-14 14:15:22.688049 I | loading WAL at term 0 and index 0
2017-11-14 14:15:22.692346 I | loading WAL at term 0 and index 0
raft2017/11/14 14:15:22 INFO: 1 became follower at term 0
raft2017/11/14 14:15:22 INFO: newRaft 1 [peers: [], term: 0, commit: 0, applied: 0, lastindex: 0, lastterm: 0]
raft2017/11/14 14:15:22 INFO: 1 became follower at term 1
2017-11-14 14:15:22.723675 I | rafthttp: starting peer 2...
2017-11-14 14:15:22.723763 I | rafthttp: started HTTP pipelining with peer 2
raft2017/11/14 14:15:22 INFO: 2 became follower at term 0
raft2017/11/14 14:15:22 INFO: newRaft 2 [peers: [], term: 0, commit: 0, applied: 0, lastindex: 0, lastterm: 0]
raft2017/11/14 14:15:22 INFO: 2 became follower at term 1
2017-11-14 14:15:22.724830 I | rafthttp: starting peer 1...
2017-11-14 14:15:22.725313 I | rafthttp: started HTTP pipelining with peer 1
2017-11-14 14:15:22.728586 I | rafthttp: started streaming with peer 1 (writer)
2017-11-14 14:15:22.730262 I | rafthttp: started streaming with peer 2 (writer)
2017-11-14 14:15:22.731361 I | rafthttp: started streaming with peer 2 (writer)
2017-11-14 14:15:22.731446 I | rafthttp: started streaming with peer 1 (writer)
2017-11-14 14:15:22.731471 I | rafthttp: started peer 2
2017-11-14 14:15:22.731502 I | rafthttp: added peer 2
2017-11-14 14:15:22.731627 I | rafthttp: started streaming with peer 2 (stream Message reader)
raft2017/11/14 14:15:22 INFO: 3 became follower at term 0
raft2017/11/14 14:15:22 INFO: newRaft 3 [peers: [], term: 0, commit: 0, applied: 0, lastindex: 0, lastterm: 0]
raft2017/11/14 14:15:22 INFO: 3 became follower at term 1
2017-11-14 14:15:22.732112 I | rafthttp: started peer 1
2017-11-14 14:15:22.732363 I | rafthttp: started streaming with peer 1 (stream Message reader)
2017-11-14 14:15:22.732612 I | rafthttp: started streaming with peer 1 (stream MsgApp v2 reader)
2017-11-14 14:15:22.732890 I | rafthttp: started streaming with peer 2 (stream MsgApp v2 reader)
2017-11-14 14:15:22.733104 I | rafthttp: starting peer 3...
2017-11-14 14:15:22.733193 I | rafthttp: starting peer 1...
2017-11-14 14:15:22.733215 I | rafthttp: started HTTP pipelining with peer 1
2017-11-14 14:15:22.736588 I | rafthttp: started peer 1
2017-11-14 14:15:22.738525 I | rafthttp: added peer 1
2017-11-14 14:15:22.738605 I | rafthttp: starting peer 2...
2017-11-14 14:15:22.738692 I | rafthttp: started HTTP pipelining with peer 2
2017-11-14 14:15:22.740200 I | rafthttp: started streaming with peer 1 (stream Message reader)
2017-11-14 14:15:22.740656 I | rafthttp: started HTTP pipelining with peer 3
2017-11-14 14:15:22.743893 I | rafthttp: started peer 3
2017-11-14 14:15:22.744030 I | rafthttp: added peer 3
2017-11-14 14:15:22.744108 I | rafthttp: started streaming with peer 1 (writer)
2017-11-14 14:15:22.744148 I | rafthttp: started streaming with peer 1 (stream MsgApp v2 reader)
2017-11-14 14:15:22.744692 I | rafthttp: started streaming with peer 3 (writer)
2017-11-14 14:15:22.744879 I | rafthttp: added peer 1
2017-11-14 14:15:22.746630 I | rafthttp: starting peer 3...
2017-11-14 14:15:22.746800 I | rafthttp: started HTTP pipelining with peer 3
2017-11-14 14:15:22.749621 I | rafthttp: started peer 3
2017-11-14 14:15:22.749724 I | rafthttp: added peer 3
2017-11-14 14:15:22.749914 I | rafthttp: started streaming with peer 2 (writer)
2017-11-14 14:15:22.750034 I | rafthttp: started peer 2
2017-11-14 14:15:22.750088 I | rafthttp: added peer 2
2017-11-14 14:15:22.750300 I | rafthttp: started streaming with peer 3 (writer)
2017-11-14 14:15:22.750477 I | rafthttp: started streaming with peer 3 (writer)
2017-11-14 14:15:22.750550 I | rafthttp: started streaming with peer 3 (stream MsgApp v2 reader)
2017-11-14 14:15:22.750844 I | rafthttp: started streaming with peer 3 (stream Message reader)
2017-11-14 14:15:22.752265 I | rafthttp: started streaming with peer 2 (stream Message reader)
2017-11-14 14:15:22.753252 I | rafthttp: started streaming with peer 2 (stream MsgApp v2 reader)
2017-11-14 14:15:22.753633 I | rafthttp: started streaming with peer 2 (writer)
2017-11-14 14:15:22.753707 I | rafthttp: started streaming with peer 1 (writer)
2017-11-14 14:15:22.753769 I | rafthttp: started streaming with peer 3 (writer)
2017-11-14 14:15:22.753841 I | rafthttp: started streaming with peer 3 (stream MsgApp v2 reader)
2017-11-14 14:15:22.754213 I | rafthttp: started streaming with peer 3 (stream Message reader)
2017-11-14 14:15:22.756099 I | rafthttp: peer 1 became active
2017-11-14 14:15:22.756162 I | rafthttp: established a TCP streaming connection with peer 1 (stream MsgApp v2 writer)
2017-11-14 14:15:22.756518 I | rafthttp: established a TCP streaming connection with peer 1 (stream Message writer)
2017-11-14 14:15:22.758505 I | rafthttp: peer 3 became active
2017-11-14 14:15:22.761426 I | rafthttp: established a TCP streaming connection with peer 3 (stream MsgApp v2 writer)
2017-11-14 14:15:22.761543 I | rafthttp: established a TCP streaming connection with peer 3 (stream Message writer)
2017-11-14 14:15:22.762480 I | rafthttp: established a TCP streaming connection with peer 1 (stream Message reader)
2017-11-14 14:15:22.763081 I | rafthttp: peer 2 became active
2017-11-14 14:15:22.763148 I | rafthttp: established a TCP streaming connection with peer 2 (stream Message writer)
2017-11-14 14:15:22.763320 I | rafthttp: established a TCP streaming connection with peer 1 (stream MsgApp v2 reader)
2017-11-14 14:15:22.763827 I | rafthttp: established a TCP streaming connection with peer 3 (stream Message reader)
2017-11-14 14:15:22.763918 I | rafthttp: established a TCP streaming connection with peer 3 (stream MsgApp v2 reader)
2017-11-14 14:15:22.764326 I | rafthttp: peer 3 became active
2017-11-14 14:15:22.765949 I | rafthttp: established a TCP streaming connection with peer 3 (stream Message writer)
2017-11-14 14:15:22.766014 I | rafthttp: established a TCP streaming connection with peer 3 (stream MsgApp v2 writer)
2017-11-14 14:15:22.766096 I | rafthttp: established a TCP streaming connection with peer 2 (stream MsgApp v2 reader)
2017-11-14 14:15:22.766660 I | rafthttp: established a TCP streaming connection with peer 2 (stream Message reader)
2017-11-14 14:15:22.769385 I | rafthttp: established a TCP streaming connection with peer 3 (stream Message reader)
2017-11-14 14:15:22.770036 I | rafthttp: established a TCP streaming connection with peer 3 (stream MsgApp v2 reader)
2017-11-14 14:15:22.770143 I | rafthttp: established a TCP streaming connection with peer 2 (stream MsgApp v2 writer)
2017-11-14 14:15:22.834139 I | rafthttp: peer 1 became active
2017-11-14 14:15:22.835077 I | rafthttp: peer 2 became active
2017-11-14 14:15:22.835174 I | rafthttp: established a TCP streaming connection with peer 2 (stream Message writer)
2017-11-14 14:15:22.835198 I | rafthttp: established a TCP streaming connection with peer 2 (stream Message reader)
2017-11-14 14:15:22.835239 I | rafthttp: established a TCP streaming connection with peer 1 (stream Message writer)
2017-11-14 14:15:22.835271 I | rafthttp: established a TCP streaming connection with peer 1 (stream Message reader)
2017-11-14 14:15:22.840869 I | rafthttp: established a TCP streaming connection with peer 2 (stream MsgApp v2 writer)
2017-11-14 14:15:22.841657 I | rafthttp: established a TCP streaming connection with peer 2 (stream MsgApp v2 reader)
2017-11-14 14:15:22.842629 I | rafthttp: established a TCP streaming connection with peer 1 (stream MsgApp v2 reader)
2017-11-14 14:15:22.842778 I | rafthttp: established a TCP streaming connection with peer 1 (stream MsgApp v2 writer)
raft2017/11/14 14:15:24 INFO: 1 is starting a new election at term 1
raft2017/11/14 14:15:24 INFO: 1 became candidate at term 2
raft2017/11/14 14:15:24 INFO: 1 received MsgVoteResp from 1 at term 2
raft2017/11/14 14:15:24 INFO: 1 [logterm: 1, index: 3] sent MsgVote request to 2 at term 2
raft2017/11/14 14:15:24 INFO: 1 [logterm: 1, index: 3] sent MsgVote request to 3 at term 2
raft2017/11/14 14:15:24 INFO: 3 [term: 1] received a MsgVote message with higher term from 1 [term: 2]
raft2017/11/14 14:15:24 INFO: 3 became follower at term 2
raft2017/11/14 14:15:24 INFO: 3 [logterm: 1, index: 3, vote: 0] cast MsgVote for 1 [logterm: 1, index: 3] at term 2
raft2017/11/14 14:15:24 INFO: 2 [term: 1] received a MsgVote message with higher term from 1 [term: 2]
raft2017/11/14 14:15:24 INFO: 2 became follower at term 2
raft2017/11/14 14:15:24 INFO: 2 [logterm: 1, index: 3, vote: 0] cast MsgVote for 1 [logterm: 1, index: 3] at term 2
raft2017/11/14 14:15:24 INFO: 1 received MsgVoteResp from 3 at term 2
raft2017/11/14 14:15:24 INFO: 1 [quorum:2] has received 2 MsgVoteResp votes and 0 vote rejections
raft2017/11/14 14:15:24 INFO: 1 became leader at term 2
raft2017/11/14 14:15:24 INFO: raft.node: 1 elected leader 1 at term 2
raft2017/11/14 14:15:24 INFO: raft.node: 3 elected leader 1 at term 2
raft2017/11/14 14:15:24 INFO: raft.node: 2 elected leader 1 at term 2
2017-11-14 14:15:24.168658 I | rafthttp: stopping peer 2...
2017-11-14 14:15:24.169904 I | rafthttp: closed the TCP streaming connection with peer 2 (stream MsgApp v2 writer)
2017-11-14 14:15:24.169957 I | rafthttp: stopped streaming with peer 2 (writer)
2017-11-14 14:15:24.170609 I | rafthttp: closed the TCP streaming connection with peer 2 (stream Message writer)
2017-11-14 14:15:24.170662 I | rafthttp: stopped streaming with peer 2 (writer)
2017-11-14 14:15:24.171268 I | rafthttp: stopped HTTP pipelining with peer 2
2017-11-14 14:15:24.171741 W | rafthttp: lost the TCP streaming connection with peer 2 (stream MsgApp v2 reader)
2017-11-14 14:15:24.171791 E | rafthttp: failed to read 2 on stream MsgApp v2 (context canceled)
2017-11-14 14:15:24.171809 I | rafthttp: peer 2 became inactive
2017-11-14 14:15:24.171830 I | rafthttp: stopped streaming with peer 2 (stream MsgApp v2 reader)
2017-11-14 14:15:24.172563 W | rafthttp: lost the TCP streaming connection with peer 2 (stream Message reader)
2017-11-14 14:15:24.172624 I | rafthttp: stopped streaming with peer 2 (stream Message reader)
2017-11-14 14:15:24.172648 I | rafthttp: stopped peer 2
2017-11-14 14:15:24.172667 I | rafthttp: stopping peer 3...
2017-11-14 14:15:24.173197 W | rafthttp: lost the TCP streaming connection with peer 1 (stream Message reader)
2017-11-14 14:15:24.173807 I | rafthttp: closed the TCP streaming connection with peer 3 (stream MsgApp v2 writer)
2017-11-14 14:15:24.173855 I | rafthttp: stopped streaming with peer 3 (writer)
2017-11-14 14:15:24.174159 I | rafthttp: closed the TCP streaming connection with peer 3 (stream Message writer)
2017-11-14 14:15:24.174159 I | rafthttp: stopped streaming with peer 3 (writer)
2017-11-14 14:15:24.175080 I | rafthttp: stopped HTTP pipelining with peer 3
2017-11-14 14:15:24.175536 W | rafthttp: lost the TCP streaming connection with peer 3 (stream MsgApp v2 reader)
2017-11-14 14:15:24.175808 E | rafthttp: failed to read 3 on stream MsgApp v2 (context canceled)
2017-11-14 14:15:24.176164 I | rafthttp: peer 3 became inactive
2017-11-14 14:15:24.176566 I | rafthttp: stopped streaming with peer 3 (stream MsgApp v2 reader)
2017-11-14 14:15:24.177033 W | rafthttp: lost the TCP streaming connection with peer 3 (stream Message reader)
2017-11-14 14:15:24.177215 I | rafthttp: stopped streaming with peer 3 (stream Message reader)
2017-11-14 14:15:24.177558 I | rafthttp: stopped peer 3
2017-11-14 14:15:24.187146 W | rafthttp: failed to process raft message (context canceled)
2017-11-14 14:15:24.187853 W | rafthttp: failed to process raft message (context canceled)
2017-11-14 14:15:24.193677 W | rafthttp: failed to process raft message (context canceled)
2017-11-14 14:15:24.193844 I | rafthttp: stopping peer 1...
2017-11-14 14:15:24.194219 I | rafthttp: closed the TCP streaming connection with peer 1 (stream MsgApp v2 writer)
2017-11-14 14:15:24.194271 I | rafthttp: stopped streaming with peer 1 (writer)
2017-11-14 14:15:24.194775 I | rafthttp: closed the TCP streaming connection with peer 1 (stream Message writer)
2017-11-14 14:15:24.194822 I | rafthttp: stopped streaming with peer 1 (writer)
2017-11-14 14:15:24.194844 W | rafthttp: lost the TCP streaming connection with peer 1 (stream Message reader)
2017-11-14 14:15:24.194879 W | rafthttp: lost the TCP streaming connection with peer 1 (stream MsgApp v2 reader)
2017-11-14 14:15:24.194900 W | rafthttp: lost the TCP streaming connection with peer 1 (stream MsgApp v2 reader)
2017-11-14 14:15:24.195394 I | rafthttp: stopped HTTP pipelining with peer 1
2017-11-14 14:15:24.195438 I | rafthttp: stopped streaming with peer 1 (stream MsgApp v2 reader)
2017-11-14 14:15:24.195467 E | rafthttp: failed to dial 1 on stream Message (context canceled)
2017-11-14 14:15:24.195483 I | rafthttp: peer 1 became inactive
2017-11-14 14:15:24.195497 I | rafthttp: stopped streaming with peer 1 (stream Message reader)
2017-11-14 14:15:24.195511 I | rafthttp: stopped peer 1
2017-11-14 14:15:24.195524 I | rafthttp: stopping peer 3...
2017-11-14 14:15:24.196029 I | rafthttp: closed the TCP streaming connection with peer 3 (stream MsgApp v2 writer)
2017-11-14 14:15:24.196086 I | rafthttp: stopped streaming with peer 3 (writer)
2017-11-14 14:15:24.196716 I | rafthttp: closed the TCP streaming connection with peer 3 (stream Message writer)
2017-11-14 14:15:24.196785 I | rafthttp: stopped streaming with peer 3 (writer)
2017-11-14 14:15:24.197192 I | rafthttp: stopped HTTP pipelining with peer 3
2017-11-14 14:15:24.197390 W | rafthttp: lost the TCP streaming connection with peer 3 (stream MsgApp v2 reader)
2017-11-14 14:15:24.197411 E | rafthttp: failed to read 3 on stream MsgApp v2 (context canceled)
2017-11-14 14:15:24.197424 I | rafthttp: peer 3 became inactive
2017-11-14 14:15:24.197895 I | rafthttp: stopped streaming with peer 3 (stream MsgApp v2 reader)
2017-11-14 14:15:24.198154 W | rafthttp: lost the TCP streaming connection with peer 3 (stream Message reader)
2017-11-14 14:15:24.198301 I | rafthttp: stopped streaming with peer 3 (stream Message reader)
2017-11-14 14:15:24.198441 I | rafthttp: stopped peer 3
2017-11-14 14:15:24.201701 W | rafthttp: lost the TCP streaming connection with peer 2 (stream MsgApp v2 reader)
2017-11-14 14:15:24.202274 E | rafthttp: failed to dial 1 on stream Message (dial tcp 127.0.0.1:10000: getsockopt: connection refused)
2017-11-14 14:15:24.202581 I | rafthttp: peer 1 became inactive
2017-11-14 14:15:24.203230 W | rafthttp: lost the TCP streaming connection with peer 2 (stream Message reader)
2017-11-14 14:15:24.204173 I | rafthttp: stopping peer 1...
2017-11-14 14:15:24.205647 I | rafthttp: closed the TCP streaming connection with peer 1 (stream MsgApp v2 writer)
2017-11-14 14:15:24.205894 I | rafthttp: stopped streaming with peer 1 (writer)
2017-11-14 14:15:24.206292 I | rafthttp: closed the TCP streaming connection with peer 1 (stream Message writer)
2017-11-14 14:15:24.206582 I | rafthttp: stopped streaming with peer 1 (writer)
2017-11-14 14:15:24.207012 I | rafthttp: stopped HTTP pipelining with peer 1
2017-11-14 14:15:24.207215 I | rafthttp: stopped streaming with peer 1 (stream MsgApp v2 reader)
2017-11-14 14:15:24.207619 I | rafthttp: stopped streaming with peer 1 (stream Message reader)
2017-11-14 14:15:24.207972 I | rafthttp: stopped peer 1
2017-11-14 14:15:24.208447 I | rafthttp: stopping peer 2...
2017-11-14 14:15:24.208712 E | rafthttp: failed to dial 2 on stream Message (dial tcp 127.0.0.1:10001: getsockopt: connection refused)
2017-11-14 14:15:24.209032 I | rafthttp: peer 2 became inactive
2017-11-14 14:15:24.209684 I | rafthttp: closed the TCP streaming connection with peer 2 (stream MsgApp v2 writer)
2017-11-14 14:15:24.209841 I | rafthttp: stopped streaming with peer 2 (writer)
2017-11-14 14:15:24.210597 I | rafthttp: closed the TCP streaming connection with peer 2 (stream Message writer)
2017-11-14 14:15:24.210761 I | rafthttp: stopped streaming with peer 2 (writer)
2017-11-14 14:15:24.211209 I | rafthttp: stopped HTTP pipelining with peer 2
2017-11-14 14:15:24.211771 I | rafthttp: stopped streaming with peer 2 (stream MsgApp v2 reader)
2017-11-14 14:15:24.212042 I | rafthttp: stopped streaming with peer 2 (stream Message reader)
2017-11-14 14:15:24.212316 I | rafthttp: stopped peer 2
--- PASS: TestProposeOnCommit (1.57s)
PASS
ok  	github.com/tangfeixiong/go-to-bigdata/raft-grpc/pkg/server	1.587s
```


## VirtualBox

Could not run in vboxsf
```
[vagrant@localhost go-to-bigdata]$ mkdir -p coreos0x2Fetcd0x2Fcontrib0x2Fraftexample
[vagrant@localhost go-to-bigdata]$ cp vendor/github.com/coreos/etcd/contrib/raftexample/*.go coreos0x2Fetcd0x2Fcontrib0x2Fraftexample/
[vagrant@localhost go-to-bigdata]$ go test -test.run ProposeOnCommit -test.v ./coreos0x2Fetcd0x2Fcontrib0x2Fraftexample/
=== RUN   TestProposeOnCommit
2017-11-14 11:29:01.026378 I | replaying WAL of member 1
2017-11-14 11:29:01.028227 I | replaying WAL of member 3
2017-11-14 11:29:01.030043 I | replaying WAL of member 2
2017-11-14 11:29:01.149912 I | raftexample: create wal error (sync .: invalid argument)
exit status 1
FAIL	github.com/tangfeixiong/go-to-bigdata/coreos0x2Fetcd0x2Fcontrib0x2Fraftexample	0.786s
```

Change waldir into _raft.go_
```
		waldir:      fmt.Sprintf("/tmp/raftexample-%d", id),
		snapdir:     fmt.Sprintf("/tmp/raftexample-%d-snap", id),
```

Succeeded
```
[vagrant@localhost go-to-bigdata]$ go test -test.run ProposeOnCommit -test.v ./coreos0x2Fetcd0x2Fcontrib0x2Fraftexample/
=== RUN   TestProposeOnCommit
2017-11-14 11:29:29.858170 I | replaying WAL of member 1
2017-11-14 11:29:29.862175 I | replaying WAL of member 2
2017-11-14 11:29:29.869373 I | replaying WAL of member 3
2017-11-14 11:29:29.887072 I | loading WAL at term 0 and index 0
raft2017/11/14 11:29:29 INFO: 1 became follower at term 0
raft2017/11/14 11:29:29 INFO: newRaft 1 [peers: [], term: 0, commit: 0, applied: 0, lastindex: 0, lastterm: 0]
raft2017/11/14 11:29:29 INFO: 1 became follower at term 1
2017-11-14 11:29:29.928201 I | rafthttp: starting peer 2...
2017-11-14 11:29:29.928531 I | rafthttp: started HTTP pipelining with peer 2
2017-11-14 11:29:29.935946 I | loading WAL at term 0 and index 0
2017-11-14 11:29:29.939425 I | rafthttp: started streaming with peer 2 (writer)
2017-11-14 11:29:29.939517 I | rafthttp: started streaming with peer 2 (writer)
2017-11-14 11:29:29.940705 I | loading WAL at term 0 and index 0
2017-11-14 11:29:29.956696 I | rafthttp: started peer 2
2017-11-14 11:29:29.957213 I | rafthttp: added peer 2
2017-11-14 11:29:29.959122 I | rafthttp: starting peer 3...
2017-11-14 11:29:29.959885 I | rafthttp: started streaming with peer 2 (stream MsgApp v2 reader)
2017-11-14 11:29:29.960753 I | rafthttp: started streaming with peer 2 (stream Message reader)
2017-11-14 11:29:29.961834 I | rafthttp: started HTTP pipelining with peer 3
2017-11-14 11:29:29.969094 I | rafthttp: started streaming with peer 3 (writer)
2017-11-14 11:29:29.970162 I | rafthttp: started streaming with peer 3 (writer)
2017-11-14 11:29:29.970265 I | rafthttp: started peer 3
2017-11-14 11:29:29.970321 I | rafthttp: added peer 3
2017-11-14 11:29:29.971268 I | rafthttp: started streaming with peer 3 (stream Message reader)
2017-11-14 11:29:29.973422 I | rafthttp: started streaming with peer 3 (stream MsgApp v2 reader)
raft2017/11/14 11:29:29 INFO: 2 became follower at term 0
raft2017/11/14 11:29:29 INFO: newRaft 2 [peers: [], term: 0, commit: 0, applied: 0, lastindex: 0, lastterm: 0]
raft2017/11/14 11:29:29 INFO: 2 became follower at term 1
2017-11-14 11:29:29.987640 I | rafthttp: starting peer 1...
2017-11-14 11:29:29.987682 I | rafthttp: started HTTP pipelining with peer 1
raft2017/11/14 11:29:29 INFO: 3 became follower at term 0
raft2017/11/14 11:29:29 INFO: newRaft 3 [peers: [], term: 0, commit: 0, applied: 0, lastindex: 0, lastterm: 0]
raft2017/11/14 11:29:29 INFO: 3 became follower at term 1
2017-11-14 11:29:29.988759 I | rafthttp: started streaming with peer 1 (writer)
2017-11-14 11:29:29.989385 I | rafthttp: started streaming with peer 1 (writer)
2017-11-14 11:29:29.989442 I | rafthttp: starting peer 1...
2017-11-14 11:29:29.989566 I | rafthttp: started HTTP pipelining with peer 1
2017-11-14 11:29:29.992714 I | rafthttp: started streaming with peer 1 (writer)
2017-11-14 11:29:29.995209 I | rafthttp: started peer 1
2017-11-14 11:29:29.995812 I | rafthttp: added peer 1
2017-11-14 11:29:29.995934 I | rafthttp: starting peer 3...
2017-11-14 11:29:29.996078 I | rafthttp: started HTTP pipelining with peer 3
2017-11-14 11:29:29.996505 I | rafthttp: started streaming with peer 1 (stream Message reader)
2017-11-14 11:29:29.997082 I | rafthttp: started streaming with peer 1 (stream MsgApp v2 reader)
2017-11-14 11:29:29.999205 I | rafthttp: started peer 1
2017-11-14 11:29:29.999843 I | rafthttp: added peer 1
2017-11-14 11:29:30.000103 I | rafthttp: starting peer 2...
2017-11-14 11:29:30.000562 I | rafthttp: started HTTP pipelining with peer 2
2017-11-14 11:29:30.003880 I | rafthttp: started peer 2
2017-11-14 11:29:30.003970 I | rafthttp: added peer 2
2017-11-14 11:29:30.004405 I | rafthttp: started peer 3
2017-11-14 11:29:30.004529 I | rafthttp: added peer 3
2017-11-14 11:29:30.004862 I | rafthttp: started streaming with peer 2 (writer)
2017-11-14 11:29:30.004969 I | rafthttp: started streaming with peer 2 (writer)
2017-11-14 11:29:30.005088 I | rafthttp: started streaming with peer 2 (stream MsgApp v2 reader)
2017-11-14 11:29:30.005774 I | rafthttp: started streaming with peer 2 (stream Message reader)
2017-11-14 11:29:30.006367 I | rafthttp: started streaming with peer 1 (stream Message reader)
2017-11-14 11:29:30.007281 I | rafthttp: peer 1 became active
2017-11-14 11:29:30.007457 I | rafthttp: established a TCP streaming connection with peer 1 (stream MsgApp v2 reader)
2017-11-14 11:29:30.008399 I | rafthttp: established a TCP streaming connection with peer 1 (stream Message reader)
2017-11-14 11:29:30.008500 I | rafthttp: started streaming with peer 1 (writer)
2017-11-14 11:29:30.008603 I | rafthttp: started streaming with peer 1 (stream MsgApp v2 reader)
2017-11-14 11:29:30.008989 I | rafthttp: started streaming with peer 3 (stream Message reader)
2017-11-14 11:29:30.009273 I | rafthttp: started streaming with peer 3 (writer)
2017-11-14 11:29:30.009375 I | rafthttp: started streaming with peer 3 (writer)
2017-11-14 11:29:30.009413 I | rafthttp: started streaming with peer 3 (stream MsgApp v2 reader)
2017-11-14 11:29:30.009831 I | rafthttp: peer 2 became active
2017-11-14 11:29:30.009928 I | rafthttp: established a TCP streaming connection with peer 2 (stream Message writer)
2017-11-14 11:29:30.010002 I | rafthttp: established a TCP streaming connection with peer 2 (stream MsgApp v2 writer)
2017-11-14 11:29:30.014028 I | rafthttp: peer 3 became active
2017-11-14 11:29:30.015618 I | rafthttp: established a TCP streaming connection with peer 3 (stream Message writer)
2017-11-14 11:29:30.015708 I | rafthttp: established a TCP streaming connection with peer 3 (stream MsgApp v2 writer)
2017-11-14 11:29:30.015816 I | rafthttp: established a TCP streaming connection with peer 3 (stream Message reader)
2017-11-14 11:29:30.015960 I | rafthttp: peer 2 became active
2017-11-14 11:29:30.016105 I | rafthttp: established a TCP streaming connection with peer 2 (stream MsgApp v2 reader)
2017-11-14 11:29:30.016743 I | rafthttp: established a TCP streaming connection with peer 2 (stream Message writer)
2017-11-14 11:29:30.018618 I | rafthttp: established a TCP streaming connection with peer 2 (stream Message reader)
2017-11-14 11:29:30.019133 I | rafthttp: peer 3 became active
2017-11-14 11:29:30.020146 I | rafthttp: established a TCP streaming connection with peer 3 (stream Message writer)
2017-11-14 11:29:30.020713 I | rafthttp: established a TCP streaming connection with peer 3 (stream MsgApp v2 writer)
2017-11-14 11:29:30.020804 I | rafthttp: established a TCP streaming connection with peer 3 (stream MsgApp v2 reader)
2017-11-14 11:29:30.020933 I | rafthttp: established a TCP streaming connection with peer 2 (stream MsgApp v2 writer)
2017-11-14 11:29:30.021015 I | rafthttp: peer 1 became active
2017-11-14 11:29:30.021416 I | rafthttp: established a TCP streaming connection with peer 1 (stream Message reader)
2017-11-14 11:29:30.021568 I | rafthttp: established a TCP streaming connection with peer 1 (stream MsgApp v2 reader)
2017-11-14 11:29:30.062564 I | rafthttp: established a TCP streaming connection with peer 1 (stream Message writer)
2017-11-14 11:29:30.063389 I | rafthttp: established a TCP streaming connection with peer 1 (stream MsgApp v2 writer)
2017-11-14 11:29:30.064085 I | rafthttp: established a TCP streaming connection with peer 2 (stream Message reader)
2017-11-14 11:29:30.064790 I | rafthttp: established a TCP streaming connection with peer 2 (stream MsgApp v2 reader)
2017-11-14 11:29:30.074993 I | rafthttp: established a TCP streaming connection with peer 1 (stream Message writer)
2017-11-14 11:29:30.075704 I | rafthttp: established a TCP streaming connection with peer 3 (stream Message reader)
2017-11-14 11:29:30.075771 I | rafthttp: established a TCP streaming connection with peer 1 (stream MsgApp v2 writer)
2017-11-14 11:29:30.075817 I | rafthttp: established a TCP streaming connection with peer 3 (stream MsgApp v2 reader)
raft2017/11/14 11:29:31 INFO: 2 is starting a new election at term 1
raft2017/11/14 11:29:31 INFO: 2 became candidate at term 2
raft2017/11/14 11:29:31 INFO: 2 received MsgVoteResp from 2 at term 2
raft2017/11/14 11:29:31 INFO: 2 [logterm: 1, index: 3] sent MsgVote request to 3 at term 2
raft2017/11/14 11:29:31 INFO: 2 [logterm: 1, index: 3] sent MsgVote request to 1 at term 2
raft2017/11/14 11:29:31 INFO: 3 [term: 1] received a MsgVote message with higher term from 2 [term: 2]
raft2017/11/14 11:29:31 INFO: 3 became follower at term 2
raft2017/11/14 11:29:31 INFO: 3 [logterm: 1, index: 3, vote: 0] cast MsgVote for 2 [logterm: 1, index: 3] at term 2
raft2017/11/14 11:29:31 INFO: 1 [term: 1] received a MsgVote message with higher term from 2 [term: 2]
raft2017/11/14 11:29:31 INFO: 1 became follower at term 2
raft2017/11/14 11:29:31 INFO: 1 [logterm: 1, index: 3, vote: 0] cast MsgVote for 2 [logterm: 1, index: 3] at term 2
raft2017/11/14 11:29:31 INFO: 2 received MsgVoteResp from 3 at term 2
raft2017/11/14 11:29:31 INFO: 2 [quorum:2] has received 2 MsgVoteResp votes and 0 vote rejections
raft2017/11/14 11:29:31 INFO: 2 became leader at term 2
raft2017/11/14 11:29:31 INFO: raft.node: 2 elected leader 2 at term 2
raft2017/11/14 11:29:31 INFO: raft.node: 3 elected leader 2 at term 2
raft2017/11/14 11:29:31 INFO: raft.node: 1 elected leader 2 at term 2
2017-11-14 11:29:31.027851 I | rafthttp: stopping peer 2...
2017-11-14 11:29:31.028171 I | rafthttp: closed the TCP streaming connection with peer 2 (stream MsgApp v2 writer)
2017-11-14 11:29:31.028233 I | rafthttp: stopped streaming with peer 2 (writer)
2017-11-14 11:29:31.028439 I | rafthttp: closed the TCP streaming connection with peer 2 (stream Message writer)
2017-11-14 11:29:31.028465 I | rafthttp: stopped streaming with peer 2 (writer)
2017-11-14 11:29:31.028829 I | rafthttp: stopped HTTP pipelining with peer 2
2017-11-14 11:29:31.028972 W | rafthttp: lost the TCP streaming connection with peer 2 (stream MsgApp v2 reader)
2017-11-14 11:29:31.029026 E | rafthttp: failed to read 2 on stream MsgApp v2 (context canceled)
2017-11-14 11:29:31.029049 I | rafthttp: peer 2 became inactive
2017-11-14 11:29:31.029082 I | rafthttp: stopped streaming with peer 2 (stream MsgApp v2 reader)
2017-11-14 11:29:31.029134 W | rafthttp: lost the TCP streaming connection with peer 2 (stream Message reader)
2017-11-14 11:29:31.029163 I | rafthttp: stopped streaming with peer 2 (stream Message reader)
2017-11-14 11:29:31.029188 I | rafthttp: stopped peer 2
2017-11-14 11:29:31.029209 I | rafthttp: stopping peer 3...
2017-11-14 11:29:31.029503 I | rafthttp: closed the TCP streaming connection with peer 3 (stream MsgApp v2 writer)
2017-11-14 11:29:31.029571 I | rafthttp: stopped streaming with peer 3 (writer)
2017-11-14 11:29:31.029766 I | rafthttp: closed the TCP streaming connection with peer 3 (stream Message writer)
2017-11-14 11:29:31.029887 I | rafthttp: stopped streaming with peer 3 (writer)
2017-11-14 11:29:31.030464 I | rafthttp: stopped HTTP pipelining with peer 3
2017-11-14 11:29:31.030908 W | rafthttp: lost the TCP streaming connection with peer 3 (stream MsgApp v2 reader)
2017-11-14 11:29:31.030963 I | rafthttp: stopped streaming with peer 3 (stream MsgApp v2 reader)
2017-11-14 11:29:31.031314 W | rafthttp: lost the TCP streaming connection with peer 3 (stream Message reader)
2017-11-14 11:29:31.031409 I | rafthttp: stopped streaming with peer 3 (stream Message reader)
2017-11-14 11:29:31.031428 I | rafthttp: stopped peer 3
2017-11-14 11:29:31.033004 W | rafthttp: lost the TCP streaming connection with peer 1 (stream MsgApp v2 reader)
2017-11-14 11:29:31.034177 W | rafthttp: lost the TCP streaming connection with peer 1 (stream Message reader)
2017-11-14 11:29:31.037049 W | rafthttp: lost the TCP streaming connection with peer 1 (stream Message reader)
2017-11-14 11:29:31.037481 W | rafthttp: lost the TCP streaming connection with peer 1 (stream MsgApp v2 reader)
2017-11-14 11:29:31.037981 E | rafthttp: failed to dial 1 on stream MsgApp v2 (read tcp 127.0.0.1:55434->127.0.0.1:10000: read: connection reset by peer)
2017-11-14 11:29:31.038714 I | rafthttp: peer 1 became inactive
2017-11-14 11:29:31.039162 I | rafthttp: stopping peer 1...
2017-11-14 11:29:31.040498 I | rafthttp: closed the TCP streaming connection with peer 1 (stream MsgApp v2 writer)
2017-11-14 11:29:31.040781 I | rafthttp: stopped streaming with peer 1 (writer)
2017-11-14 11:29:31.041497 I | rafthttp: closed the TCP streaming connection with peer 1 (stream Message writer)
2017-11-14 11:29:31.041756 I | rafthttp: stopped streaming with peer 1 (writer)
2017-11-14 11:29:31.042289 I | rafthttp: stopped HTTP pipelining with peer 1
2017-11-14 11:29:31.042607 I | rafthttp: stopped streaming with peer 1 (stream MsgApp v2 reader)
2017-11-14 11:29:31.042735 I | rafthttp: stopped streaming with peer 1 (stream Message reader)
2017-11-14 11:29:31.042821 I | rafthttp: stopped peer 1
2017-11-14 11:29:31.042910 I | rafthttp: stopping peer 3...
2017-11-14 11:29:31.043895 I | rafthttp: closed the TCP streaming connection with peer 3 (stream MsgApp v2 writer)
2017-11-14 11:29:31.044007 I | rafthttp: stopped streaming with peer 3 (writer)
2017-11-14 11:29:31.044584 I | rafthttp: closed the TCP streaming connection with peer 3 (stream Message writer)
2017-11-14 11:29:31.044683 I | rafthttp: stopped streaming with peer 3 (writer)
2017-11-14 11:29:31.044860 W | rafthttp: failed to process raft message (context canceled)
2017-11-14 11:29:31.045118 I | rafthttp: stopped HTTP pipelining with peer 3
2017-11-14 11:29:31.045266 W | rafthttp: lost the TCP streaming connection with peer 3 (stream MsgApp v2 reader)
2017-11-14 11:29:31.045452 I | rafthttp: stopped streaming with peer 3 (stream MsgApp v2 reader)
2017-11-14 11:29:31.045610 W | rafthttp: lost the TCP streaming connection with peer 3 (stream Message reader)
2017-11-14 11:29:31.045696 E | rafthttp: failed to read 3 on stream Message (context canceled)
2017-11-14 11:29:31.045774 I | rafthttp: peer 3 became inactive
2017-11-14 11:29:31.045855 I | rafthttp: stopped streaming with peer 3 (stream Message reader)
2017-11-14 11:29:31.045936 I | rafthttp: stopped peer 3
2017-11-14 11:29:31.047443 I | rafthttp: stopping peer 1...
2017-11-14 11:29:31.048007 I | rafthttp: closed the TCP streaming connection with peer 1 (stream MsgApp v2 writer)
2017-11-14 11:29:31.048187 I | rafthttp: stopped streaming with peer 1 (writer)
2017-11-14 11:29:31.048648 I | rafthttp: closed the TCP streaming connection with peer 1 (stream Message writer)
2017-11-14 11:29:31.048742 I | rafthttp: stopped streaming with peer 1 (writer)
2017-11-14 11:29:31.049141 I | rafthttp: stopped HTTP pipelining with peer 1
2017-11-14 11:29:31.049251 I | rafthttp: stopped streaming with peer 1 (stream MsgApp v2 reader)
2017-11-14 11:29:31.049436 I | rafthttp: stopped streaming with peer 1 (stream Message reader)
2017-11-14 11:29:31.049523 I | rafthttp: stopped peer 1
2017-11-14 11:29:31.049596 I | rafthttp: stopping peer 2...
2017-11-14 11:29:31.051126 I | rafthttp: closed the TCP streaming connection with peer 2 (stream MsgApp v2 writer)
2017-11-14 11:29:31.051476 I | rafthttp: stopped streaming with peer 2 (writer)
2017-11-14 11:29:31.051988 I | rafthttp: closed the TCP streaming connection with peer 2 (stream Message writer)
2017-11-14 11:29:31.052260 I | rafthttp: stopped streaming with peer 2 (writer)
2017-11-14 11:29:31.052823 I | rafthttp: stopped HTTP pipelining with peer 2
2017-11-14 11:29:31.056580 W | rafthttp: lost the TCP streaming connection with peer 2 (stream MsgApp v2 reader)
2017-11-14 11:29:31.056958 I | rafthttp: stopped streaming with peer 2 (stream MsgApp v2 reader)
2017-11-14 11:29:31.057553 W | rafthttp: lost the TCP streaming connection with peer 2 (stream Message reader)
2017-11-14 11:29:31.057920 I | rafthttp: stopped streaming with peer 2 (stream Message reader)
2017-11-14 11:29:31.058261 I | rafthttp: stopped peer 2
--- PASS: TestProposeOnCommit (1.21s)
PASS
ok  	github.com/tangfeixiong/go-to-bigdata/coreos0x2Fetcd0x2Fcontrib0x2Fraftexample	1.219s
```

data
```
[vagrant@localhost go-to-bigdata]$ ls /tmp/raft*
/tmp/raftexample-1:
0000000000000000-0000000000000000.wal

/tmp/raftexample-1-snap:

/tmp/raftexample-2:
0000000000000000-0000000000000000.wal

/tmp/raftexample-2-snap:

/tmp/raftexample-3:
0000000000000000-0000000000000000.wal

/tmp/raftexample-3-snap:
```

### coreos/etcd

project
```
[vagrant@localhost etcd]$ pwd
/data/src/github.com/coreos/etcd
```

git
```
[vagrant@localhost etcd]$ git log
commit eb19ab14e240657219f05df0dfddbe7382cd1472
Author: Gyu-Ho Lee <gyuhox@gmail.com>
Date:   Mon Nov 13 11:00:35 2017 -0800

    Merge pull request #8656 from gyuho/readme
    
    README: update badges
```

gopath
```
[vagrant@localhost etcd]$ export GOPATH=/data
```

vender
```
[vagrant@localhost etcd]$ ls vendor/
github.com  golang.org  google.golang.org
[vagrant@localhost etcd]$ ls vendor/github.com/
beorn7  coreos  golang  matttproud  prometheus  xiang90
[vagrant@localhost etcd]$ ls vendor/github.com/beorn7/
perks
[vagrant@localhost etcd]$ ls vendor/github.com/coreos/
go-semver  go-systemd  pkg
[vagrant@localhost etcd]$ ls vendor/github.com/golang/
protobuf
[vagrant@localhost etcd]$ ls vendor/github.com/matttproud/
golang_protobuf_extensions
[vagrant@localhost etcd]$ ls vendor/github.com/prometheus/
client_golang  client_model  common  procfs
[vagrant@localhost etcd]$ ls vendor/github.com/xiang90/
probing
[vagrant@localhost etcd]$ ls vendor/golang.org/x/
crypto  net  sys  time
[vagrant@localhost etcd]$ ls vendor/google.golang.org/
grpc
```

```
[vagrant@localhost etcd]$ go test -test.run ProposeOnCommit -test.v ./contrib/raftexample/
=== RUN   TestProposeOnCommit
2017-11-14 08:40:06.001011 I | replaying WAL of member 3
2017-11-14 08:40:06.001767 I | replaying WAL of member 2
2017-11-14 08:40:06.002132 I | replaying WAL of member 1
2017-11-14 08:40:06.042999 I | loading WAL at term 0 and index 0
2017-11-14 08:40:06.068245 I | loading WAL at term 0 and index 0
2017-11-14 08:40:06.069725 I | loading WAL at term 0 and index 0
raft2017/11/14 08:40:06 INFO: 1 became follower at term 0
raft2017/11/14 08:40:06 INFO: newRaft 1 [peers: [], term: 0, commit: 0, applied: 0, lastindex: 0, lastterm: 0]
raft2017/11/14 08:40:06 INFO: 1 became follower at term 1
2017-11-14 08:40:06.079152 I | rafthttp: starting peer 2...
2017-11-14 08:40:06.079656 I | rafthttp: started HTTP pipelining with peer 2
raft2017/11/14 08:40:06 INFO: 2 became follower at term 0
raft2017/11/14 08:40:06 INFO: newRaft 2 [peers: [], term: 0, commit: 0, applied: 0, lastindex: 0, lastterm: 0]
raft2017/11/14 08:40:06 INFO: 2 became follower at term 1
2017-11-14 08:40:06.080279 I | rafthttp: starting peer 1...
2017-11-14 08:40:06.080338 I | rafthttp: started HTTP pipelining with peer 1
raft2017/11/14 08:40:06 INFO: 3 became follower at term 0
raft2017/11/14 08:40:06 INFO: newRaft 3 [peers: [], term: 0, commit: 0, applied: 0, lastindex: 0, lastterm: 0]
raft2017/11/14 08:40:06 INFO: 3 became follower at term 1
2017-11-14 08:40:06.083591 I | rafthttp: started streaming with peer 1 (writer)
2017-11-14 08:40:06.084286 I | rafthttp: starting peer 1...
2017-11-14 08:40:06.084326 I | rafthttp: started HTTP pipelining with peer 1
2017-11-14 08:40:06.086852 I | rafthttp: started streaming with peer 1 (writer)
2017-11-14 08:40:06.086937 I | rafthttp: started peer 2
2017-11-14 08:40:06.086960 I | rafthttp: added peer 2
2017-11-14 08:40:06.086999 I | rafthttp: starting peer 3...
2017-11-14 08:40:06.087031 I | rafthttp: started HTTP pipelining with peer 3
2017-11-14 08:40:06.089583 I | rafthttp: started streaming with peer 2 (writer)
2017-11-14 08:40:06.089719 I | rafthttp: started streaming with peer 1 (writer)
2017-11-14 08:40:06.089844 I | rafthttp: started streaming with peer 2 (stream Message reader)
2017-11-14 08:40:06.091658 I | rafthttp: started streaming with peer 1 (writer)
2017-11-14 08:40:06.091865 I | rafthttp: started streaming with peer 2 (writer)
2017-11-14 08:40:06.091970 I | rafthttp: started streaming with peer 2 (stream MsgApp v2 reader)
2017-11-14 08:40:06.096309 I | rafthttp: started peer 1
2017-11-14 08:40:06.096420 I | rafthttp: added peer 1
2017-11-14 08:40:06.096455 I | rafthttp: starting peer 2...
2017-11-14 08:40:06.096516 I | rafthttp: started HTTP pipelining with peer 2
2017-11-14 08:40:06.097081 I | rafthttp: started streaming with peer 1 (stream MsgApp v2 reader)
2017-11-14 08:40:06.100534 I | rafthttp: started peer 2
2017-11-14 08:40:06.100614 I | rafthttp: added peer 2
2017-11-14 08:40:06.102246 I | rafthttp: started streaming with peer 1 (stream MsgApp v2 reader)
2017-11-14 08:40:06.102613 I | rafthttp: started streaming with peer 1 (stream Message reader)
2017-11-14 08:40:06.103413 I | rafthttp: started peer 1
2017-11-14 08:40:06.103543 I | rafthttp: added peer 1
2017-11-14 08:40:06.103583 I | rafthttp: starting peer 3...
2017-11-14 08:40:06.103657 I | rafthttp: started HTTP pipelining with peer 3
2017-11-14 08:40:06.105176 I | rafthttp: started streaming with peer 3 (writer)
2017-11-14 08:40:06.105293 I | rafthttp: started streaming with peer 2 (writer)
2017-11-14 08:40:06.105323 I | rafthttp: started streaming with peer 2 (writer)
2017-11-14 08:40:06.105561 I | rafthttp: started streaming with peer 2 (stream MsgApp v2 reader)
2017-11-14 08:40:06.105964 I | rafthttp: started streaming with peer 2 (stream Message reader)
2017-11-14 08:40:06.106279 I | rafthttp: started streaming with peer 3 (writer)
2017-11-14 08:40:06.107041 I | rafthttp: started streaming with peer 1 (stream Message reader)
2017-11-14 08:40:06.110402 I | rafthttp: started peer 3
2017-11-14 08:40:06.110495 I | rafthttp: added peer 3
2017-11-14 08:40:06.110971 I | rafthttp: started peer 3
2017-11-14 08:40:06.113624 I | rafthttp: added peer 3
2017-11-14 08:40:06.114582 I | rafthttp: started streaming with peer 3 (stream Message reader)
2017-11-14 08:40:06.115354 I | rafthttp: peer 2 became active
2017-11-14 08:40:06.115736 I | rafthttp: established a TCP streaming connection with peer 2 (stream Message writer)
2017-11-14 08:40:06.116110 I | rafthttp: started streaming with peer 3 (writer)
2017-11-14 08:40:06.116616 I | rafthttp: started streaming with peer 3 (stream MsgApp v2 reader)
2017-11-14 08:40:06.117195 I | rafthttp: started streaming with peer 3 (stream Message reader)
2017-11-14 08:40:06.117692 I | rafthttp: started streaming with peer 3 (writer)
2017-11-14 08:40:06.117957 I | rafthttp: started streaming with peer 3 (stream MsgApp v2 reader)
2017-11-14 08:40:06.118731 I | rafthttp: peer 1 became active
2017-11-14 08:40:06.119012 I | rafthttp: established a TCP streaming connection with peer 1 (stream Message reader)
2017-11-14 08:40:06.119778 I | rafthttp: peer 2 became active
2017-11-14 08:40:06.120062 I | rafthttp: established a TCP streaming connection with peer 2 (stream Message writer)
2017-11-14 08:40:06.120576 I | rafthttp: peer 3 became active
2017-11-14 08:40:06.120988 I | rafthttp: established a TCP streaming connection with peer 3 (stream Message reader)
2017-11-14 08:40:06.121920 I | rafthttp: peer 1 became active
2017-11-14 08:40:06.122368 I | rafthttp: established a TCP streaming connection with peer 1 (stream Message writer)
2017-11-14 08:40:06.122659 I | rafthttp: peer 3 became active
2017-11-14 08:40:06.122685 I | rafthttp: established a TCP streaming connection with peer 3 (stream MsgApp v2 reader)
2017-11-14 08:40:06.122817 I | rafthttp: established a TCP streaming connection with peer 2 (stream MsgApp v2 writer)
2017-11-14 08:40:06.123057 I | rafthttp: established a TCP streaming connection with peer 1 (stream MsgApp v2 writer)
2017-11-14 08:40:06.123218 I | rafthttp: established a TCP streaming connection with peer 3 (stream MsgApp v2 reader)
2017-11-14 08:40:06.123598 I | rafthttp: established a TCP streaming connection with peer 3 (stream Message reader)
2017-11-14 08:40:06.192025 I | rafthttp: established a TCP streaming connection with peer 1 (stream Message writer)
2017-11-14 08:40:06.192301 I | rafthttp: established a TCP streaming connection with peer 2 (stream Message reader)
2017-11-14 08:40:06.208845 I | rafthttp: established a TCP streaming connection with peer 3 (stream Message writer)
2017-11-14 08:40:06.217584 I | rafthttp: established a TCP streaming connection with peer 1 (stream MsgApp v2 writer)
2017-11-14 08:40:06.217684 I | rafthttp: established a TCP streaming connection with peer 3 (stream MsgApp v2 writer)
2017-11-14 08:40:06.217714 I | rafthttp: established a TCP streaming connection with peer 2 (stream MsgApp v2 writer)
2017-11-14 08:40:06.218157 I | rafthttp: established a TCP streaming connection with peer 3 (stream Message writer)
2017-11-14 08:40:06.218415 I | rafthttp: established a TCP streaming connection with peer 1 (stream MsgApp v2 reader)
2017-11-14 08:40:06.218557 I | rafthttp: established a TCP streaming connection with peer 1 (stream MsgApp v2 reader)
2017-11-14 08:40:06.218669 I | rafthttp: established a TCP streaming connection with peer 2 (stream MsgApp v2 reader)
2017-11-14 08:40:06.218826 I | rafthttp: established a TCP streaming connection with peer 1 (stream Message reader)
2017-11-14 08:40:06.219254 I | rafthttp: established a TCP streaming connection with peer 2 (stream Message reader)
2017-11-14 08:40:06.219563 I | rafthttp: established a TCP streaming connection with peer 3 (stream MsgApp v2 writer)
2017-11-14 08:40:06.219607 I | rafthttp: established a TCP streaming connection with peer 2 (stream MsgApp v2 reader)
raft2017/11/14 08:40:07 INFO: 2 is starting a new election at term 1
raft2017/11/14 08:40:07 INFO: 2 became candidate at term 2
raft2017/11/14 08:40:07 INFO: 2 received MsgVoteResp from 2 at term 2
raft2017/11/14 08:40:07 INFO: 2 [logterm: 1, index: 3] sent MsgVote request to 1 at term 2
raft2017/11/14 08:40:07 INFO: 2 [logterm: 1, index: 3] sent MsgVote request to 3 at term 2
raft2017/11/14 08:40:07 INFO: 1 [term: 1] received a MsgVote message with higher term from 2 [term: 2]
raft2017/11/14 08:40:07 INFO: 1 became follower at term 2
raft2017/11/14 08:40:07 INFO: 1 [logterm: 1, index: 3, vote: 0] cast MsgVote for 2 [logterm: 1, index: 3] at term 2
raft2017/11/14 08:40:07 INFO: 3 [term: 1] received a MsgVote message with higher term from 2 [term: 2]
raft2017/11/14 08:40:07 INFO: 3 became follower at term 2
raft2017/11/14 08:40:07 INFO: 3 [logterm: 1, index: 3, vote: 0] cast MsgVote for 2 [logterm: 1, index: 3] at term 2
raft2017/11/14 08:40:07 INFO: 2 received MsgVoteResp from 1 at term 2
raft2017/11/14 08:40:07 INFO: 2 [quorum:2] has received 2 MsgVoteResp votes and 0 vote rejections
raft2017/11/14 08:40:07 INFO: 2 became leader at term 2
raft2017/11/14 08:40:07 INFO: raft.node: 2 elected leader 2 at term 2
raft2017/11/14 08:40:07 INFO: raft.node: 3 elected leader 2 at term 2
raft2017/11/14 08:40:07 INFO: raft.node: 1 elected leader 2 at term 2
2017-11-14 08:40:07.494103 I | rafthttp: stopping peer 2...
2017-11-14 08:40:07.494464 I | rafthttp: closed the TCP streaming connection with peer 2 (stream MsgApp v2 writer)
2017-11-14 08:40:07.494511 I | rafthttp: stopped streaming with peer 2 (writer)
2017-11-14 08:40:07.494754 W | rafthttp: lost the TCP streaming connection with peer 1 (stream MsgApp v2 reader)
2017-11-14 08:40:07.495442 I | rafthttp: closed the TCP streaming connection with peer 2 (stream Message writer)
2017-11-14 08:40:07.495881 I | rafthttp: stopped streaming with peer 2 (writer)
2017-11-14 08:40:07.496494 W | rafthttp: lost the TCP streaming connection with peer 1 (stream Message reader)
2017-11-14 08:40:07.496569 I | rafthttp: stopped HTTP pipelining with peer 2
2017-11-14 08:40:07.496650 W | rafthttp: lost the TCP streaming connection with peer 2 (stream MsgApp v2 reader)
2017-11-14 08:40:07.496719 E | rafthttp: failed to read 2 on stream MsgApp v2 (context canceled)
2017-11-14 08:40:07.496748 I | rafthttp: peer 2 became inactive
2017-11-14 08:40:07.496774 I | rafthttp: stopped streaming with peer 2 (stream MsgApp v2 reader)
2017-11-14 08:40:07.496846 W | rafthttp: lost the TCP streaming connection with peer 2 (stream Message reader)
2017-11-14 08:40:07.496882 I | rafthttp: stopped streaming with peer 2 (stream Message reader)
2017-11-14 08:40:07.496908 I | rafthttp: stopped peer 2
2017-11-14 08:40:07.496948 I | rafthttp: stopping peer 3...
2017-11-14 08:40:07.497245 I | rafthttp: closed the TCP streaming connection with peer 3 (stream MsgApp v2 writer)
2017-11-14 08:40:07.497315 I | rafthttp: stopped streaming with peer 3 (writer)
2017-11-14 08:40:07.497456 I | rafthttp: closed the TCP streaming connection with peer 3 (stream Message writer)
2017-11-14 08:40:07.498090 I | rafthttp: stopped streaming with peer 3 (writer)
2017-11-14 08:40:07.498864 I | rafthttp: stopped HTTP pipelining with peer 3
2017-11-14 08:40:07.499015 W | rafthttp: lost the TCP streaming connection with peer 3 (stream MsgApp v2 reader)
2017-11-14 08:40:07.499035 E | rafthttp: failed to read 3 on stream MsgApp v2 (context canceled)
2017-11-14 08:40:07.499049 I | rafthttp: peer 3 became inactive
2017-11-14 08:40:07.499071 I | rafthttp: stopped streaming with peer 3 (stream MsgApp v2 reader)
2017-11-14 08:40:07.499111 W | rafthttp: lost the TCP streaming connection with peer 3 (stream Message reader)
2017-11-14 08:40:07.499130 I | rafthttp: stopped streaming with peer 3 (stream Message reader)
2017-11-14 08:40:07.499143 I | rafthttp: stopped peer 3
2017-11-14 08:40:07.499302 E | rafthttp: failed to find member 2 in cluster 1000
2017-11-14 08:40:07.499468 W | rafthttp: lost the TCP streaming connection with peer 1 (stream Message reader)
2017-11-14 08:40:07.499556 W | rafthttp: lost the TCP streaming connection with peer 1 (stream MsgApp v2 reader)
2017-11-14 08:40:07.499739 I | rafthttp: stopping peer 1...
2017-11-14 08:40:07.499769 E | rafthttp: failed to find member 2 in cluster 1000
2017-11-14 08:40:07.499889 I | rafthttp: closed the TCP streaming connection with peer 1 (stream MsgApp v2 writer)
2017-11-14 08:40:07.499904 I | rafthttp: stopped streaming with peer 1 (writer)
2017-11-14 08:40:07.499958 E | rafthttp: failed to dial 1 on stream MsgApp v2 (peer 1 failed to find local node 2)
2017-11-14 08:40:07.499985 I | rafthttp: peer 1 became inactive
2017-11-14 08:40:07.500026 I | rafthttp: closed the TCP streaming connection with peer 1 (stream Message writer)
2017-11-14 08:40:07.500041 I | rafthttp: stopped streaming with peer 1 (writer)
2017-11-14 08:40:07.500207 I | rafthttp: stopped HTTP pipelining with peer 1
2017-11-14 08:40:07.500234 I | rafthttp: stopped streaming with peer 1 (stream MsgApp v2 reader)
2017-11-14 08:40:07.500254 I | rafthttp: stopped streaming with peer 1 (stream Message reader)
2017-11-14 08:40:07.500268 I | rafthttp: stopped peer 1
2017-11-14 08:40:07.500281 I | rafthttp: stopping peer 3...
2017-11-14 08:40:07.500434 I | rafthttp: closed the TCP streaming connection with peer 3 (stream MsgApp v2 writer)
2017-11-14 08:40:07.500451 I | rafthttp: stopped streaming with peer 3 (writer)
2017-11-14 08:40:07.500884 I | rafthttp: closed the TCP streaming connection with peer 3 (stream Message writer)
2017-11-14 08:40:07.500946 I | rafthttp: stopped streaming with peer 3 (writer)
2017-11-14 08:40:07.501261 I | rafthttp: stopped HTTP pipelining with peer 3
2017-11-14 08:40:07.501457 W | rafthttp: lost the TCP streaming connection with peer 2 (stream MsgApp v2 reader)
2017-11-14 08:40:07.501577 W | rafthttp: lost the TCP streaming connection with peer 3 (stream MsgApp v2 reader)
2017-11-14 08:40:07.501854 E | rafthttp: failed to read 3 on stream MsgApp v2 (context canceled)
2017-11-14 08:40:07.501878 I | rafthttp: peer 3 became inactive
2017-11-14 08:40:07.501902 I | rafthttp: stopped streaming with peer 3 (stream MsgApp v2 reader)
2017-11-14 08:40:07.502685 E | rafthttp: failed to dial 1 on stream MsgApp v2 (dial tcp 127.0.0.1:10000: getsockopt: connection refused)
2017-11-14 08:40:07.503048 I | rafthttp: peer 1 became inactive
2017-11-14 08:40:07.503745 W | rafthttp: lost the TCP streaming connection with peer 2 (stream Message reader)
2017-11-14 08:40:07.503831 W | rafthttp: lost the TCP streaming connection with peer 3 (stream Message reader)
2017-11-14 08:40:07.503885 I | rafthttp: stopped streaming with peer 3 (stream Message reader)
2017-11-14 08:40:07.503911 I | rafthttp: stopped peer 3
2017-11-14 08:40:07.504528 E | rafthttp: failed to find member 3 in cluster 1000
2017-11-14 08:40:07.505178 E | rafthttp: failed to dial 2 on stream Message (dial tcp 127.0.0.1:10001: getsockopt: connection reset by peer)
2017-11-14 08:40:07.505367 I | rafthttp: peer 2 became inactive
2017-11-14 08:40:07.506034 I | rafthttp: stopping peer 1...
2017-11-14 08:40:07.506559 I | rafthttp: closed the TCP streaming connection with peer 1 (stream MsgApp v2 writer)
2017-11-14 08:40:07.506909 I | rafthttp: stopped streaming with peer 1 (writer)
2017-11-14 08:40:07.507537 I | rafthttp: closed the TCP streaming connection with peer 1 (stream Message writer)
2017-11-14 08:40:07.507731 I | rafthttp: stopped streaming with peer 1 (writer)
2017-11-14 08:40:07.508159 I | rafthttp: stopped HTTP pipelining with peer 1
2017-11-14 08:40:07.508480 I | rafthttp: stopped streaming with peer 1 (stream MsgApp v2 reader)
2017-11-14 08:40:07.508839 I | rafthttp: stopped streaming with peer 1 (stream Message reader)
2017-11-14 08:40:07.509116 I | rafthttp: stopped peer 1
2017-11-14 08:40:07.509673 I | rafthttp: stopping peer 2...
2017-11-14 08:40:07.510990 I | rafthttp: closed the TCP streaming connection with peer 2 (stream MsgApp v2 writer)
2017-11-14 08:40:07.511471 I | rafthttp: stopped streaming with peer 2 (writer)
2017-11-14 08:40:07.513223 I | rafthttp: closed the TCP streaming connection with peer 2 (stream Message writer)
2017-11-14 08:40:07.513474 I | rafthttp: stopped streaming with peer 2 (writer)
2017-11-14 08:40:07.513982 I | rafthttp: stopped HTTP pipelining with peer 2
2017-11-14 08:40:07.514136 I | rafthttp: stopped streaming with peer 2 (stream MsgApp v2 reader)
2017-11-14 08:40:07.514251 I | rafthttp: stopped streaming with peer 2 (stream Message reader)
2017-11-14 08:40:07.514348 I | rafthttp: stopped peer 2
--- PASS: TestProposeOnCommit (1.51s)
PASS
ok  	github.com/coreos/etcd/contrib/raftexample	1.523s
```