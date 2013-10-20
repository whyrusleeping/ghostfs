package main
import (
	"fmt"
)

/* packet type 1, containing client info */
type PktFileDelete struct {
	pktid int
	client_id int
	path string
}

func (p* PktFileDelete) GetPkid() int {
	return p.pktid;
}

func (p* PktFileDelete) Print() {
	fmt.Println("I'm a PktFileDelete struct.")
}
