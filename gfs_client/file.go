package main

import (
	"fmt"
	"time"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse"
)

const BlockSize = 1024

type FileChunk struct {
	Data []byte
	Hash []byte
	Index int
}

type GfsFile struct {
	//nodefs.File
	cli *GfsCli
	data []*FileChunk
	size uint32
	perms uint32
	in chan *FileChunk
	path string
}

func (f *GfsFile) Read(dest []byte, off int64) (fuse.ReadResult, fuse.Status) {
	t := []byte("hello")
	copy(dest, t)
	return fuse.ReadResultData(dest), fuse.OK
	fmt.Println("Reading data from file")
	blk_i := off / BlockSize
	blk := f.getBlock(blk_i)
	o := off - (BlockSize * blk_i)
	toread := int64(len(dest))
	thisread := toread
	if toread > (BlockSize - o) {
		thisread = (BlockSize - o)
	}
	copy(dest, blk.Data[:o])
	toread -= thisread
	return fuse.ReadResultData(dest), fuse.OK
}

func (f *GfsFile) Write(data []byte, off int64) (written uint32, code fuse.Status) {
	fmt.Println("Writing data to file.\n")
	return 0, fuse.OK
}

//func (f *GfsFile) 
func (f *GfsFile) Utimens(atime *time.Time, mtime *time.Time) fuse.Status {
	fmt.Println("Called Utimens.")
	return fuse.OK
}

func (f *GfsFile) Truncate(size uint64) fuse.Status {
	fmt.Printf("Called truncate: %d\n", size)
	return fuse.OK
}

func (f *GfsFile) SetInode(ino *nodefs.Inode) {
	fmt.Println("Called SetInode.")
}

func (f *GfsFile) Release() {
	fmt.Println("Called release")
}

func (f *GfsFile) GetAttr(out *fuse.Attr) fuse.Status {
	fmt.Println("Calling GetAttr.")
	return fuse.OK
}

func (f *GfsFile) Fsync(flags int) fuse.Status {
	fmt.Printf("Calling fsync: %d\n", flags)
	return fuse.OK
}

func (f *GfsFile) Chown(uid uint32, gid uint32) fuse.Status {
	fmt.Println("Chown called.")
	return fuse.OK
}

func (f *GfsFile) Flush() fuse.Status {
	fmt.Println("Flush Called.")
	return fuse.OK
}

func (f *GfsFile) getBlock(blk_i int64) *FileChunk {
	if blk_i >= int64(len(f.data)) {
		//Request file data
	}
	blk := f.data[blk_i]
	if blk == nil {
		//Request block
		f.cli.RequestFileData(f.path, int(blk_i), f.in)
		f.data[blk_i] = <-f.in
	}
	return f.data[blk_i]
}


func (f *GfsFile) Chmod(perms uint32) fuse.Status {
	f.perms = perms
	return fuse.OK
}

func (f *GfsFile) Allocate(off, size uint64, mode uint32) fuse.Status {
	if uint64(len(f.data)) * BlockSize < off + size {
		//Need to allocate more blocks
	}
	return fuse.OK
}

func (f *GfsFile) InnerFile() nodefs.File {
	return nil
}

func (f *GfsFile) String() string {
	return "A File!"
}
