package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	goruntime "runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/coreos/etcd/raft/raftpb"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/tangfeixiong/go-to-bigdata/raft-grpc/pb"
	//"github.com/tangfeixiong/go-to-bigdata/raft-grpc/pkg/..."
)

type myServer struct {
	clusterName string
	nodeName    string
	dataStore   string
	peerHost    string
	grpcHost    string
	httpHost    string

	packgesHome            string
	redisSentinelAddresses string
	redisAddresses         []string
	redisDB                int
	etcdAddresses          string
	mysqlAddress           string
	gnatsdAddresses        string
	kafkaAddresses         string
	zookeeperAddresses     string
	rabbitAddress          string
	priorityCMDB           []string
	priorityMQ             []string
	subSubject             string
	subQueue               string
	pubSubject             string
	cmCache                string
	unsubCh                chan string
	dispatchersignal       chan bool
}

type myCreation struct {
	myServer
	//exportermanager *exporter.ExporterManager
	peers       []string
	commitC     []<-chan *string
	errorC      []<-chan error
	proposeC    []chan string
	confChangeC []chan raftpb.ConfChange
}

//type myCollector struct {
//	myServer
//	collectormanager *collector.CollectorManager
//}

type ServerConfigure interface {
	ClusterNamed(name string) ServerConfigure
	NodeNamed(name string) ServerConfigure
	DataDrived(driver string) ServerConfigure
	Peers(ipv4, port string) ServerConfigure
	GrpcHost(host string) ServerConfigure
	HttpHost(host string) ServerConfigure
	Run()
}

type ServerOptionFunc func(sc ServerConfigure) error

func PeerHostOption(ipv4 string, port string) ServerOptionFunc {
	return func(sc ServerConfigure) error {
		sc.Peers(ipv4, port)
		return nil
	}
}

func GrpcHostOption(addr string) ServerOptionFunc {
	return func(sc ServerConfigure) error {
		sc.GrpcHost(addr)
		return nil
	}
}

func HttpHostOption(addr string) ServerOptionFunc {
	return func(sc ServerConfigure) error {
		sc.HttpHost(addr)
		return nil
	}
}

func NewServer(options ...ServerOptionFunc) *myCreation {
	m := new(myCreation)
	m.peerHost = "127.0.0.1:12347"
	m.grpcHost = ":12345"
	m.httpHost = ":12346"

	//cluster := flag.String("cluster", "http://127.0.0.1:9021", "comma separated cluster peers")
	//id := flag.Int("id", 1, "node ID")
	//kvport := flag.Int("port", 9121, "key-value server port")
	//join := flag.Bool("join", false, "join an existing cluster")
	//flag.Parse()
	var cluster *string
	*cluster = "http://127.0.0.1:12346"
	var id *int
	*id = 1
	var kvport *int
	*kvport = 12347
	var join *bool
	*join = false

	proposeC := make(chan string)
	defer close(proposeC)
	confChangeC := make(chan raftpb.ConfChange)
	defer close(confChangeC)

	// raft provides a commit stream for the proposals from the http api
	var kvs *kvstore
	getSnapshot := func() ([]byte, error) { return kvs.getSnapshot() }
	commitC, errorC, snapshotterReady := newRaftNode(*id, strings.Split(*cluster, ","), *join, getSnapshot, proposeC, confChangeC)

	kvs = newKVStore(<-snapshotterReady, proposeC, commitC, errorC)

	// the key-value http handler will propose updates to raft
	serveHttpKVAPI(kvs, *kvport, confChangeC, errorC)

	//m.exportermanager = new(exporter.ExporterManager)
	//m.exportermanager.Dispatchers = make(map[string]exporter.MeterDispatcher)
	//m.exportermanager.MeteringNameURLs = make(map[string][]string)
	// m.exportermanager.MetricsCollectorRPC = "localhost:12305"
	// m.exportermanager.MeteringNameURLs["docker"] = []string{"unix:///var/run/docker.sock"}

	if v, ok := os.LookupEnv("METERINGEXPORTER_GRPC_PORT"); ok && 0 != len(v) {
		if strings.Contains(v, ":") {
			m.grpcHost = v
		} else {
			m.grpcHost = "localhost:" + v
		}
	}
	if v, ok := os.LookupEnv("METERINGEXPORTER_HTTP_PORT"); ok && 0 != len(v) {
		if strings.Contains(v, ":") {
			m.httpHost = v
		} else {
			m.httpHost = ":" + v
		}
	}

	if v, ok := os.LookupEnv("DATA_STORE_DRIVER"); ok && 0 != len(v) {
		for _, pair := range strings.Split(v, ",") {
			nu := strings.Split(pair, "=")
			if len(nu) != 2 {
				panic("Invalid environment value: DATA_STORE_DRIVER")
			}
		}
		m.dataStore = v
	}

	return m
}

