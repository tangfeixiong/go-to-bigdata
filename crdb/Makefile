
GOPATHP?=/Users/fanhongling/Downloads/workspace
GOPATHD?=/home/vagrant/go

IMG_NS?=docker.io/tangfeixiong
IMG_REPO?=crdb
IMG_TAG?=latest
GIT_COMMIT=$(shell date +%y%m%d%H%M)-git$(shell git rev-parse --short=7 HEAD)
DOCKER_FILE?=Dockerfile.busybox
REGISTRY_HOST?=172.17.4.50:5000

all: protoc-grpc docker-push

protoc-grpc:
	@protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATHD}/src/github.com/gogo/protobuf \
		-I${GOPATHD}/src/github.com/coreos/etcd/raft \
		-I${GOPATHD}/src \
		--gogo_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,Mgogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto,Mraftpb/raft.proto=github.com/coreos/etcd/raft/raftpb,Mpb/demo.proto=github.com/tangfeixiong/go-to-kubernetes/examples/raft/pb,plugins=grpc:. \
		pb/service.proto pb/demo.proto
	@protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATHD}/src/github.com/gogo/protobuf \
		-I${GOPATHD}/src/github.com/coreos/etcd/raft \
		-I${GOPATHD}/src \
		--grpc-gateway_out=logtostderr=true:. \
		pb/service.proto

go-install:
	@go install -v ./

docker-push: docker-build
	@docker push $(IMG_NS)/$(IMG_REPO):$(IMG_TAG)

.PHONY: all protoc-grpc go-install go-build docker-build docker-push docker-build-exporter docker-build-collector docker-cgo docker-run
