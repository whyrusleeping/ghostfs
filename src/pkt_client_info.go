package main
import (
	"fmt"
)

/* packet type 1, containing client info */
type PktClientInfo struct {
	pktid int
	client_id int
	ipaddr string // addr:port
}

func (p* PktClientInfo) GetPkid() int {
	return p.pktid;
}

func (p* PktClientInfo) GetClientId() int {
	return p.client_id;
}

func (p* PktClientInfo) Print() {
	fmt.Println("I'm a PktClientInfo struct.")
}

