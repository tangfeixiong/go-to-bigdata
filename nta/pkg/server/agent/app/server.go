package agent

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/coreos/go-systemd/daemon"
	"github.com/golang/glog"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/server/healthz"
	"k8s.io/kubernetes/pkg/client/chaosclient"
	"k8s.io/kubernetes/pkg/cloudprovider"
	"k8s.io/kubernetes/pkg/kubelet/certificate/bootstrap"
	"k8s.io/kubernetes/pkg/kubelet/config"
	"k8s.io/kubernetes/pkg/util/flock"
	nodeutil "k8s.io/kubernetes/pkg/util/node"
	"k8s.io/kubernetes/pkg/util/nsenter"
	"k8s.io/kubernetes/pkg/version"

	"github.com/tangfeixiong/go-to-bigdata/nta/pkg/server/agent"
	"github.com/tangfeixiong/go-to-bigdata/nta/pkg/server/agent/app/options"
)

func Start(cfg *agent.Config, stopCh <-chan struct{}) {
	// Refer to https://github.com/kubernetes/kubernetes/blob/master/cmd/kubelet/app/server.go#236

	// construct a KubeletServer from kubeletFlags and kubeletConfig
	//	kubeletServer := &options.KubeletServer{
	//		KubeletFlags:         *kubeletFlags,
	//		KubeletConfiguration: *kubeletConfig,
	//	}
	agentServer := &options.AgentServer{}

	// use kubeletServer to construct the default KubeletDeps
	agentDeps /* kubeletDeps */, err := UnsecuredDependencies( /* kubeletServer */ agentServer)
	if err != nil {
		glog.Fatal(err)
	}

	// run the kubelet
	//	glog.V(5).Infof("KubeletConfiguration: %#v", kubeletServer.KubeletConfiguration)
	if err := Run( /* kubeletServer */ agentServer /* kubeletDeps */, agentDeps, stopCh); err != nil {
		glog.Fatal(err)
	}
}

// Refer to https://github.com/kubernetes/kubernetes/blob/master/cmd/kubelet/app/server.go#353

// UnsecuredDependencies returns a Dependencies suitable for being run, or an error if the server setup
// is not valid.  It will not start any background processes, and does not include authentication/authorization
func UnsecuredDependencies(s /* *options.KubeletServer */ *options.AgentServer) ( /* *kubelet.Dependencies */ *agent.Dependencies, error) {
	// Initialize the TLS Options
	//	tlsOptions, err := InitializeTLS(&s.KubeletFlags, &s.KubeletConfiguration)
	//	if err != nil {
	//		return nil, err
	//	}

	//	mounter := mount.New(s.ExperimentalMounterPath)
	//	var writer kubeio.Writer = &kubeio.StdWriter{}
	//	if s.Containerized {
	//		glog.V(2).Info("Running kubelet in containerized mode")
	//		ne, err := nsenter.NewNsenter(nsenter.DefaultHostRootFsPath, exec.New())
	//		if err != nil {
	//			return nil, err
	//		}
	//		mounter = mount.NewNsenterMounter(s.RootDirectory, ne)
	//		writer = kubeio.NewNsenterWriter(ne)
	//	}

	//	var dockerClientConfig *dockershim.ClientConfig
	//	if s.ContainerRuntime == kubetypes.DockerContainerRuntime {
	//		dockerClientConfig = &dockershim.ClientConfig{
	//			DockerEndpoint:            s.DockerEndpoint,
	//			RuntimeRequestTimeout:     s.RuntimeRequestTimeout.Duration,
	//			ImagePullProgressDeadline: s.ImagePullProgressDeadline.Duration,
	//		}
	//	}

	return &kubelet.Dependencies{
		//		Auth:                nil, // default does not enforce auth[nz]
		CAdvisorInterface: nil, // cadvisor.New launches background processes (bg http.ListenAndServe, and some bg cleaners), not set here
		//		Cloud:               nil, // cloud provider might start background processes
		//		ContainerManager:    nil,
		//		DockerClientConfig:  dockerClientConfig,
		KubeClient:         nil,
		HeartbeatClient:    nil,
		ExternalKubeClient: nil,
		EventClient:        nil,
		//		Mounter:             mounter,
		//		OOMAdjuster:         oom.NewOOMAdjuster(),
		//		OSInterface:         kubecontainer.RealOS{},
		//		Writer:              writer,
		//		VolumePlugins:       ProbeVolumePlugins(),
		//		DynamicPluginProber: GetDynamicPluginProber(s.VolumePluginDir),
		//		TLSOptions:          tlsOptions
	}, nil
}

