/*
  Inspired by:
  - https://github.com/kubernetes/kubernetes/blob/master/cmd/kubelet/app/server_unsupported.go
*/

// +build !linux

package app

import "errors"

func watchForLockfileContention(path string, done chan struct{}) error {
	return errors.New("kubelet unsupported in this build")
}
