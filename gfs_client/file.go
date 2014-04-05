package main

import (
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse"
)

type GfsFile struct {
	//nodefs.File
	cli *GfsCli
}

func (f *GfsFile) Read(dest []byte, off int64) (fuse.ReadResult, fuse.Status) {
}

Write(data []byte, off int64) (written uint32, code fuse.Status)

func (f *GfsFile) Chmod(perms uint32) fuse.Status {
	return fuse.OK
}

func (f *GfsFile) Allocate(off, size uint64, mode uint32) fuse.Status {
	return fuse.OK
}

func (f *GfsFile) InnerFile() nodefs.File {
	return nil
}

func (f *GfsFile) String() string {
	return "A File!"
}