// Run runs the specified KubeletServer with the given Dependencies. This should never exit.
// The kubeDeps argument may be nil - if so, it is initialized from the settings on KubeletServer.
// Otherwise, the caller is assumed to have set up the Dependencies object and a default one will
// not be generated.
func Run(s /* *options.KubeletServer */ *options.AgentServer /* kubeDeps */, agentDeps /* *kubelet.Dependencies */ *agent.Dependencies, stopCh <-chan struct{}) error {
	// To help debugging, immediately log version
	glog.Infof("Version: %+v", "0.0.1" /* version.Get() */)
	if err := initForOS( /* s.KubeletFlags.WindowsService */ s.AgentFlags.WindowsService); err != nil {
		return fmt.Errorf("failed OS init: %v", err)
	}
	if err := run(s, agentDeps /* kubeDeps*/, stopCh); err != nil {
		//		return fmt.Errorf("failed to run Kubelet: %v", err)
		return fmt.Errorf("failed to run Agent: %v", err)
	}
	return nil
}

func checkPermissions() error {
	if uid := os.Getuid(); uid != 0 {
		return fmt.Errorf("Kubelet needs to run as uid `0`. It is being run as %d", uid)
	}
	// TODO: Check if kubelet is running in the `initial` user namespace.
	// http://man7.org/linux/man-pages/man7/user_namespaces.7.html
	return nil
}

//func setConfigz(cz *configz.Config, kc *kubeletconfiginternal.KubeletConfiguration) error {
//	scheme, _, err := kubeletscheme.NewSchemeAndCodecs()
//	if err != nil {
//		return err
//	}
//	versioned := kubeletconfigv1beta1.KubeletConfiguration{}
//	if err := scheme.Convert(kc, &versioned, nil); err != nil {
//		return err
//	}
//	cz.Set(versioned)
//	return nil
//}

//func initConfigz(kc *kubeletconfiginternal.KubeletConfiguration) error {
//	cz, err := configz.New("kubeletconfig")
//	if err != nil {
//		glog.Errorf("unable to register configz: %s", err)
//		return err
//	}
//	if err := setConfigz(cz, kc); err != nil {
//		glog.Errorf("unable to register config: %s", err)
//		return err
//	}
//	return nil
//}

// makeEventRecorder sets up kubeDeps.Recorder if it's nil. It's a no-op otherwise.
//func makeEventRecorder(kubeDeps *kubelet.Dependencies, nodeName types.NodeName) {
//	if kubeDeps.Recorder != nil {
//		return
//	}
//	eventBroadcaster := record.NewBroadcaster()
//	kubeDeps.Recorder = eventBroadcaster.NewRecorder(legacyscheme.Scheme, v1.EventSource{Component: componentKubelet, Host: string(nodeName)})
//	eventBroadcaster.StartLogging(glog.V(3).Infof)
//	if kubeDeps.EventClient != nil {
//		glog.V(4).Infof("Sending events to api server.")
//		eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: kubeDeps.EventClient.Events("")})
//	} else {
//		glog.Warning("No api server defined - no events will be sent to API server.")
//	}
//}

