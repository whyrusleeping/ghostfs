package main

/* packet interface */
type Packet interface {
	GetPkid() int
	Handle()
}

