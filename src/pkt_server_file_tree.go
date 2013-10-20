package main
import (
	"fmt"
)

/* packet type 1, containing client info */
type PktServerFileTree struct {
	pktid int
	client_id int
	ipaddr string // addr:port
}

func (p* PktServerFileTree) GetPkid() int {
	return p.pktid;
}

func (p* PktServerFileTree) Print() {
	fmt.Println("I'm a PktServerFileTree struct.")
}