func run(s /* *options.KubeletServer */ *options.AgentServer /* kubeDeps */, agentDeps /* *kubelet.Dependencies */ *agent.Dependencies, stopCh <-chan struct{}) (err error) {
	// Set global feature gates based on the value on the initial KubeletServer
	//	err = utilfeature.DefaultFeatureGate.SetFromMap(s.KubeletConfiguration.FeatureGates)
	//	if err != nil {
	//		return err
	//	}
	// validate the initial KubeletServer (we set feature gates first, because this validation depends on feature gates)
	//	if err := options.ValidateKubeletServer(s); err != nil {
	//		return err
	//	}

	// Obtain Kubelet Lock File
	if s.ExitOnLockContention && s.LockFilePath == "" {
		return errors.New("cannot exit on lock file contention: no lock file specified")
	}
	done := make(chan struct{})
	if s.LockFilePath != "" {
		glog.Infof("acquiring file lock on %q", s.LockFilePath)
		if err := flock.Acquire(s.LockFilePath); err != nil {
			return fmt.Errorf("unable to acquire file lock on %q: %v", s.LockFilePath, err)
		}
		if s.ExitOnLockContention {
			glog.Infof("watching for inotify events for: %v", s.LockFilePath)
			if err := watchForLockfileContention(s.LockFilePath, done); err != nil {
				return err
			}
		}
	}

	// Register current configuration with /configz endpoint
	//	err = initConfigz(&s.KubeletConfiguration)
	//	if err != nil {
	//		glog.Errorf("unable to register KubeletConfiguration with configz, error: %v", err)
	//	}

	// About to get clients and such, detect standaloneMode
	standaloneMode := true
	if len(s.KubeConfig) > 0 {
		standaloneMode = false
	}

	if /* kubeDeps */ agentDeps == nil {
		agentDeps /* kubeDeps*/, err = UnsecuredDependencies(s)
		if err != nil {
			return err
		}
	}

	//	if kubeDeps.Cloud == nil {
	//		if !cloudprovider.IsExternal(s.CloudProvider) {
	//			cloud, err := cloudprovider.InitCloudProvider(s.CloudProvider, s.CloudConfigFile)
	//			if err != nil {
	//				return err
	//			}
	//			if cloud == nil {
	//				glog.V(2).Infof("No cloud provider specified: %q from the config file: %q\n", s.CloudProvider, s.CloudConfigFile)
	//			} else {
	//				glog.V(2).Infof("Successfully initialized cloud provider: %q from the config file: %q\n", s.CloudProvider, s.CloudConfigFile)
	//			}
	//			kubeDeps.Cloud = cloud
	//		}
	//	}

	nodeName, err := getNodeName( /* kubeDeps.Cloud */ nil, nodeutil.GetHostname(s.HostnameOverride))
	if err != nil {
		return err
	}

	if s.BootstrapKubeconfig != "" {
		if err := bootstrap.LoadClientCert(s.KubeConfig, s.BootstrapKubeconfig, s.CertDirectory, nodeName); err != nil {
			return err
		}
	}

	// if in standalone mode, indicate as much by setting all clients to nil
	if standaloneMode {
		//		kubeDeps.KubeClient = nil
		//		kubeDeps.ExternalKubeClient = nil
		//		kubeDeps.EventClient = nil
		//		kubeDeps.HeartbeatClient = nil
		agentDeps.KubeClient = nil
		agentDeps.ExternalKubeClient = nil
		agentDeps.EventClient = nil
		agentDeps.HeartbeatClient = nil
		glog.Warningf("standalone mode, no API client")
	} else if kubeDeps.KubeClient == nil || kubeDeps.ExternalKubeClient == nil || kubeDeps.EventClient == nil || kubeDeps.HeartbeatClient == nil {
		// initialize clients if not standalone mode and any of the clients are not provided
		var kubeClient clientset.Interface
		var eventClient v1core.EventsGetter
		var heartbeatClient v1core.CoreV1Interface
		var externalKubeClient clientset.Interface

		clientConfig, err := createAPIServerClientConfig(s)
		if err != nil {
			return fmt.Errorf("invalid kubeconfig: %v", err)
		}

		//		var clientCertificateManager certificate.Manager
		//		if s.RotateCertificates && utilfeature.DefaultFeatureGate.Enabled(features.RotateKubeletClientCertificate) {
		//			clientCertificateManager, err = kubeletcertificate.NewKubeletClientCertificateManager(s.CertDirectory, nodeName, clientConfig.CertData, clientConfig.KeyData, clientConfig.CertFile, clientConfig.KeyFile)
		//			if err != nil {
		//				return err
		//			}
		//		}
		// we set exitAfter to five minutes because we use this client configuration to request new certs - if we are unable
		// to request new certs, we will be unable to continue normal operation. Exiting the process allows a wrapper
		// or the bootstrapping credentials to potentially lay down new initial config.
		//		closeAllConns, err := kubeletcertificate.UpdateTransport(wait.NeverStop, clientConfig, clientCertificateManager, 5*time.Minute)
		//		if err != nil {
		//			return err
		//		}

		kubeClient, err = clientset.NewForConfig(clientConfig)
		if err != nil {
			glog.Warningf("New kubeClient from clientConfig error: %v", err)
		} else if kubeClient.CertificatesV1beta1() != nil && clientCertificateManager != nil {
			glog.V(2).Info("Starting client certificate rotation.")
			clientCertificateManager.SetCertificateSigningRequestClient(kubeClient.CertificatesV1beta1().CertificateSigningRequests())
			clientCertificateManager.Start()
		}
		externalKubeClient, err = clientset.NewForConfig(clientConfig)
		if err != nil {
			glog.Warningf("New kubeClient from clientConfig error: %v", err)
		}

		// make a separate client for events
		eventClientConfig := *clientConfig
		//		eventClientConfig.QPS = float32(s.EventRecordQPS)
		//		eventClientConfig.Burst = int(s.EventBurst)
		eventClient, err = v1core.NewForConfig(&eventClientConfig)
		if err != nil {
			glog.Warningf("Failed to create API Server client for Events: %v", err)
		}

		// make a separate client for heartbeat with throttling disabled and a timeout attached
		heartbeatClientConfig := *clientConfig
		//		heartbeatClientConfig.Timeout = s.KubeletConfiguration.NodeStatusUpdateFrequency.Duration
		heartbeatClientConfig.QPS = float32(-1)
		heartbeatClient, err = v1core.NewForConfig(&heartbeatClientConfig)
		if err != nil {
			glog.Warningf("Failed to create API Server client for heartbeat: %v", err)
		}

		kubeDeps.KubeClient = kubeClient
		kubeDeps.ExternalKubeClient = externalKubeClient
		if heartbeatClient != nil {
			kubeDeps.HeartbeatClient = heartbeatClient
			//			kubeDeps.OnHeartbeatFailure = closeAllConns
		}
		if eventClient != nil {
			kubeDeps.EventClient = eventClient
		}
	}

	// If the kubelet config controller is available, and dynamic config is enabled, start the config and status sync loops
	//	if utilfeature.DefaultFeatureGate.Enabled(features.DynamicKubeletConfig) && len(s.DynamicConfigDir.Value()) > 0 &&
	//		kubeDeps.KubeletConfigController != nil && !standaloneMode && !s.RunOnce {
	//		if err := kubeDeps.KubeletConfigController.StartSync(kubeDeps.KubeClient, kubeDeps.EventClient, string(nodeName)); err != nil {
	//			return err
	//		}
	//	}

	//	if kubeDeps.Auth == nil {
	//		auth, err := BuildAuth(nodeName, kubeDeps.ExternalKubeClient, s.KubeletConfiguration)
	//		if err != nil {
	//			return err
	//		}
	//		kubeDeps.Auth = auth
	//	}

	if kubeDeps.CAdvisorInterface == nil {
		imageFsInfoProvider := cadvisor.NewImageFsInfoProvider(s.ContainerRuntime, s.RemoteRuntimeEndpoint)
		kubeDeps.CAdvisorInterface, err = cadvisor.New(s.Address, uint(s.CAdvisorPort), imageFsInfoProvider, s.RootDirectory, cadvisor.UsingLegacyCadvisorStats(s.ContainerRuntime, s.RemoteRuntimeEndpoint))
		if err != nil {
			return err
		}
	}

	// Setup event recorder if required.
	//	makeEventRecorder(kubeDeps, nodeName)

	//	if kubeDeps.ContainerManager == nil {
	//		if s.CgroupsPerQOS && s.CgroupRoot == "" {
	//			glog.Infof("--cgroups-per-qos enabled, but --cgroup-root was not specified.  defaulting to /")
	//			s.CgroupRoot = "/"
	//		}
	//		kubeReserved, err := parseResourceList(s.KubeReserved)
	//		if err != nil {
	//			return err
	//		}
	//		systemReserved, err := parseResourceList(s.SystemReserved)
	//		if err != nil {
	//			return err
	//		}
	//		var hardEvictionThresholds []evictionapi.Threshold
	//		// If the user requested to ignore eviction thresholds, then do not set valid values for hardEvictionThresholds here.
	//		if !s.ExperimentalNodeAllocatableIgnoreEvictionThreshold {
	//			hardEvictionThresholds, err = eviction.ParseThresholdConfig([]string{}, s.EvictionHard, nil, nil, nil)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//		experimentalQOSReserved, err := cm.ParseQOSReserved(s.QOSReserved)
	//		if err != nil {
	//			return err
	//		}

	//		devicePluginEnabled := utilfeature.DefaultFeatureGate.Enabled(features.DevicePlugins)

	//		kubeDeps.ContainerManager, err = cm.NewContainerManager(
	//			kubeDeps.Mounter,
	//			kubeDeps.CAdvisorInterface,
	//			cm.NodeConfig{
	//				RuntimeCgroupsName:    s.RuntimeCgroups,
	//				SystemCgroupsName:     s.SystemCgroups,
	//				KubeletCgroupsName:    s.KubeletCgroups,
	//				ContainerRuntime:      s.ContainerRuntime,
	//				CgroupsPerQOS:         s.CgroupsPerQOS,
	//				CgroupRoot:            s.CgroupRoot,
	//				CgroupDriver:          s.CgroupDriver,
	//				KubeletRootDir:        s.RootDirectory,
	//				ProtectKernelDefaults: s.ProtectKernelDefaults,
	//				NodeAllocatableConfig: cm.NodeAllocatableConfig{
	//					KubeReservedCgroupName:   s.KubeReservedCgroup,
	//					SystemReservedCgroupName: s.SystemReservedCgroup,
	//					EnforceNodeAllocatable:   sets.NewString(s.EnforceNodeAllocatable...),
	//					KubeReserved:             kubeReserved,
	//					SystemReserved:           systemReserved,
	//					HardEvictionThresholds:   hardEvictionThresholds,
	//				},
	//				QOSReserved:                           *experimentalQOSReserved,
	//				ExperimentalCPUManagerPolicy:          s.CPUManagerPolicy,
	//				ExperimentalCPUManagerReconcilePeriod: s.CPUManagerReconcilePeriod.Duration,
	//				ExperimentalPodPidsLimit:              s.PodPidsLimit,
	//				EnforceCPULimits:                      s.CPUCFSQuota,
	//			},
	//			s.FailSwapOn,
	//			devicePluginEnabled,
	//			kubeDeps.Recorder)

	//		if err != nil {
	//			return err
	//		}
	//	}

	if err := checkPermissions(); err != nil {
		glog.Error(err)
	}

	//	utilruntime.ReallyCrash = s.ReallyCrashForTesting

	rand.Seed(time.Now().UTC().UnixNano())

	// TODO(vmarmol): Do this through container config.
	//	oomAdjuster := kubeDeps.OOMAdjuster
	//	if err := oomAdjuster.ApplyOOMScoreAdj(0, int(s.OOMScoreAdj)); err != nil {
	//		glog.Warning(err)
	//	}

	if err := RunKubelet(&s.KubeletFlags, &s.KubeletConfiguration, kubeDeps, s.RunOnce, stopCh); err != nil {
		return err
	}

	if s.HealthzPort > 0 {
		healthz.DefaultHealthz()
		go wait.Until(func() {
			err := http.ListenAndServe(net.JoinHostPort(s.HealthzBindAddress, strconv.Itoa(int(s.HealthzPort))), nil)
			if err != nil {
				glog.Errorf("Starting health server failed: %v", err)
			}
		}, 5*time.Second, wait.NeverStop)
	}

	if s.RunOnce {
		return nil
	}

	// If systemd is used, notify it that we have started
	go daemon.SdNotify(false, "READY=1")

	select {
	case <-done:
		break
	case <-stopCh:
		break
	}

	return nil
}

