package sfs

/* packet interface */
type Packet interface {
	GetPkid() int
	GetClientId() int
	Handle()
	Print()
}
