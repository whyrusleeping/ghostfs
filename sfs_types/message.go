package sfs

import (
	"os"
)

type EntryInfo struct {
	Name string
	Mode os.FileMode
}

type DirInfo struct {
	Entries []*EntryInfo
	Mode os.FileMode
}

type Message interface {

}

type DirInfoMessage struct {
	Inf *DirInfo
}
