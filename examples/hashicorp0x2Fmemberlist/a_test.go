package hashicorp0x2Fmemberlist

import (
	"net"
	"os"
	"sync"
	"testing"
)

func Test_Hostname(t *testing.T) {
	hn, err := os.Hostname()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(hn)
	}
}

var bindLock sync.Mutex
var bindNum byte = 10

func getBindAddr() net.IP {
	bindLock.Lock()
	defer bindLock.Unlock()

	result := net.IPv4(127, 0, 0, bindNum)
	bindNum++
	if bindNum > 255 {
		bindNum = 10
	}

	return result
}

func TestHM_getBindAddr(t *testing.T) {
	t.Log(getBindAddr())
	t.Log(getBindAddr())
}
