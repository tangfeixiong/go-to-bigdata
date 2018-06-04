package agent

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/golang/glog"
	cadvisorapi "github.com/google/cadvisor/info/v1"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	clientset "k8s.io/client-go/kubernetes"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/kubernetes/pkg/kubelet/cadvisor"
	kubecontainer "k8s.io/kubernetes/pkg/kubelet/container"
	"k8s.io/kubernetes/pkg/scheduler/algorithm/predicates"
	utilipt "k8s.io/kubernetes/pkg/util/iptables"
	nodeutil "k8s.io/kubernetes/pkg/util/node"

	"github.com/tangfeixiong/go-to-bigdata/nta/pkg/server/agent/app"
	"github.com/tangfeixiong/go-to-bigdata/nta/pkg/server/config"
)

type Config struct {
	Name            string
	RemoteGrpc      string
	RemoteHttp      string
	SecureTransport bool

	Common *config.Config
}

func Start(cfg *Config, stopCh <-chan struct{}) {
	a := &Agent{
		Config:          cfg,
		nodeIPValidator: validateNodeIP,
	}

	// klet.setNodeStatusFuncs = klet.defaultNodeStatusFuncs()
	a.setNodeStatusFuncs = a.defaultNodeStatusFuncs()

	a.start(&Config{})

}

// Refer to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet.go#178

// Option is a functional option type for Kubelet
type Option func( /* *Kubelet */ Agent)

// Refer to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet.go#L226

// Dependencies is a bin for things we might consider "injected dependencies" -- objects constructed
// at runtime that are necessary for running the Kubelet. This is a temporary solution for grouping
// these objects while we figure out a more comprehensive dependency injection story for the Kubelet.
type Dependencies struct {
	Options []Option

	// Injected Dependencies
	//	Auth                    server.AuthInterface
	CAdvisorInterface cadvisor.Interface
	Cloud             cloudprovider.Interface
	//	ContainerManager        cm.ContainerManager
	//	DockerClientConfig      *dockershim.ClientConfig
	EventClient        v1core.EventsGetter
	HeartbeatClient    v1core.CoreV1Interface
	OnHeartbeatFailure func()
	KubeClient         clientset.Interface
	ExternalKubeClient clientset.Interface
	//	Mounter                 mount.Interface
	//	OOMAdjuster             *oom.OOMAdjuster
	//	OSInterface             kubecontainer.OSInterface
	//	PodConfig               *config.PodConfig
	Recorder record.EventRecorder
	//	Writer                  kubeio.Writer
	//	VolumePlugins           []volume.VolumePlugin
	//	DynamicPluginProber     volume.DynamicPluginProber
	//	TLSOptions              *server.TLSOptions
	//	KubeletConfigController *kubeletconfig.Controller
}

type Agent struct {
	Config *Config

	signal         os.Signal
	grpcClientConn *connection

	// Refer to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet.go#L883
	hostname        String
	nodeName        types.NodeName
	runtimeCache    kubecontainer.RuntimeCache
	kubeClient      clientset.Interface
	heartbeatClient v1core.CoreV1Interface
	iptClient       utilipt.Interface
	rootDirectory   string

	// Refer to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet.go#L916

	// cAdvisor used for container information.
	cadvisor cadvisor.Interface

	// Refer to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet.go#L935

	// nodeInfo knows how to get information about the node for this kubelet.
	nodeInfo predicates.NodeInfo

	// a list of node labels to register
	nodeLabels map[string]string

	// Refer to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet.go#L975

	// Cached MachineInfo returned by cadvisor.
	machineInfo *cadvisorapi.MachineInfo

	// Refer to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet.go#L1005

	// Reference to this node.
	nodeRef *v1.ObjectReference

	// Refer to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet.go#L1029

	// nodeStatusUpdateFrequency specifies how often kubelet posts node status to master.
	// Note: be cautious when changing the constant, it must work with nodeMonitorGracePeriod
	// in nodecontroller. There are several constraints:
	// 1. nodeMonitorGracePeriod must be N times more than nodeStatusUpdateFrequency, where
	//    N means number of retries allowed for kubelet to post node status. It is pointless
	//    to make nodeMonitorGracePeriod be less than nodeStatusUpdateFrequency, since there
	//    will only be fresh values from Kubelet at an interval of nodeStatusUpdateFrequency.
	//    The constant must be less than podEvictionTimeout.
	// 2. nodeStatusUpdateFrequency needs to be large enough for kubelet to generate node
	//    status. Kubelet may fail to update node status reliably if the value is too small,
	//    as it takes time to gather all necessary node information.
	nodeStatusUpdateFrequency time.Duration

	// Refer to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet.go#L1093

	// oneTimeInitializer is used to initialize modules that are dependent on the runtime to be up.
	oneTimeInitializer sync.Once

	// If non-nil, use this IP address for the node
	nodeIP net.IP

	// use this function to validate the kubelet nodeIP
	nodeIPValidator func(net.IP) error

	// Refer to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet.go#L1107

	// handlers called during the tryUpdateNodeStatus cycle
	setNodeStatusFuncs []func( /* *v1.Node */ *pb.AgentNode) error
}

