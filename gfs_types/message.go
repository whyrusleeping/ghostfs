package gfs

import (
	"encoding/gob"
	
	"github.com/hanwen/go-fuse/fuse"
)

func init() {
	gob.Register(DirInfoMessage{})
	gob.Register(&EntryInfo{})
	gob.Register(&DirInfoRequest{})
}

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

type DirInfoRequest struct {
	Path string
}
