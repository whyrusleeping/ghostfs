package main

/* packet interface */
type Packet interface {
	GetPkid() int
	GetClientId() int
	Handle()
	Print()
}
