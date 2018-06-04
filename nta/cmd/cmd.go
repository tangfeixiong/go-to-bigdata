package cmd

import (
	"flag"
	"path/filepath"
	"strconv"
	// "fmt"
	// "log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	// "google.golang.org/grpc"

	// "k8s.io/kubernetes/pkg/util/rand"

	"github.com/tangfeixiong/go-to-bigdata/nta/pkg/server"
	agentserver "github.com/tangfeixiong/go-to-bigdata/nta/pkg/server/agent/app"
	"github.com/tangfeixiong/go-to-bigdata/nta/pkg/server/config"
	"github.com/tangfeixiong/go-to-bigdata/pkg/util/homedir"
)

func RootCommandFor(name string, stopCh <-chan struct{}) *cobra.Command {
	var cfg config.Config
	// in, out, errout := os.Stdin, os.Stdout, os.Stderr
	collectorcfg := &service.Config{
		Common: &cfg,
	}
	agentcfg := &agent.Config{
		Common: &cfg,
	}

	root := &cobra.Command{
		Use:   name,
		Short: "Collector server for NTA with gRPC & ReST API",
		Long: `
        Collector server for NTA
        
        This is a ..., and acting as a client of Apache HBase.
        
        It is inspired by some github projects.
        `,
	}
	root.AddCommand(createCollectorCommand(collectorcfg, stopCh))
	root.AddCommand(createAgentCommand(agentcfg, stopCh))

	root.PersistentFlags().StringVar(&cfg.Kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file. it means running out of cluster if supplied")
	if home := homedir.HomeDir(); home != "" {
		root.PersistentFlags().Lookup("kubeconfig").NoOptDefVal = filepath.Join(home, ".kube", "config")
	}
	root.PersistentFlags().IntVar(&cfg.LogLevel, "log_level", 2, "for glog")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	command.Flags().StringVar(&cfg.HttpAddress, "http_addr", "0.0.0.0:10016", "IP:port format. Serve HTTP, or No HTTP if empty")
	command.Flags().StringVar(&cfg.GrpcAddress, "grpc_addr", "0.0.0.0:10017", "IP:port format")
	command.Flags().BoolVar(&cfg.SecureProtocol, "secure_protocol", false, "Currently not used, if both HTTP address and HTTPS flag not set, just gRPC noly")

	return root
}

func createCollectorCommand(cfg *server.Config, stopCh <-chan struct{}) *cobra.Command {

	command := &cobra.Command{
		Use:   "collect",
		Short: "Serving with gRPC and a gRPC-Gateway",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", strconv.Itoa(cfg.Common.LogLevel))
			flag.Parse()
			server.Start(cfg, stopCh)
		},
	}

	//	command.PersistentFlags().StringVar(&config.BaseDomain, "domain", "cluster.local", "Domain of K8s DNS")
	//	command.PersistentFlags().StringVar(&config.CustomResourceName, "custom_resource", "", "custom resource name")
	//	command.PersistentFlags().StringVar(&config.ClusterID, "hdfs_cluster_id", "", "HDFS cluster name, auto-creating by default")
	//	command.PersistentFlags().StringVar(&config.NodeType, "hdfs_node_type", "namenode", "Or datanode")
	//	command.PersistentFlags().StringVar(&config.Dir, "hadoop_dir", "/hadoop-3.0.0", "Directory of etc")
	//	command.PersistentFlags().IntVar(&config.Port, "service_port", 9000, "Service port")
	// command.Flags().AddGoFlagSet(flag.CommandLine)

	return command
}

func createAgentCommand(cfg *agent.Config, stopCh <-chan struct{}) *cobra.Command {

	command := &cobra.Command{
		Use:   "agent",
		Short: "Serving with gRPC",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", strconv.Itoa(cfg.LogLevel))
			flag.Parse()

			agentserver.Start(cfg, stopCh)
		},
	}

	command.Flags().StringVar(&cfg.Name, "name", "", "Unique agent name to identify a worker process")
	command.Flags().StringVar(&cfg.RemoteGrpc, "remote_grpc", "127.0.0.1:10017", "Kubernetes Service object name")
	command.Flags().BoolVar(&cfg.SecureRemote, "secure_transport", false, "Kubernetes namespace, or lookup value from env name POD_NAMESPACE, otherwise 'default'")
	return command
}
