package types

import (
	"os"
	"path"
	"time"
	"net"
)

type File struct {
	Path  string
	Hash  string
	Mtime time.Time
}

type ServerFileTree struct {
	Files []File
}

type Node struct {
	Entries []*Node
	Type os.FileMode
	Name string
}

type Client struct {
	Conn net.Conn
	Id int
}

type Cclient struct {
	Conn net.Conn
	Id int
	Live bool
	Addr string
}

func (n *Node) BuildTree(rel string) error {
	abs_path := path.Join(rel, n.Name)
	dir,err := os.Open(abs_path)
	if err != nil {
		return err
	}
	dinfo,err := dir.Readdir(0)
	if err != nil {
		return err
	}
	for _,e := range dinfo {
		n := new(Node)
		n.Name = e.Name()
		n.Type = e.Mode()
		if e.IsDir() {
			n.BuildTree(abs_path)
		}
	}
	return nil
}
