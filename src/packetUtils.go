package main
import (
	"fmt"
	"net"
	"os"
	"encoding/gob"
)

type connecInfo struct {
	ipaddr string //port and addr
}

/* packet interface */
type packet interface { }

/* packet type 1, containing client info */
type clientPacket struct {
	ipaddr string
}

/* packet type 2, containing info on all files */
type serverFileTree struct {
	files []serverFiles
}

/* file info */
type file struct {
	string modTime
	string hash
	string path
}

func getPacket (conn net.Conn) {
	decode = gob.NewDecoder(conn)
	id := 0
	fmt.Fscanf(conn, "%d", &id)

	var p packet
	switch (id) {
		case 1: 
			decode.Decode(&p)
			p.handleInfo()
	}
}

func idPacket(id int, pkt packet) {
	if id == 1 {
		packet.handleInfo()
	}
}

func (c *clientPacket) 
