/*
  Inspired by:
  - https://github.com/kubernetes/kubernetes/blob/master/cmd/kubelet/app/init_others.go
*/

// +build !windows

package app

func initForOS(service bool) error {
	return nil
}
