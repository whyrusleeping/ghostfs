package main 
import (
	"fmt"
	"net"
	"sync"
//	"encoding/gob"
	"os"
)


var masterFiles ServerFileTree
var mutex sync.Mutex
var pkt chan Packet
var count int
var clients []client

func main () {
	rootpath := os.Args[1]
	ln, _ := net.Listen("tcp", ":8080")
	count = 1;
	mutex.Lock()
	masterFiles := TraverseDir(rootpath)
	mutex.Unlock()

	for i:=0; i<len(masterFiles.files); i++ {
		fmt.Println(masterFiles.files[i])
	}

	go handleIncomingPkts()

	for {
		conn, err := ln.Accept()
		if err != nil {
			 continue
		}
		go handleConnection(conn)
	}
}
