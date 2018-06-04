package options

import (
	//"k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig"
	"k8s.io/kubernetes/pkg/kubelet/config"

	"github.com/tangfeixiong/go-to-bigdata/nta/pkg/server/agent"
)

// Refer to https://github.com/kubernetes/kubernetes/blob/master/cmd/kubelet/app/options/options.go#48

// A configuration field should go in KubeletFlags instead of KubeletConfiguration if any of these are true:
// - its value will never, or cannot safely be changed during the lifetime of a node
// - its value cannot be safely shared between nodes at the same time (e.g. a hostname)
//   KubeletConfiguration is intended to be shared between nodes
// In general, please try to avoid adding flags or configuration fields,
// we already have a confusingly large amount of them.
//type KubeletFlags struct {
type AgentFlags struct {
	KubeConfig          string
	BootstrapKubeconfig string

	// Insert a probability of random errors during calls to the master.
	ChaosChance float64
	// Crash immediately, rather than eating panics.
	ReallyCrashForTesting bool

	// TODO(mtaufen): It is increasingly looking like nobody actually uses the
	//                Kubelet's runonce mode anymore, so it may be a candidate
	//                for deprecation and removal.
	// If runOnce is true, the Kubelet will check the API server once for pods,
	// run those in addition to the pods specified by static pod files, and exit.
	RunOnce bool

	// HostnameOverride is the hostname used to identify the kubelet instead
	// of the actual hostname.
	HostnameOverride string
	// NodeIP is IP address of the node.
	// If set, kubelet will use this IP address for the node.
	NodeIP string

	// Container-runtime-specific options.
	config.ContainerRuntimeOptions

	// certDirectory is the directory where the TLS certs are located (by
	// default /var/run/kubernetes). If tlsCertFile and tlsPrivateKeyFile
	// are provided, this flag will be ignored.
	CertDirectory string

	// cloudProvider is the provider for cloud services.
	// +optional
	CloudProvider string

	// cAdvisorPort is the port of the localhost cAdvisor endpoint (set to 0 to disable)
	CAdvisorPort int32

	// WindowsService should be set to true if kubelet is running as a service on Windows
	// Its corresponding flag only gets registered in Windows builds
	WindowsService bool

	// containerized should be set to true if kubelet is running in a container.
	Containerized bool
	// remoteRuntimeEndpoint is the endpoint of remote runtime service
	RemoteRuntimeEndpoint string

	// lockFilePath is the path that kubelet will use to as a lock file.
	// It uses this file as a lock to synchronize with other kubelet processes
	// that may be running.
	LockFilePath string
	// ExitOnLockContention is a flag that signifies to the kubelet that it is running
	// in "bootstrap" mode. This requires that 'LockFilePath' has been set.
	// This will cause the kubelet to listen to inotify events on the lock file,
	// releasing it and exiting when another process tries to open that file.
	ExitOnLockContention bool
}

//func ValidateKubeletFlags(f *KubeletFlags) error {
//	// ensure that nobody sets DynamicConfigDir if the dynamic config feature gate is turned off
//	if f.DynamicConfigDir.Provided() && !utilfeature.DefaultFeatureGate.Enabled(features.DynamicKubeletConfig) {
//		return fmt.Errorf("the DynamicKubeletConfig feature gate must be enabled in order to use the --dynamic-config-dir flag")
//	}
//	if f.CAdvisorPort != 0 && utilvalidation.IsValidPortNum(int(f.CAdvisorPort)) != nil {
//		return fmt.Errorf("invalid configuration: CAdvisorPort (--cadvisor-port) %v must be between 0 and 65535, inclusive", f.CAdvisorPort)
//	}
//	if f.NodeStatusMaxImages < -1 {
//		return fmt.Errorf("invalid configuration: NodeStatusMaxImages (--node-status-max-images) must be -1 or greater")
//	}
//	return nil
//}

// Refer to https://github.com/kubernetes/kubernetes/blob/master/cmd/kubelet/app/options/options.go#299

// KubeletServer encapsulates all of the parameters necessary for starting up
// a kubelet. These can either be set via command line or directly.
//type KubeletServer struct {
//	KubeletFlags
//	kubeletconfig.KubeletConfiguration
//}
type AgentServer struct {
	AgentFlags
	agent.Config
}

// validateKubeletServer validates configuration of KubeletServer and returns an error if the input configuration is invalid
//func ValidateKubeletServer(s *KubeletServer) error {
//	// please add any KubeletConfiguration validation to the kubeletconfigvalidation.ValidateKubeletConfiguration function
//	if err := kubeletconfigvalidation.ValidateKubeletConfiguration(&s.KubeletConfiguration); err != nil {
//		return err
//	}
//	if err := ValidateKubeletFlags(&s.KubeletFlags); err != nil {
//		return err
//	}
//	return nil
//}
