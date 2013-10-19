package main
import (
	"time"
	"fmt"
)

/* packet type 1, containing client info */
type PktFileUpdate struct {
	pktid int
	path string
	modTime time.Time
	hash []byte
}

func (p* PktFileUpdate) GetPkid() int {
	return p.pktid;
}

func (p* PktFileUpdate) Print() {
	fmt.Println("I'm a PktFileUpdate struct.")
}
