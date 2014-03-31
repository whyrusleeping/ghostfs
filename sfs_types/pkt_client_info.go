package types
import (
	"fmt"
)

/* packet type 1, containing client info */
type PktClientInfo struct {
	Pktid int
	Client_id int
	Ipaddr string // addr:port
}

func (p* PktClientInfo) GetPkid() int {
	return p.Pktid;
}

func (p* PktClientInfo) GetClientId() int {
	return p.Client_id;
}

func (p* PktClientInfo) Print() {
	fmt.Println("I'm a PktClientInfo struct.")
}

