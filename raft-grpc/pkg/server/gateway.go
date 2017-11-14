package server

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/tangfeixiong/go-to-bigdata/raft-grpc/pb"
)

func (m *myCreation) Demo(ctx context.Context, req *pb.DemoReqResp) (*pb.DemoReqResp, error) {
	fmt.Printf("go to demo: %q\n", req)

	resp := new(pb.DemoReqResp)
	resp.StateCode = 2
	resp.StateMessage = "Not Implemented"

	return resp, nil
}
