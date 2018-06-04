/*
  Inspired by:
  - https://github.com/kubernetes/kubernetes/blob/master/cmd/kubelet/app/init_others.go
*/

// +build windows

package app

import (
	"k8s.io/kubernetes/pkg/windows/service"
)

const (
	serviceName = "agent"
)

func initForOS(windowsService bool) error {
	if windowsService {
		return service.InitService(serviceName)
	}
	return nil
}
