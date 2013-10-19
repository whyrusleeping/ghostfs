package main
import (
	"fmt"
)

/* packet type 1, containing client info */
type PktFileCreate struct {
	pktid int
	path string
}

func (p* PktFileCreate) GetPkid() int {
	return p.pktid;
}

func (p* PktFileCreate) Print() {
	fmt.Println("I'm a PktFileCreate struct.")
}
