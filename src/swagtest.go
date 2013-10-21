package main

import (
	"flag"
	"log"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

//Example program to run our fuse filesystem
func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Usage:\n  hello MOUNTPOINT")
	}

	//Make our filesystem structure
	swag := MakeSwag()

	//Use it to create a file system interface
	nfs := pathfs.NewPathNodeFs(swag, nil)

	//Mount our filesystem
	server, _, err := nodefs.MountFileSystem(flag.Arg(0), nfs, nil)
	if err != nil {
		log.Fatalf("Mount fail: %s\n", err)
	}
	server.Serve()
}
