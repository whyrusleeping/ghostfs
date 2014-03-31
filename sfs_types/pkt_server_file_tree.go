package types
import (
	"fmt"
)

/* packet type 1, containing client info */
type PktServerFileTree struct {
	Pktid int
	Client_id int
	//Sft ServerFileTree
	Files []File
}

func (p* PktServerFileTree) GetPkid() int {
	return p.Pktid;
}

func (p* PktServerFileTree) GetClientId() int {
	return p.Client_id;
}

func (p* PktServerFileTree) Print() {
	fmt.Println("I'm a PktServerFileTree struct.")
}
