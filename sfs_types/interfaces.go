package types

/* packet interface */
type Packet interface {
	GetPkid() int
	GetClientId() int
	Handle()
	Print()
}
