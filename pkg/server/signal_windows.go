/*
  Inspired by:
  - https://github.com/kubernetes/apiserver/blob/master/pkg/server/signal_windows.go

  Build Constraints: all windows
*/

package server

import (
	"os"
)

var shutdownSignals = []os.Signal{os.Interrupt}
