package main

import (
	"os"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"fmt"
	"github.com/whyrusleeping/ghostfs/gfs_types"
)

//Everything in our filesystem is an 'Entry'
type Entry interface {
	Name() string
	Attr() *fuse.Attr
	GetInfo() fuse.DirEntry
}

//Embeddable struct to ease of coding
type mEntry struct {
	attr *fuse.Attr
	name string
}

type NotLoadedError struct {}
func (n NotLoadedError) Error() string {
	return "Tryed to access unloaded directory."
}

func MakeEntry(e *gfs.EntryInfo) Entry {
	if (e.Attr.Mode & uint32(os.ModeDir)) > 0 {
		d := new(Dir)
		d.name = e.Name
		d.Entries = make(map[string]Entry)
		d.attr = &e.Attr
		return d
	}
	f := new(File)
	f.name = e.Name
	f.attr = &e.Attr
	return f
}

//Directory
type Dir struct {
	//Entries []Entry
	Entries map[string]Entry
	Loaded bool
	mEntry
}

func MakeDir(name string) *Dir {
	d := new(Dir)
	d.name = name
	d.attr = &fuse.Attr{Mode: fuse.S_IFDIR | 0755}
	d.Entries = make(map[string]Entry)
	return d
}

func (d *Dir) AddEntry(e Entry) {
	if e == nil {
		return
	}
	//d.Entries = append(d.Entries, e)
	d.Entries[e.Name()] = e
	d.attr.Nlink = uint32(len(d.Entries))
}

func (d *Dir) RemoveChild(name string) {
	delete(d.Entries, name)
	/*
	for i, e := range d.Entries {
		if name == e.Name() {
			d.Entries = append(d.Entries[:i], d.Entries[i+1:]...)
			return
		}
	}
	*/
}

func (d *Dir) GetEntry(toks []string) (Entry,error) {
	if len(toks) == 0 {
		fmt.Println("returning self")
		return d,nil
	}

	e,ok := d.Entries[toks[0]]
	if !ok {
		return nil,nil
	}
	if len(toks) == 1 {
		return e,nil
	}
	sub, ok := e.(*Dir)
	if !ok {
		//Was not a dir
		return nil,nil
	}
	if !sub.Loaded {
		return nil,NotLoadedError{}
	}
	return sub.GetEntry(toks[1:])
	/*
	for _,e := range(d.Entries) {
		if e.Name() == toks[0] {
			if len(toks) == 1 {
				//This is it!
				return e
			} else {
				//Need to search deeper
				sub, ok := e.(*Dir)
				if !ok {
					return nil
				}
				return sub.GetEntry(toks[1:])
			}
		}
	}
	return nil
	*/
}

func (d *Dir) GetInfo() fuse.DirEntry {
	fmt.Printf("Getinfo called, mode = %d\n", d.attr.Mode)
	return fuse.DirEntry{Name: d.name, Mode: d.attr.Mode}
}

func (d *Dir) Name() string {
	return d.name
}

func (d *Dir) Attr() *fuse.Attr {
	return d.attr
}

//Normal file
type File struct {
	Content string
	//TODO: Implement custom nodefs.File
	FileData nodefs.File
	Chunks int
	RealSize int //Full size of the file
	LocalSize int //Size of this instance on disk
	mEntry
}

func MakeFile(name string) *File {
	f := new(File)
	f.name = name
	f.attr = &fuse.Attr{ Mode: fuse.S_IFREG | 0644, Size: uint64(len(name))}
	f.FileData = nodefs.NewDefaultFile()
	return f
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Attr() *fuse.Attr {
	return f.attr
}

func (f *File) GetInfo() fuse.DirEntry {
	fmt.Printf("Calling getinfo, mode= %d\n", f.attr.Mode)
	return fuse.DirEntry{Name: f.name, Mode: f.attr.Mode}
}

//Link to a file or directory (or another link... interesting...)
type Link struct {
	To Entry
	Lname string
}

func (l *Link) Name() string {
	return l.Lname
}

func (l *Link) Attr() *fuse.Attr {
	return l.To.Attr()
}
