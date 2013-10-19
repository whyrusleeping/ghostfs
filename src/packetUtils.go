package main

import (
	"net"
	"encoding/gob"
)

type connecInfo struct {
	ipaddr string //port and addr
}

/* file info */

func getPacket (conn net.Conn) {
	decode := gob.NewDecoder(conn)

	var p Packet
	decode.Decode(&p)
	p.HandleInfo()

}

func idPacket(id int, pkt Packet) {
	if id == 1 {
		pkt.HandleInfo()
	}
}
