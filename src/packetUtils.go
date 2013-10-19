package main

import (
	"net"
	"encoding/gob"
)

type connecInfo struct {
	ipaddr string //port and addr
}

/* file info */

func GetPacket (conn net.Conn) {
	decode := gob.NewDecoder(conn)
	var p Packet
	decode.Decode(&p)
	p.Print()
	p.Handle()
}
