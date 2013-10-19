package main

import (
	"fmt"
	"net"
	"os"
	"encoding/gob"
	"time"
)

type connecInfo struct {
	ipaddr string //port and addr
}

/* file info */

func getPacket (conn net.Conn) {
	decode = gob.NewDecoder(conn)
	id := 0

	var p packet
	decode.Decode(&p)
	p.handleInfo()

}

func idPacket(id int, pkt packet) {
	if id == 1 {
		packet.handleInfo()
	}
}
