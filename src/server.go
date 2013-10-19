package main 
import (
	"fmt"
	"net"
	"gob"
	"os"
)

var masterFiles serverFileTree
var mutex sync.Mutex
var pkt chan Packet

func main () {
	ln, err := net.Listen("tcp", ":8080") 
	
	go handleIncomingPkts()
	for {
		conn, err := ln.Accept()
		if err != nil {
			 continue
		}
		go handleConnection()
	}
}
