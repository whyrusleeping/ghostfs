package main

import (
	"flag"
	"log"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Usage:\n  hello MOUNTPOINT")
	}
	swag := &swagfs{FileSystem: pathfs.NewDefaultFileSystem()}
	swag.Root = new(Dir)
	swag.Root.name = "root"
	swag.Root.attr = new(fuse.Attr)
	swag.Root.Entries = append(swag.Root.Entries, MakeFile("coolfile"))
	mydir := MakeDir("adir")
	mydir.Entries = append(mydir.Entries, MakeFile("lamefile"))
	swag.Root.Entries = append(swag.Root.Entries, mydir)
	nfs := pathfs.NewPathNodeFs(swag, nil)
	server, _, err := nodefs.MountFileSystem(flag.Arg(0), nfs, nil)
	if err != nil {
		log.Fatalf("Mount fail: %s\n", err)
	}
	server.Serve()
}
