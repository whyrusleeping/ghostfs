package main

/* packet type 1, containing client info */
type PktFileRequestChunkMap struct {
	pktid int
	path string
}

func (p* PktFileRequestChunkMap) GetPkid() int {
	return p.pktid;
}
