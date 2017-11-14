/*
   ok
   https://groups.google.com/forum/#!topic/golang-nuts/-L5wDGwkMMI
*/
package main

import (
	"fmt"
	"net"
	"regexp"
)

func main() {
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		fmt.Println(inter.Name, inter.HardwareAddr)
		if addrs, err := inter.Addrs(); err == nil {
			for _, addr := range addrs {
				if inter.Flags&net.FlagLoopback != 0 {
					fmt.Println(inter.Name, "->", addr, "loopback")
					continue // loopback interface
				}
				if ok, err := regexp.MatchString(`\d+\.\d+\.\d+\.\d+[:/]\d+`, addr.String()); err == nil && ok {
					if ipv4Addr, ipv4Net, err := net.ParseCIDR(addr.String()); err == nil {
						fmt.Println(inter.Name, "->", ipv4Addr, ipv4Net, "ip-v4")
						continue
					}
				}
				fmt.Println(inter.Name, "->", addr)
			}
		}
	}
}