// getNodeName returns the node name according to the cloud provider
// if cloud provider is specified. Otherwise, returns the hostname of the node.
func getNodeName(cloud cloudprovider.Interface, hostname string) (types.NodeName, error) {
	if cloud == nil {
		return types.NodeName(hostname), nil
	}

	instances, ok := cloud.Instances()
	if !ok {
		return "", fmt.Errorf("failed to get instances from cloud provider")
	}

	nodeName, err := instances.CurrentNodeName(context.TODO(), hostname)
	if err != nil {
		return "", fmt.Errorf("error fetching current node name from cloud provider: %v", err)
	}

	glog.V(2).Infof("cloud provider determined current node name to be %s", nodeName)

	return nodeName, nil
}

// InitializeTLS checks for a configured TLSCertFile and TLSPrivateKeyFile: if unspecified a new self-signed
// certificate and key file are generated. Returns a configured server.TLSOptions object.
//func InitializeTLS(kf *options.KubeletFlags, kc *kubeletconfiginternal.KubeletConfiguration) (*server.TLSOptions, error) {
//	if !kc.ServerTLSBootstrap && kc.TLSCertFile == "" && kc.TLSPrivateKeyFile == "" {
//		kc.TLSCertFile = path.Join(kf.CertDirectory, "kubelet.crt")
//		kc.TLSPrivateKeyFile = path.Join(kf.CertDirectory, "kubelet.key")

