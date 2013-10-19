package main 
import (
	//"fmt"
	"net"
	"sync"
//	"encoding/gob"
//	"os"
)

type ServerFileTree struct {
	files []file
}

var masterFiles ServerFileTree
var mutex sync.Mutex
var pkt chan Packet

func main () {
	ln, _ := net.Listen("tcp", ":8080") 
	
	go handleIncomingPkts()
	for {
		conn, err := ln.Accept()
		if err != nil {
			 continue
		}
		go handleConnection(conn)
	}
}
