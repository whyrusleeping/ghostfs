package main
import (
	"fmt"
)

/* packet type 1, containing client info */
type PktServerFileTree struct {
	pktid int
	client_id int
	sft ServerFileTree
}

func (p* PktServerFileTree) GetPkid() int {
	return p.pktid;
}

func (p* PktServerFileTree) GetClientId() int {
	return p.client_id;
}

func (p* PktServerFileTree) Print() {
	fmt.Println("I'm a PktServerFileTree struct.")
}
