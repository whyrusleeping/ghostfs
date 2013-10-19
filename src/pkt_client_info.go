package main

/* packet type 1, containing client info */
type PktClientInfo struct {
	pktid int
	ipaddr string // addr:port
}

func (p* PktClientInfo) GetPkid() int {
	return p.pktid;
}
