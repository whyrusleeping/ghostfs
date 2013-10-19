package main
import (
	"time"
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
