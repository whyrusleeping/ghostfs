package main

import (
	"fmt"
	"net"
	"os"
)

type Client struct {
	Name string
	Conn net.Conn
}

func main() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}
	fmt.Println("Available IP")
	for _, a := range addrs {
		fmt.Println(a)
	}

	clientList := list.New()
	in := make(chan string)
	go IOHandler(in, clientList)

	service := ":7789"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)

	if error != nil {
		Log("Error: Could not resolve address")
		return
	}
	netListen, err := net.Listen(tcpAddr.Network(), tcpAddr.String())
	if err != nil {
		Log(err)
	} else {
		defer netListen.Close()
        for {
            Log("Waiting for Clients")
            connection, err := netListen.Accept()
            if err != nil {
                Log("CLient error: ", err)
            } else {
                go ClientHandler(connection, in, clientList)
            }
        }
	}

}
