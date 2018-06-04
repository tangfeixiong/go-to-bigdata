/*
  Inspired by:
  - https://github.com/kubernetes/apiserver/blob/master/pkg/server/signal_posix.go

  Build Constraints: *nix
*/

// +build !windows

package server

import (
	"os"
	"syscall"
)

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}
