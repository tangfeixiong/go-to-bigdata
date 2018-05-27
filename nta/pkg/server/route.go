package server

import (
	"errors"
	"fmt"

	"github.com/golang/glog"

	"golang.org/x/net/context"

	"github.com/tangfeixiong/go-to-bigdata/nta/pb"
)

func (s *Server) CreateContact(ctx context.Context, req *pb.ContactReqResp) (*pb.ContactReqResp, error) {
	fmt.Println("Request to create Contact:", req)
	resp := &pb.ContactReqResp{
		Recipe: new(pb.Recipient),
	}
	if req == nil || req.Recipe == nil {
		glog.Error("Contact recipe is required")
		resp.StateCode = 100
		resp.StateMessage = "Contact recipe is required"
		return resp, errors.New(resp.StateMessage)
	}
	if req.Recipe.Plural == "" || req.Recipe.Group == "" {
		glog.Error("Empty Contact field is not allowed")
		resp.StateCode = 101
		resp.StateMessage = "Empty Contact field is not allowed"
		return resp, errors.New(resp.StateMessage)
	}

	resp.Recipe.Group = req.Recipe.Group
	resp.Recipe.Version = req.Recipe.Version
	resp.Recipe.Scope = req.Recipe.Scope
	resp.Recipe.Plural = req.Recipe.Plural
	resp.Recipe.Singular = req.Recipe.Singular
	resp.Recipe.Kind = req.Recipe.Kind
	//	err := s.ops["rabbitmq-operator"].CreateCRD(req.Recipe)
	//	if err != nil {
	//		glog.Infof("Create Contact failed: %s", err.Error())
	//		resp.StateCode = 10
	//		resp.StateMessage = err.Error()
	//		return resp, err
	//	}
	return resp, nil
}

func (s *Server) ReapContact(ctx context.Context, req *pb.ContactReqResp) (*pb.ContactReqResp, error) {
	fmt.Println("Requesting:", req)
	return new(pb.ContactReqResp), nil
}