// Refer to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet.go#L1318

// initializeRuntimeDependentModules will initialize internal modules that require the container runtime to be up.
//func (kl *Kubelet) initializeRuntimeDependentModules() {
func (a *Agent) initializeRuntimeDependentModules() {
	if err := kl.cadvisor.Start(); err != nil {
		// Fail kubelet and rely on the babysitter to retry starting kubelet.
		// TODO(random-liu): Add backoff logic in the babysitter
		glog.Fatalf("Failed to start cAdvisor %v", err)
	}

	// trigger on-demand stats collection once so that we have capacity information for ephemeral storage.
	// ignore any errors, since if stats collection is not successful, the container manager will fail to start below.
	kl.StatsProvider.GetCgroupStats("/", true)
	// Start container manager.
	node, err := kl.getNodeAnyWay()
	if err != nil {
		// Fail kubelet and rely on the babysitter to retry starting kubelet.
		glog.Fatalf("Kubelet failed to get node info: %v", err)
	}
	// containerManager must start after cAdvisor because it needs filesystem capacity information
	if err := kl.containerManager.Start(node, kl.GetActivePods, kl.sourcesReady, kl.statusManager, kl.runtimeService); err != nil {
		// Fail kubelet and rely on the babysitter to retry starting kubelet.
		glog.Fatalf("Failed to start ContainerManager %v", err)
	}
	// eviction manager must start after cadvisor because it needs to know if the container runtime has a dedicated imagefs
	kl.evictionManager.Start(kl.StatsProvider, kl.GetActivePods, kl.podResourcesAreReclaimed, evictionMonitoringPeriod)

	// container log manager must start after container runtime is up to retrieve information from container runtime
	// and inform container to reopen log file after log rotation.
	kl.containerLogManager.Start()
}

// Run starts the kubelet reacting to config updates
func (kl *Kubelet) Run(updates <-chan kubetypes.PodUpdate) {
	if kl.logServer == nil {
		kl.logServer = http.StripPrefix("/logs/", http.FileServer(http.Dir("/var/log/")))
	}
	if kl.kubeClient == nil {
		glog.Warning("No api server defined - no node status update will be sent.")
	}

	if err := kl.initializeModules(); err != nil {
		kl.recorder.Eventf(kl.nodeRef, v1.EventTypeWarning, events.KubeletSetupFailed, err.Error())
		glog.Fatal(err)
	}

	// Start volume manager
	go kl.volumeManager.Run(kl.sourcesReady, wait.NeverStop)

	if kl.kubeClient != nil {
		// Start syncing node status immediately, this may set up things the runtime needs to run.
		go wait.Until(kl.syncNodeStatus, kl.nodeStatusUpdateFrequency, wait.NeverStop)
	}
	go wait.Until(kl.updateRuntimeUp, 5*time.Second, wait.NeverStop)

	// Start loop to sync iptables util rules
	if kl.makeIPTablesUtilChains {
		go wait.Until(kl.syncNetworkUtil, 1*time.Minute, wait.NeverStop)
	}

	// Start a goroutine responsible for killing pods (that are not properly
	// handled by pod workers).
	go wait.Until(kl.podKiller, 1*time.Second, wait.NeverStop)

	// Start gorouting responsible for checking limits in resolv.conf
	if kl.dnsConfigurer.ResolverConfig != "" {
		go wait.Until(func() { kl.dnsConfigurer.CheckLimitsForResolvConf() }, 30*time.Second, wait.NeverStop)
	}

	// Start component sync loops.
	kl.statusManager.Start()
	kl.probeManager.Start()

	// Start the pod lifecycle event generator.
	kl.pleg.Start()
	kl.syncLoop(updates, kl)
}

func (a *Agent) start() {
	//	ch := make(chan string)
	//	wg := sync.WaitGroup{}

	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		s.startGRPC(ch)
	//	}()

	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		s.startGateway(ch)
	//	}()

	go func() {
		a.begin()
	}()

	/*
	   https://github.com/kubernetes/kubernetes/blob/release-1.1/build/pause/pause.go
	*/
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	// Block until a signal is received.
	s.signal = <-c

	go func() {
		done := make(chan bool, 1)
		a.terminate(done)
	}()
	fmt.Println(<-done)
	//wg.Wait()
}

func (a *Agent) terminate(chan<- bool) {

}

func (a *Agent) CaptureOpenstackIdentityData() {

}

func (a *Agent) CaptureOpenstackNetworkData() {

}

func (a *Agent) CaptureOpenstackComputeData() {

}

func (a *Agent) CaptureOpenstackMeterData() {

}

func (a *Agent) CaptureOpenstackData() {

}
