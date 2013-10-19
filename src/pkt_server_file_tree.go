package main

/* packet type 1, containing client info */
type PktServerFileTree struct {
	pktid int
	ipaddr string // addr:port
}

func (p* PktServerFileTree) GetPkid() int {
	return p.pktid;
}
