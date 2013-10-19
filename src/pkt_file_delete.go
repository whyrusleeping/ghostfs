package main

/* packet type 1, containing client info */
type PktFileDelete struct {
	pktid int
	path string
}

func (p* PktFileDelete) GetPkid() int {
	return p.pktid;
}
