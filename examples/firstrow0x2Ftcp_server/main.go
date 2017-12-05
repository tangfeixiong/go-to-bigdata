package main

import (
	"fmt"

	"github.com/firstrow/tcp_server"
)

func main() {
	ts := tcp_server.New("0.0.0.0:2017")

	ts.OnNewClient(func(c *tcp_server.Client) {
		// new client connected
		// lets send some message
		fmt.Println("Client accepted")
	})
	ts.OnNewMessage(func(c *tcp_server.Client, message string) {
		// new message received
		fmt.Println(message)
		c.Send(message)
	})
	ts.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		// connection with client lost
		fmt.Println("closed, err: ", err)
	})

	ts.Listen()
}