//		canReadCertAndKey, err := certutil.CanReadCertAndKey(kc.TLSCertFile, kc.TLSPrivateKeyFile)
//		if err != nil {
//			return nil, err
//		}
//		if !canReadCertAndKey {
//			cert, key, err := certutil.GenerateSelfSignedCertKey(nodeutil.GetHostname(kf.HostnameOverride), nil, nil)
//			if err != nil {
//				return nil, fmt.Errorf("unable to generate self signed cert: %v", err)
//			}

//			if err := certutil.WriteCert(kc.TLSCertFile, cert); err != nil {
//				return nil, err
//			}

//			if err := certutil.WriteKey(kc.TLSPrivateKeyFile, key); err != nil {
//				return nil, err
//			}

//			glog.V(4).Infof("Using self-signed cert (%s, %s)", kc.TLSCertFile, kc.TLSPrivateKeyFile)
//		}
//	}

//	tlsCipherSuites, err := flag.TLSCipherSuites(kc.TLSCipherSuites)
//	if err != nil {
//		return nil, err
//	}

//	minTLSVersion, err := flag.TLSVersion(kc.TLSMinVersion)
//	if err != nil {
//		return nil, err
//	}

//	tlsOptions := &server.TLSOptions{
//		Config: &tls.Config{
//			MinVersion:   minTLSVersion,
//			CipherSuites: tlsCipherSuites,
//		},
//		CertFile: kc.TLSCertFile,
//		KeyFile:  kc.TLSPrivateKeyFile,
//	}

