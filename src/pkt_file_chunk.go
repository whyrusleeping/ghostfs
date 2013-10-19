package main

import (
	"fmt"
)
/* packet type 1, containing client info */
type PktFileChunk struct {
	pktid int
	path string
	chunk []byte
}

func (p* PktFileChunk) GetPkid() int {
	return p.pktid;
}

func (p* PktFileChunk) Print() {
	fmt.Println("I'm a PktFileChunk struct.")
}
