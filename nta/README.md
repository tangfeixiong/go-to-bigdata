# NTA Collector and Agent

Network Traffic Analysis

## Prerequisites

go
```
[vagrant@kubedev-172-17-4-59 ~]$ go version
go version go1.10 linux/amd64
```

```
[vagrant@kubedev-172-17-4-59 ~]$ go env GOBIN; go env GOPATH
/Users/fanhongling/Downloads/workspace/bin
/home/vagrant/go
```

protobuf
```
[vagrant@kubedev-172-17-4-59 ~]$ protoc --version
libprotoc 3.5.0
```

linux
```
[vagrant@kubedev-172-17-4-59 ~]$ uname -r
4.11.8-300.fc26.x86_64
```

```
[vagrant@kubedev-172-17-4-59 ~]$ cat /proc/version 
Linux version 4.11.8-300.fc26.x86_64 (mockbuild@bkernel02.phx2.fedoraproject.org) (gcc version 7.1.1 20170622 (Red Hat 7.1.1-3) (GCC) ) #1 SMP Thu Jun 29 20:09:48 UTC 2017
```

## Development

### Protocol buffer

protoc
```
[vagrant@kubedev-172-17-4-59 nta]$ GOPATH=/Users/fanhongling/Downloads/workspace:/home/vagrant/go make protoc-grpc
```

### Web UI

bind h5
```
[vagrant@kubedev-172-17-4-59 nta]$ GOPATH=/Users/fanhongling/Downloads/workspace:/home/vagrant/go make go-bindata-web
```

### Swagger UI

bind swagger h5
```
[vagrant@kubedev-172-17-4-59 nta]$ GOPATH=/Users/fanhongling/Downloads/workspace:/home/vagrant/go make go-bindata-swagger
#@pkg=artifact; src=template/...; output_file=pkg/spec/${pkg}/artifacts.go; \
#	go-bindata -nocompress -o ${output_file} -prefix ${PWD} -pkg ${pkg} ${src}
```

### Build 

With `go install`
```
[vagrant@kubedev-172-17-4-59 nta]$ GOPATH=/Users/fanhongling/Downloads/workspace:/home/vagrant/go make go-install
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/spf13/pflag
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/spf13/cobra
github.com/tangfeixiong/go-to-bigdata/nta/pkg/hbase
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/elazarl/go-bindata-assetfs
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/golang/glog
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/gorilla/websocket
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/golang/protobuf/proto
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/golang/protobuf/ptypes/struct
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/golang/protobuf/jsonpb
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/grpc-ecosystem/grpc-gateway/runtime/internal
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/grpc-ecosystem/grpc-gateway/utilities
github.com/tangfeixiong/go-to-bigdata/vendor/golang.org/x/net/context
github.com/tangfeixiong/go-to-bigdata/vendor/golang.org/x/net/http2/hpack
github.com/tangfeixiong/go-to-bigdata/vendor/golang.org/x/text/transform
github.com/tangfeixiong/go-to-bigdata/vendor/golang.org/x/text/unicode/bidi
github.com/tangfeixiong/go-to-bigdata/vendor/golang.org/x/text/secure/bidirule
github.com/tangfeixiong/go-to-bigdata/vendor/golang.org/x/text/unicode/norm
github.com/tangfeixiong/go-to-bigdata/vendor/golang.org/x/net/idna
github.com/tangfeixiong/go-to-bigdata/vendor/golang.org/x/net/lex/httplex
github.com/tangfeixiong/go-to-bigdata/vendor/golang.org/x/net/http2
github.com/tangfeixiong/go-to-bigdata/vendor/golang.org/x/net/internal/timeseries
github.com/tangfeixiong/go-to-bigdata/vendor/golang.org/x/net/trace
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/grpc/codes
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/grpc/credentials
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/grpc/grpclb/grpc_lb_v1
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/grpc/grpclog
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/grpc/internal
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/grpc/keepalive
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/grpc/metadata
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/grpc/naming
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/grpc/peer
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/grpc/stats
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/golang/protobuf/ptypes/any
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/genproto/googleapis/rpc/status
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/grpc/status
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/grpc/tap
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/grpc/transport
github.com/tangfeixiong/go-to-bigdata/vendor/google.golang.org/grpc
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/grpc-ecosystem/grpc-gateway/runtime
github.com/philips/grpc-gateway-example/pkg/ui/data/swagger
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/beorn7/perks/quantile
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/prometheus/client_model/go
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/matttproud/golang_protobuf_extensions/pbutil
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/prometheus/common/internal/bitbucket.org/ww/goautoneg
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/prometheus/common/model
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/prometheus/common/expfmt
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/prometheus/procfs/xfs
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/prometheus/procfs
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/prometheus/client_golang/prometheus
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/prometheus/client_golang/prometheus/promhttp
github.com/tangfeixiong/go-to-bigdata/pkg/util/httpfs
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/gogo/protobuf/proto
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/golang/protobuf/protoc-gen-go/descriptor
github.com/tangfeixiong/go-to-bigdata/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api
github.com/tangfeixiong/go-to-bigdata/nta/pb
github.com/tangfeixiong/go-to-bigdata/nta/pkg/ui/data/webapp
github.com/gogo/protobuf/proto
github.com/golang/protobuf/proto
github.com/golang/protobuf/ptypes/struct
github.com/golang/protobuf/jsonpb
github.com/grpc-ecosystem/grpc-gateway/runtime/internal
github.com/grpc-ecosystem/grpc-gateway/utilities
golang.org/x/net/context
golang.org/x/net/http2/hpack
golang.org/x/text/transform
golang.org/x/text/unicode/bidi
golang.org/x/text/secure/bidirule
golang.org/x/text/unicode/norm
golang.org/x/net/idna
golang.org/x/net/lex/httplex
golang.org/x/net/http2
golang.org/x/net/internal/timeseries
golang.org/x/net/trace
google.golang.org/grpc/codes
google.golang.org/grpc/credentials
google.golang.org/grpc/grpclb/grpc_lb_v1
google.golang.org/grpc/grpclog
google.golang.org/grpc/internal
google.golang.org/grpc/keepalive
google.golang.org/grpc/metadata
google.golang.org/grpc/naming
google.golang.org/grpc/peer
google.golang.org/grpc/stats
github.com/golang/protobuf/ptypes/any
google.golang.org/genproto/googleapis/rpc/status
google.golang.org/grpc/status
google.golang.org/grpc/tap
google.golang.org/grpc/transport
google.golang.org/grpc
github.com/grpc-ecosystem/grpc-gateway/runtime
github.com/golang/protobuf/protoc-gen-go/descriptor
github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api
github.com/tangfeixiong/go-to-bigdata/vendor/golang.org/x/net/websocket
github.com/tangfeixiong/go-to-bigdata/nta/pkg/server
github.com/tangfeixiong/go-to-bigdata/pkg/util/homedir
github.com/tangfeixiong/go-to-bigdata/nta/cmd
github.com/tangfeixiong/go-to-bigdata/nta
```