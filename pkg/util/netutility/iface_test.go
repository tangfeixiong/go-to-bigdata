/*
  For example:
    [vagrant@kubedev-172-17-4-59 go-to-bigdata]$ GOPATH=/Users/fanhongling/Downloads/workspace:/Users/fanhongling/go go test -v -run Net ./pkg/util/netutility/
*/

package netutility

import (
	"testing"
)

func TestNet_interfaces(t *testing.T) {
	SurveyInterfaces()
}
