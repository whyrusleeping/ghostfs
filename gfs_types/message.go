package gfs

import (
	"fmt"
	"encoding/gob"
	"github.com/hanwen/go-fuse/fuse"
)

func init() {
	fmt.Println("Registering gob types.")
	gob.Register(&DirInfoMessage{})
	gob.Register(&EntryInfo{})
	gob.Register(&DirInfoRequest{})
	gob.Register(&FileDataRequest{})
	gob.Register(&FileDataResponse{})
}

type EntryInfo struct {
	Name string
	Attr fuse.Attr
}

type DirInfo struct {
	Entries []*EntryInfo
	Attr fuse.Attr
}

type callback struct {}
func (c *callback) SetCallback(chan Message) {}
func (c *callback) GetCallback() chan Message {return nil}

type Message interface {
	SetCallback(chan Message)
	GetCallback() chan Message
}

type Request interface {
	SetCallback(chan Message)
	GetCallback() chan Message
}

type DirInfoMessage struct {
	Inf *DirInfo
	RelPath string
	callback
}

type DirInfoRequest struct {
	Path string
	callback
}

type FileDataRequest struct {
	Path string
	Offset int64
	Size int
	callback chan Message
}

func (f *FileDataRequest) SetCallback(c chan Message) {
	f.callback = c
}

func (f *FileDataRequest) GetCallback() chan Message {
	return f.callback
}

type FileDataResponse struct {
	Path string
	Hash [16]byte
	Data []byte
	Error string
	callback
}
