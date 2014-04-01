package main

import (
	"os"
	"fmt"
	"path"
	"strings"
	"github.com/whyrusleeping/swagfs/sfs_types"
)

type Node struct {
	Entries []*Node
	Mode os.FileMode
	Name string
}

func (n *Node) Child(name string) *Node {
	for _,c := range n.Entries {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func (n *Node) Find(path string) *Node {
	parts := strings.Split(path, "/")
	cur := n
	for _,i := range parts {
		cur = cur.Child(i)
		if cur == nil {
			return nil
		}
	}
	return cur
}

func (n *Node) GetDirInfo() *sfs.DirInfo {
	di := new(sfs.DirInfo)
	di.Mode = n.Mode
	for _,e := range n.Entries {
		di.Entries = append(di.Entries, &sfs.EntryInfo{e.Name, e.Mode})
	}
	fmt.Printf("Dir Info: %d items\n", len(di.Entries))
	return di
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
		newn := new(Node)
		newn.Name = e.Name()
		newn.Mode = e.Mode()
		fmt.Printf("Found: %s\n", e.Name())
		if e.IsDir() {
			newn.BuildTree(abs_path)
		}
		n.Entries = append(n.Entries, newn)
	}
	return nil
}