//	if len(kc.Authentication.X509.ClientCAFile) > 0 {
//		clientCAs, err := certutil.NewPool(kc.Authentication.X509.ClientCAFile)
//		if err != nil {
//			return nil, fmt.Errorf("unable to load client CA file %s: %v", kc.Authentication.X509.ClientCAFile, err)
//		}
//		// Specify allowed CAs for client certificates
//		tlsOptions.Config.ClientCAs = clientCAs
//		// Populate PeerCertificates in requests, but don't reject connections without verified certificates
//		tlsOptions.Config.ClientAuth = tls.RequestClientCert
//	}

//	return tlsOptions, nil
//}

func kubeconfigClientConfig(s /* *options.KubeletServer */ *options.AgentServer) (*restclient.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: s.KubeConfig},
		&clientcmd.ConfigOverrides{},
	).ClientConfig()
}

// createClientConfig creates a client configuration from the command line arguments.
// If --kubeconfig is explicitly set, it will be used.
func createClientConfig(s /* *options.KubeletServer */ *options.AgentServer) (*restclient.Config, error) {
	if s.BootstrapKubeconfig != "" || len(s.KubeConfig) > 0 {
		return kubeconfigClientConfig(s)
	} else {
		return nil, fmt.Errorf("createClientConfig called in standalone mode")
	}
}

// createAPIServerClientConfig generates a client.Config from command line flags
// via createClientConfig and then injects chaos into the configuration via addChaosToClientConfig.
func createAPIServerClientConfig(s /* *options.KubeletServer */ *options.AgentServer) (*restclient.Config, error) {
	clientConfig, err := createClientConfig(s)
	if err != nil {
		return nil, err
	}

	//	clientConfig.ContentType = s.ContentType
	// Override kubeconfig qps/burst settings from flags
	//	clientConfig.QPS = float32(s.KubeAPIQPS)
	//	clientConfig.Burst = int(s.KubeAPIBurst)

	addChaosToClientConfig(s, clientConfig)
	return clientConfig, nil
}

// addChaosToClientConfig injects random errors into client connections if configured.
func addChaosToClientConfig(s *options.KubeletServer, config *restclient.Config) {
	if s.ChaosChance != 0.0 {
		config.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
			seed := chaosclient.NewSeed(1)
			// TODO: introduce a standard chaos package with more tunables - this is just a proof of concept
			// TODO: introduce random latency and stalls
			return chaosclient.NewChaosRoundTripper(rt, chaosclient.LogChaos, seed.P(s.ChaosChance, chaosclient.ErrSimulatedConnectionResetByPeer))
		}
	}
}

// Refer to https://github.com/kubernetes/kubernetes/blob/master/cmd/kubelet/app/server.go#885

