package main

/* packet type 1, containing client info */
type PktFileUpdate struct {
	pktid int
	path string
}

func (p* PktFileUpdate) GetPkid() int {
	return p.pktid;
}
