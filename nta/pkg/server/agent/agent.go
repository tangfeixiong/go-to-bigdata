package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"k8s.io/kubernetes/pkg/util/rand"

	"github.com/tangfeixiong/go-to-bigdata/nta/pb"
	"github.com/tangfeixiong/go-to-bigdata/pkg/util/demotls"
)

type Config struct {
	SecureAddress   string
	InsecureAddress string
	SecureHTTP      bool

	RemoteGRPC   string
	SecureRemote bool

	LogLevel int
	HBase    hbase.Config
}

type Agent struct {
}

func (a *Agent) Shakehands() {
	var insecure bool
	pflag.BoolVar(&insecure, "insecure", true, "using http.")
	pflag.Parse()
	if insecure {
		fmt.Println("http...")
		url := "http://localhost:10001/v1/battlefields"

		var netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			// TLSHandshakeTimeout: 5 * time.Second,
		}
		var netClient = &http.Client{
			Timeout:   time.Second * 10,
			Transport: netTransport,
		}

		in, err := json.Marshal(&pbos.OpenstackNeutronNetRequestData{Name: "test"})
		if err != nil {
			panic(err)
		}

		response, err := netClient.Post(url, "application/json", bytes.NewBuffer(in))
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		fmt.Println("response Status:", response.Status)
		fmt.Println("response Headers:", response.Header)
		respbody, _ := ioutil.ReadAll(response.Body)
		fmt.Println("response Body:", string(respbody))

		var jsonStr = []byte(`{"name": "again"}`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("X-Custom-Header", "test again")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))

		return
	}

	var client pb.SimpleGRpcServiceClient

	//	println("grpc with tls")
	//	var opts []grpc.DialOption
	//	creds := credentials.NewClientTLSFromCert(demotls.DemoCertPool, "localhost:10000")
	//	opts = append(opts, grpc.WithTransportCredentials(creds))
	//	conn, err := grpc.Dial(demoAddr, opts...)
	//	if err != nil {
	//		grpclog.Fatalf("fail to dial: %v", err)
	//	}
	//	defer conn.Close()
	//	client = pb.NewSimpleGRpcServiceClient(conn)

	println("grpc")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client = pb.NewSimpleGRpcServiceClient(conn)

	// msg, err := client.Echo(context.Background(), &pb.EchoMessage{strings.Join(os.Args[2:], " ")})
	// println(msg.Value)
	println(strings.Join(os.Args[1:], " "))

	req := &pb.ContactReqResp{
		Recipe: &pb.Recipient{
			Group: "group",
		},
	}
	// copts := []grpc.CallOption{grpc.EmptyCallOption{}}
	copts := []grpc.CallOption{}
	resp, err := client.CreateContact(context.Background(), req, copts...)
	println(resp)

}
