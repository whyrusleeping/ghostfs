package main

import (
	"os"
	"fmt"
	"path"
	"strings"
	"github.com/whyrusleeping/ghostfs/gfs_types"
	"github.com/hanwen/go-fuse/fuse"
)

type Node struct {
	Entries []*Node
	Name string
	Attr fuse.Attr
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

func (n *Node) GetDirInfo() *gfs.DirInfo {
	di := new(gfs.DirInfo)
	di.Attr = n.Attr
	for _,e := range n.Entries {
		di.Entries = append(di.Entries, &gfs.EntryInfo{e.Name, e.Attr})
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
		newn.Attr = *fuse.ToAttr(e)
		fmt.Printf("Found: %s, mode: %x\n", e.Name(), e.Mode())
		if e.IsDir() {
			newn.BuildTree(abs_path)
		}
		n.Entries = append(n.Entries, newn)
	}
	return nil
}