func (s *myCreation) ClusterNamed(name string) *myCreation {
	s.clusterName = name
	return s
}

func (s *myCreation) NodeNamed(name string) *myCreation {
	s.nodeName = name
	return s
}

func (s *myCreation) DataDrived(driver string) *myCreation {
	s.dataStore = driver
	return s
}

func (s *myCreation) Peers(ipv4, port string) *myCreation {
	s.peerHost = strings.Join([]string{ipv4, port}, "")
	return s
}

func (s *myCreation) GrpcHost(host string) *myCreation {
	s.grpcHost = host
	return s
}

func (s *myCreation) HttpHost(host string) *myCreation {
	s.httpHost = host
	return s
}

func (m *myCreation) Run() {
	wg := sync.WaitGroup{}
	ch := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()
		m.startGRPC(ch)
	}()

	select {
	case <-ch:
		return
	case <-time.After(time.Millisecond * 500):
		fmt.Println("So gRPC running")
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		m.startGateway()
	}()

	m.dispatchersignal = make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		/*
		   default to read from 'docker stats'?
		*/
		// m.exportermanager.Dispatch(m.dispatchersignal)
	}()

	/*
	   https://github.com/kubernetes/kubernetes/blob/release-1.1/build/pause/pause.go
	*/
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Block until a signal is received.
		<-c

		// to stop stuff...
		m.dispatchersignal <- false
		goruntime.Goexit()
	}()

	wg.Wait()
}

func (m *myCreation) startGRPC(ch chan<- bool) {
	host := m.grpcHost

	s := grpc.NewServer()
	pb.RegisterRaftReplicaServiceServer(s, m)

	l, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	fmt.Println("Start gRPC on host", l.Addr())
	if err := s.Serve(l); nil != err {
		panic(err)
	}
	ch <- false
}

func (m *myCreation) startGateway() {
	gRPCHost := m.grpcHost
	host := m.httpHost

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	// mux.HandleFunc("/swagger/", serveSwagger2)
	//	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
	//		io.Copy(w, strings.NewReader(healthcheckerpb.Swagger))
	//	})

	dopts := []grpc.DialOption{grpc.WithInsecure()}

	fmt.Println("Start gRPC Gateway into host", gRPCHost)
	gwmux := runtime.NewServeMux()
	if err := pb.RegisterRaftReplicaServiceHandlerFromEndpoint(ctx, gwmux, gRPCHost, dopts); err != nil {
		fmt.Println("Failed to run HTTP server. ", err)
		return
	}

	mux.Handle("/", gwmux)
	// serveSwagger(mux)
	//	fmt.Printf("Start HTTP")
	//	if err := http.ListenAndServe(host, allowCORS(mux)); nil != err {
	//		fmt.Fprintf(os.Stderr, "Server died: %s\n", err)
	//	}

	lstn, err := net.Listen("tcp", host)
	if nil != err {
		panic(err)
	}

	fmt.Printf("http on host: %s\n", lstn.Addr())
	srv := &http.Server{
		Handler: func /*allowCORS*/ (h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if origin := r.Header.Get("Origin"); origin != "" {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
						func /*preflightHandler*/ (w http.ResponseWriter, r *http.Request) {
							headers := []string{"Content-Type", "Accept"}
							w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
							methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
							w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
							glog.Infof("preflight request for %s", r.URL.Path)
							return
						}(w, r)
						return
					}
				}
				h.ServeHTTP(w, r)
			})
		}(mux),
	}

	if err := srv.Serve(lstn); nil != err {
		fmt.Fprintln(os.Stderr, "Server died.", err.Error())
	}
}

//func RunCollector(storage string) {
//	m := new(myCollector)
//	m.grpcHost = ":12305"
//	m.httpHost = ":12306"
//	m.collectormanager = new(collector.CollectorManager)
//	m.collectormanager.MetricsStorageDuration = time.Second * 10

