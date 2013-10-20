package main
import (
	"fmt"
)
/* packet type 1, containing client info */
type PktFileRequestChunk struct {
	pktid int
	client_id int
	path string
	chunk int
}

func (p* PktFileRequestChunk) GetPkid() int {
	return p.pktid;
}

func (p* PktFileRequestChunk) Print() {
	fmt.Println("I'm a PktFileRequestChunk struct.")
}
