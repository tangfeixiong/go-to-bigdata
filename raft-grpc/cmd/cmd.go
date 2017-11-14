package cmd

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	// "google.golang.org/grpc"

	// "k8s.io/kubernetes/pkg/util/rand"

	"github.com/tangfeixiong/go-to-bigdata/raft-grpc/pkg/server"
)

func ip_v4(ei string) (net.IP, *net.IPNet, error) {
	var ip net.IP
	var ipnet *net.IPNet
	var perr error

	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		fmt.Println(inter.Name, inter.HardwareAddr)
		if addrs, err := inter.Addrs(); err == nil {
			for _, addr := range addrs {
				if inter.Flags&net.FlagLoopback != 0 {
					fmt.Println(inter.Name, "->", addr, "loopback")
					continue // loopback interface
				}
				if ok, err := regexp.MatchString(`\d+\.\d+\.\d+\.\d+[:/]\d+`, addr.String()); err == nil && ok {
					if ipv4Addr, ipv4Net, err := net.ParseCIDR(addr.String()); err == nil {
						fmt.Println(inter.Name, "->", ipv4Addr, ipv4Net, "ip-v4")
						if ipnet != nil {
							continue
						}
						if strings.Compare(strings.ToLower(ei), strings.ToLower(inter.Name)) == 0 {
							ip = ipv4Addr
							ipnet = ipv4Net
							continue
						}
					}
				} else {
					if ipnet == nil {
						perr = err
					}
				}
				fmt.Println(inter.Name, "->", addr)
			}
		} else {
			if ipnet == nil {
				perr = err
			}
		}
	}
	return ip, ipnet, perr
}

func RootCommandFor(name string) *cobra.Command {
	// in, out, errout := os.Stdin, os.Stdout, os.Stderr

	root := &cobra.Command{
		Use:   name,
		Short: "Replication under RAFT",
		Long: `
        bin
        
        TBD.
        `,
	}
	root.AddCommand(buildClusterCommand())
	root.AddCommand(createJoinClusterCommand())
	// root.AddCommand(createClientCommand())

	return root
}

func buildClusterCommand() *cobra.Command {
	var clustername, nodename string
	var peerIface, peerhost, grpchost, httphost string
	var datadriver string
	var loglevel string

	command := &cobra.Command{
		Use:   "build-cluster",
		Short: "New node of a RAFT cluster with gRPC",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", loglevel)
			flag.Parse()
			ipv4, ipv4net, err := ip_v4(peerIface)
			if ipv4net == nil {
				os.Exit(1)
			} else if err != nil {
				log.Fatal(err)
			}
			if nodename == "" {
				nodename = fmt.Sprintf("%s_%s", strings.Replace(ipv4.String(), ".", "-", -1), ipv4net.Mask.String())
			}
			server.NewServer(server.PeerHostOption(ipv4.String(), peerhost),
				server.GrpcHostOption(grpchost),
				server.HttpHostOption(httphost)).ClusterNamed(clustername).NodeNamed(nodename).DataDrived(datadriver).Run()
		},
	}

	command.Flags().StringVar(&clustername, "cluster-name", "cluster1", "for cluster name")
	command.Flags().StringVar(&nodename, "node-name", "", "for node name")
	command.Flags().StringVar(&peerIface, "peer-interface", "eth1", "for interface name")
	command.Flags().StringVar(&peerhost, "peer-host", ":12347", "for cluster peer to peer, [ipv4:]port")
	command.Flags().StringVar(&grpchost, "grpc-host", ":12345", "for client rpc, [ipv4:]port")
	command.Flags().StringVar(&httphost, "http-host", ":12346", "for client insecure http [ipv4:]port")
	command.Flags().StringVar(&datadriver, "storage", "", "for underlying store, e.g. rocksdb:/tmp/rocksdb-data")
	command.Flags().StringVar(&loglevel, "loglevel", "2", "for glog")
	// command.Flags().AddGoFlagSet(flag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	return command
}

func createJoinClusterCommand() *cobra.Command {
	var storage string
	var loglevel string

	command := &cobra.Command{
		Use:   "join-cluster",
		Short: "Create and join gRPC server into an existed RAFT cluster",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", loglevel)
			flag.Parse()
			server.RunCreation("", "", storage)
		},
	}

	command.Flags().StringVar(&storage, "storage", "", "for storage address, e.g. elasticsearch=http://localhost:9200")
	command.Flags().StringVar(&loglevel, "loglevel", "2", "for glog")
	// command.Flags().AddGoFlagSet(flag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	return command
}

//func createClientCommand() *cobra.Command {
//	command := &cobra.Command{
//		Use: "client",
//	}
//	return command
//}
