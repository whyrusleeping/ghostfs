package main

import (
	"fmt"
	"net"
	"sync"
	"os"
)

//TODO: maybe put this all in a struct?
var MasterFiles *ServerFileTree
var mutex sync.Mutex //TODO: mutexes are bad, lets somehow use a channel to sync stuff up
var pkt chan Packet
var count int
var clients []client

func main () {
	if len(os.Args) < 2 {
		fmt.Println("Must specify root path")
		return
	}
	rootpath := os.Args[1]

	ln, _ := net.Listen("tcp", ":8080")
	count = 1

	MasterFiles = TraverseDir(rootpath)

	go handleIncomingPkts()

	for {
		conn, err := ln.Accept()
		if err != nil {
			 continue
		}
		go handleConnection(conn)
	}
}
