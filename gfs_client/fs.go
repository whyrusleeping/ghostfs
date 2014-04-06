package main

import (
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"path/filepath"
	"strings"
	"time"
	"fmt"
	"reflect"
)

type swagfs struct {
	pathfs.FileSystem
	Root *Dir
	cachedir string

	ncli *GfsCli
}

func MakeSwag(cli *GfsCli) *swagfs {
	swag := &swagfs{FileSystem: pathfs.NewDefaultFileSystem()}
	swag.Root = MakeDir("")
	swag.ncli = cli
	return swag
}

//Traverse the file tree and get the file specified
func (fs *swagfs) GetEntry(path string, doload bool) Entry {
	//fmt.Printf("calling get entry! %s\n", path)
	if path == "" {
		return fs.Root
	}
	toks := strings.Split(path, "/")

	e,err := fs.Root.GetEntry(toks)
	if err != nil {
		if doload {
			fmt.Printf("Requesting info for: %s\n", path)
			fs.ncli.RequestDirInfo(path)
		}
	}
	if e != nil {
		fmt.Printf("Got Entry: '%s' %s\n", e.Name(), reflect.TypeOf(e))
	}
	return e
}

func (fs *swagfs) Access(path string, mode uint32, context *fuse.Context) fuse.Status {
	return fuse.OK
}
//Require the 'name' to be the full path
func (fs *swagfs) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	//fmt.Printf("calling get attr: %s\n", name)
	if name == "" {
		fmt.Println("Returning Empty.")
		return &fuse.Attr{
			Mode: fuse.S_IFDIR | 0755,
		}, fuse.OK
	}

	e := fs.GetEntry(name, true)
	if e == nil {
		return nil, fuse.ENOENT
	}
	//fmt.Printf("Get Attr:")
	//fmt.Printf("%s: mode: %x size: %d uid: %d\n", e.Name(), e.Attr().Mode,
		//e.Attr().Size, e.Attr().Uid)
	na := &fuse.Attr{Mode: e.Attr().Mode}
	return na, fuse.OK
}

func (fs *swagfs) OpenDir(name string, context *fuse.Context) (c []fuse.DirEntry, code fuse.Status) {
	fmt.Printf("Calling open dir: %s\n", name)
	e := fs.GetEntry(name, true)
	if e == nil {
		fmt.Println("Got nil back.")
	}
	dir, ok := e.(*Dir)
	if !ok {
		fmt.Printf("Failed to open dir...")
		fmt.Println(e)
		return nil, fuse.ENOENT
	}

	var ents []fuse.DirEntry

	for _,sub := range(dir.Entries) {
		ents = append(ents, sub.GetInfo())
	}
	//fmt.Println(ents)
	return ents, fuse.OK
}

//TODO: actually manage file data
func (fs *swagfs) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	fmt.Printf("calling open: %s\n", name)
	e := fs.GetEntry(name, true)
	if e == nil {
		return nil, fuse.ENOENT
	}

	fi, ok := e.(*File)
	if !ok {
		return nil, fuse.ENOENT
	}
	return fi.FileData, fuse.OK
}

func (fs *swagfs) Utimens(name string, Atime *time.Time, Mtime *time.Time, context *fuse.Context) fuse.Status {
	e := fs.GetEntry(name, true)
	if e == nil {
		return fuse.ENOENT
	}
	attr := e.Attr()
	attr.Mtime = uint64(Mtime.Unix())
	attr.Atime = uint64(Atime.Unix())
	fmt.Printf("calling update times: %s\n", name)
	return fuse.OK
}

func (fs *swagfs) Mknod(name string, mode uint32, dev uint32, context *fuse.Context) fuse.Status {
	fmt.Printf("Calling mknod: %s\n", name)
	dir, file := filepath.Split(name)
	pardir := fs.GetEntry(dir, true)
	if pardir == nil {
		return fuse.ENOENT
	}

	pdir_t, ok := pardir.(*Dir)
	if !ok {
		return fuse.ENOENT
	}

	fi := MakeFile(file)
	fi.attr = new(fuse.Attr)
	fi.attr.Mode = mode
	pdir_t.AddEntry(fi)
	return fuse.OK
}