// RunKubelet is responsible for setting up and running a kubelet.  It is used in three different applications:
//   1 Integration tests
//   2 Kubelet binary
//   3 Standalone 'kubernetes' binary
// Eventually, #2 will be replaced with instances of #3
// func RunKubelet(kubeFlags *options.KubeletFlags, kubeCfg *kubeletconfiginternal.KubeletConfiguration, kubeDeps *kubelet.Dependencies, runOnce bool, stopCh <-chan struct{}) error {
func RunAgent(agentFlags *options.AgentFlags, kubeCfg *kubeletconfiginternal.KubeletConfiguration, agentDeps *agent.Dependencies, runOnce bool, stopCh <-chan struct{}) error {
	hostname := nodeutil.GetHostname(kubeFlags.HostnameOverride)
	// Query the cloud provider for our node name, default to hostname if kubeDeps.Cloud == nil
	nodeName, err := getNodeName( /* kubeDeps.Cloud */ nil, hostname)
	if err != nil {
		return err
	}
	// Setup event recorder if required.
	//	makeEventRecorder(kubeDeps, nodeName)

	// TODO(mtaufen): I moved the validation of these fields here, from UnsecuredKubeletConfig,
	//                so that I could remove the associated fields from KubeletConfiginternal. I would
	//                prefer this to be done as part of an independent validation step on the
	//                KubeletConfiguration. But as far as I can tell, we don't have an explicit
	//                place for validation of the KubeletConfiguration yet.
	//	hostNetworkSources, err := kubetypes.GetValidatedSources(kubeFlags.HostNetworkSources)
	//	if err != nil {
	//		return err
	//	}

	//	hostPIDSources, err := kubetypes.GetValidatedSources(kubeFlags.HostPIDSources)
	//	if err != nil {
	//		return err
	//	}

	//	hostIPCSources, err := kubetypes.GetValidatedSources(kubeFlags.HostIPCSources)
	//	if err != nil {
	//		return err
	//	}

	//	privilegedSources := capabilities.PrivilegedSources{
	//		HostNetworkSources: hostNetworkSources,
	//		HostPIDSources:     hostPIDSources,
	//		HostIPCSources:     hostIPCSources,
	//	}
	//	capabilities.Setup(kubeFlags.AllowPrivileged, privilegedSources, 0)

	//	credentialprovider.SetPreferredDockercfgPath(kubeFlags.RootDirectory)
	//	glog.V(2).Infof("Using root directory: %v", kubeFlags.RootDirectory)

	//	if kubeDeps.OSInterface == nil {
	//		kubeDeps.OSInterface = kubecontainer.RealOS{}
	//	}

	k, err := CreateAndInitKubelet(kubeCfg,
		kubeDeps,
		&kubeFlags.ContainerRuntimeOptions,
		kubeFlags.ContainerRuntime,
		kubeFlags.RuntimeCgroups,
		kubeFlags.HostnameOverride,
		kubeFlags.NodeIP,
		kubeFlags.ProviderID,
		kubeFlags.CloudProvider,
		kubeFlags.CertDirectory,
		kubeFlags.RootDirectory,
		kubeFlags.RegisterNode,
		kubeFlags.RegisterWithTaints,
		kubeFlags.AllowedUnsafeSysctls,
		kubeFlags.RemoteRuntimeEndpoint,
		kubeFlags.RemoteImageEndpoint,
		kubeFlags.ExperimentalMounterPath,
		kubeFlags.ExperimentalKernelMemcgNotification,
		kubeFlags.ExperimentalCheckNodeCapabilitiesBeforeMount,
		kubeFlags.ExperimentalNodeAllocatableIgnoreEvictionThreshold,
		kubeFlags.MinimumGCAge,
		kubeFlags.MaxPerPodContainerCount,
		kubeFlags.MaxContainerCount,
		kubeFlags.MasterServiceNamespace,
		kubeFlags.RegisterSchedulable,
		kubeFlags.NonMasqueradeCIDR,
		kubeFlags.KeepTerminatedPodVolumes,
		kubeFlags.NodeLabels,
		kubeFlags.SeccompProfileRoot,
		kubeFlags.BootstrapCheckpointPath,
		kubeFlags.NodeStatusMaxImages,
		stopCh)
	if err != nil {
		//		return fmt.Errorf("failed to create kubelet: %v", err)
		return fmt.Errorf("failed to create agent: %v", err)
	}

	// NewMainKubelet should have set up a pod source config if one didn't exist
	// when the builder was run. This is just a precaution.
	if kubeDeps.PodConfig == nil {
		return fmt.Errorf("failed to create kubelet, pod source config was nil")
	}
	podCfg := kubeDeps.PodConfig

	rlimit.RlimitNumFiles(uint64(kubeCfg.MaxOpenFiles))

	// process pods and exit.
	if runOnce {
		if _, err := k.RunOnce(podCfg.Updates()); err != nil {
			return fmt.Errorf("runonce failed: %v", err)
		}
		glog.Infof("Started kubelet as runonce")
	} else {
		//		startKubelet(k, podCfg, kubeCfg, kubeDeps, kubeFlags.EnableServer)
		startAgent(k, podCfg, kubeCfg, kubeDeps, kubeFlags.EnableServer)
		glog.Infof("Started kubelet")
	}
	return nil
}

