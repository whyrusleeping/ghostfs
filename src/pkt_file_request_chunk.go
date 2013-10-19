package main

/* packet type 1, containing client info */
type PktFileRequestChunk struct {
	pktid int
	path string
	chunk int
}

func (p* PktFileRequestChunk) GetPkid() int {
	return p.pktid;
}
