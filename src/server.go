package main 
import (
	"fmt"
	"net"
	"gob"
	"os"
)

var masterFiles serverFileTree
var mutex sync.Mutex
var 

func main () {
	ln, err := net.Listen("tcp", ":8080") 
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			 continue
		}
		go handleConnection
	}
}