//func startKubelet(k kubelet.Bootstrap, podCfg *config.PodConfig, kubeCfg *kubeletconfiginternal.KubeletConfiguration, kubeDeps *kubelet.Dependencies, enableServer bool) {
func startAgent(k kubelet.Bootstrap, podCfg *config.PodConfig, kubeCfg *kubeletconfiginternal.KubeletConfiguration, kubeDeps *kubelet.Dependencies, enableServer bool) {
	wg := sync.WaitGroup{}

	// start the kubelet
	wg.Add(1)
	go wait.Until(func() {
		wg.Done()
		k.Run(podCfg.Updates())
	}, 0, wait.NeverStop)

	// start the kubelet server
	if enableServer {
		wg.Add(1)
		go wait.Until(func() {
			wg.Done()
			k.ListenAndServe(net.ParseIP(kubeCfg.Address), uint(kubeCfg.Port), kubeDeps.TLSOptions, kubeDeps.Auth, kubeCfg.EnableDebuggingHandlers, kubeCfg.EnableContentionProfiling)
		}, 0, wait.NeverStop)
	}
	if kubeCfg.ReadOnlyPort > 0 {
		wg.Add(1)
		go wait.Until(func() {
			wg.Done()
			k.ListenAndServeReadOnly(net.ParseIP(kubeCfg.Address), uint(kubeCfg.ReadOnlyPort))
		}, 0, wait.NeverStop)
	}
	wg.Wait()
}

func CreateAndInitKubelet(kubeCfg *kubeletconfiginternal.KubeletConfiguration,
	kubeDeps *kubelet.Dependencies,
	crOptions *config.ContainerRuntimeOptions,
	containerRuntime string,
	runtimeCgroups string,
	hostnameOverride string,
	nodeIP string,
	providerID string,
	cloudProvider string,
	certDirectory string,
	rootDirectory string,
	registerNode bool,
	registerWithTaints []api.Taint,
	allowedUnsafeSysctls []string,
	remoteRuntimeEndpoint string,
	remoteImageEndpoint string,
	experimentalMounterPath string,
	experimentalKernelMemcgNotification bool,
	experimentalCheckNodeCapabilitiesBeforeMount bool,
	experimentalNodeAllocatableIgnoreEvictionThreshold bool,
	minimumGCAge metav1.Duration,
	maxPerPodContainerCount int32,
	maxContainerCount int32,
	masterServiceNamespace string,
	registerSchedulable bool,
	nonMasqueradeCIDR string,
	keepTerminatedPodVolumes bool,
	nodeLabels map[string]string,
	seccompProfileRoot string,
	bootstrapCheckpointPath string,
	nodeStatusMaxImages int32,
	stopCh <-chan struct{}) (k kubelet.Bootstrap, err error) {
	// TODO: block until all sources have delivered at least one update to the channel, or break the sync loop
	// up into "per source" synchronizations

	k, err = kubelet.NewMainKubelet(kubeCfg,
		kubeDeps,
		crOptions,
		containerRuntime,
		runtimeCgroups,
		hostnameOverride,
		nodeIP,
		providerID,
		cloudProvider,
		certDirectory,
		rootDirectory,
		registerNode,
		registerWithTaints,
		allowedUnsafeSysctls,
		remoteRuntimeEndpoint,
		remoteImageEndpoint,
		experimentalMounterPath,
		experimentalKernelMemcgNotification,
		experimentalCheckNodeCapabilitiesBeforeMount,
		experimentalNodeAllocatableIgnoreEvictionThreshold,
		minimumGCAge,
		maxPerPodContainerCount,
		maxContainerCount,
		masterServiceNamespace,
		registerSchedulable,
		nonMasqueradeCIDR,
		keepTerminatedPodVolumes,
		nodeLabels,
		seccompProfileRoot,
		bootstrapCheckpointPath,
		nodeStatusMaxImages,
		stopCh)
	if err != nil {
		return nil, err
	}

	k.BirthCry()

	k.StartGarbageCollection()

	return k, nil
}
