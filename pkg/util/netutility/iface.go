package netutility

import (
	"fmt"
	"net"
	// "os"
)

func SurveyInterfaces() {
	// Get a list of all interfaces on the machine.
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Unable to get interfaces.")
	}

	// Loop over all interfaces
	for i := range interfaces {

		// Get the interfaces Index and then load
		// the interface by Index.
		id := interfaces[i].Index
		iface, err := net.InterfaceByIndex(id)

		if err != nil {
			panic("Unable to get interface by id")
		}

		fmt.Printf("Index: %v\nName: %+v\nHWAddr: %+v\nFlags: %+v\n", id, interfaces[i].Name, iface.HardwareAddr, iface.Flags.String())

		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}

		fmt.Println("Unicasts")
		for _, v := range addrs {
			fmt.Println(v)
		}

		multicasts, err := iface.MulticastAddrs()
		if err != nil {
			panic(err)
		}

		fmt.Println("Multicasts")
		for _, v := range multicasts {
			fmt.Println(v)
		}

		fmt.Println()
	}

}
