package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"k8s.io/apiserver/pkg/util/logs"
	_ "k8s.io/kubernetes/pkg/client/metrics/prometheus" // for client metric registration
	_ "k8s.io/kubernetes/pkg/version/prometheus"        // for version metric registration

	"github.com/tangfeixiong/go-to-bigdata/nta/cmd"
	"github.com/tangfeixiong/go-to-bigdata/pkg/server"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	basename := filepath.Base(os.Args[0])
	command := cmd.RootCommandFor(basename, server.SetupSignalHandler())
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
