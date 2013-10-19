package main

/* packet type 1, containing client info */
type PktFileChunk struct {
	pktid int
	path string
	chunk []byte
}

func (p* PktFileChunk) GetPkid() int {
	return p.pktid;
}