//	if v, ok := os.LookupEnv("METERINGCOLLECTOR_GRPC_PORT"); ok && 0 != len(v) {
//		if strings.Contains(v, ":") {
//			m.grpcHost = v
//		} else {
//			m.grpcHost = "localhost:" + v
//		}
//	}
//	if v, ok := os.LookupEnv("METERINGCOLLECTOR_HTTP_PORT"); ok && 0 != len(v) {
//		if strings.Contains(v, ":") {
//			m.httpHost = v
//		} else {
//			m.httpHost = ":" + v
//		}
//	}

//	if storage != "" {
//		if err := os.Setenv("METRICS_STORAGE_DRIVER", storage); err != nil {
//			fmt.Println("Environment not set, error:", err)
//		}
//	}
//	if v, ok := os.LookupEnv("METRICS_STORAGE_DRIVER"); ok && 0 != len(v) {
//		if strings.HasPrefix(v, "elasticsearch=") {
//			m.collectormanager.MetricsStorageDriver = v
//		} else {
//			fmt.Println("RPC driver not support,", v)
//		}
//	}

//	m.run()
//}

//func (m *myCollector) run() {
//	wg := sync.WaitGroup{}
//	ch := make(chan bool)

//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		m.startGRPC(ch)
//	}()

//	select {
//	case <-ch:
//		return
//	case <-time.After(time.Millisecond * 500):
//		fmt.Println("So gRPC running")
//	}

//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		m.startGateway()
//	}()

//	m.dispatchersignal = make(chan bool)
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		/*
//		   default to write to stdout?
//		*/
//		m.collectormanager.Start(m.dispatchersignal)
//	}()

//	/*
//	   https://github.com/kubernetes/kubernetes/blob/release-1.1/build/pause/pause.go
//	*/
//	c := make(chan os.Signal, 1)
//	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		// Block until a signal is received.
//		<-c

//		// to stop stuff...
//		m.dispatchersignal <- false
//		goruntime.Goexit()
//	}()

//	wg.Wait()
//}

//func (m *myCollector) startGRPC(ch chan<- bool) {
//	host := m.grpcHost

//	s := grpc.NewServer()
//	pb.RegisterCollectorServiceServer(s, m)

//	l, err := net.Listen("tcp", host)
//	if err != nil {
//		panic(err)
//	}

//	fmt.Println("Start gRPC on host", l.Addr())
//	if err := s.Serve(l); nil != err {
//		panic(err)
//	}
//	ch <- false
//}

//func (m *myCollector) startGateway() {
//	gRPCHost := m.grpcHost
//	host := m.httpHost

//	ctx := context.Background()
//	ctx, cancel := context.WithCancel(ctx)
//	defer cancel()

//	mux := http.NewServeMux()
//	// mux.HandleFunc("/swagger/", serveSwagger2)
//	//	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
//	//		io.Copy(w, strings.NewReader(healthcheckerpb.Swagger))
//	//	})

//	dopts := []grpc.DialOption{grpc.WithInsecure()}

//	fmt.Println("Start gRPC Gateway into host", gRPCHost)
//	gwmux := runtime.NewServeMux()
//	if err := pb.RegisterCollectorServiceHandlerFromEndpoint(ctx, gwmux, gRPCHost, dopts); err != nil {
//		fmt.Println("Failed to run HTTP server. ", err)
//		return
//	}

//	mux.Handle("/", gwmux)
//	// serveSwagger(mux)
//	//	fmt.Printf("Start HTTP")
//	//	if err := http.ListenAndServe(host, allowCORS(mux)); nil != err {
//	//		fmt.Fprintf(os.Stderr, "Server died: %s\n", err)
//	//	}

//	lstn, err := net.Listen("tcp", host)
//	if nil != err {
//		panic(err)
//	}

//	fmt.Printf("http on host: %s\n", lstn.Addr())
//	srv := &http.Server{
//		Handler: func /*allowCORS*/ (h http.Handler) http.Handler {
//			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//				if origin := r.Header.Get("Origin"); origin != "" {
//					w.Header().Set("Access-Control-Allow-Origin", origin)
//					if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
//						func /*preflightHandler*/ (w http.ResponseWriter, r *http.Request) {
//							headers := []string{"Content-Type", "Accept"}
//							w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
//							methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
//							w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
//							glog.Infof("preflight request for %s", r.URL.Path)
//							return
//						}(w, r)
//						return
//					}
//				}
//				h.ServeHTTP(w, r)
//			})
//		}(mux),
//	}

//	if err := srv.Serve(lstn); nil != err {
//		fmt.Fprintln(os.Stderr, "Server died.", err.Error())
//	}
//}
