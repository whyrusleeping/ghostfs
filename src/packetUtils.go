package main

import (
	"net"
	"encoding/json"
	"fmt"
)

type connecInfo struct {
	ipaddr string //port and addr
}

/* file info */

func GetPacket (conn net.Conn) {
/*	b := make([]byte, 1)
	for {
		conn.Read(b)
		fmt.Println(string(b))
	}
	*/
	decode := json.NewDecoder(conn)
	fmt.Println("GetPacket")
	var p Packet
	decode.Decode(&p)
	fmt.Println("Have Packet")
	p.Print()
	p.Handle()
}
