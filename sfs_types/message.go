package sfs

import (
	"github.com/hanwen/go-fuse/fuse"
)

type EntryInfo struct {
	Name string
	Attr fuse.Attr
}

type DirInfo struct {
	Entries []*EntryInfo
	Attr fuse.Attr
}

type Message interface {

}

type DirInfoMessage struct {
	Inf *DirInfo
	RelPath string
}
