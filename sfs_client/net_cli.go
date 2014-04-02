package main

import (
	"fmt"
	"net"
	"encoding/gob"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/whyrusleeping/swagfs/sfs_types"
	"log"
)

func init() {
	fmt.Println("Initializing gob types.")
	gob.Register(sfs.DirInfoMessage{})
	gob.Register(&sfs.EntryInfo{})
}

type SfsCli struct {

}

type SwagSystem struct {
	fs *swagfs
}


func (ss *SwagSystem) BuildAndMount(mountpoint string) error {
	//Make our filesystem structure
	swag := MakeSwag()
	ss.fs = swag

	//Use it to create a file system interface
	nfs := pathfs.NewPathNodeFs(swag, nil)

	//Mount our filesystem
	server, _, err := nodefs.MountRoot(mountpoint, nfs.Root(), nil)
	if err != nil {
		log.Fatalf("Mount fail: %s\n", err)
		return err
	}
	go server.Serve()
	return nil
}

func (s *SfsCli) Start(host, mount string) error {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println("Connection to server failed. Exiting...")
		return err
	}

	/*
	fmt.Printf("%-30s", "Handshake...");
	id, err = handshake(conn);
	if err != nil	{
		fmt.Println(err);
		return
	}

	fmt.Println("Id: ", id);
	*/
	fmt.Printf("[OK]\n");

	ss := new(SwagSystem)
	ss.BuildAndMount(mount)

	//Get and handle packets
	dec := gob.NewDecoder(conn)
	var m sfs.Message
	for {
		fmt.Println("Wait for message...")
		err := dec.Decode(&m)
		if err != nil {
			panic(err)
		}
		fmt.Println(m)
		switch m := m.(type) {
			case sfs.DirInfoMessage:
				fmt.Println("DirInfoMessage:")
				e := ss.fs.GetEntry(m.RelPath)
				dir,ok := e.(*Dir)
				if !ok {
					fmt.Println("Recieved Dir info for non dir...")
				} else {
					for _,d := range m.Inf.Entries {
						fmt.Println(d.Name)
						dir.AddEntry(MakeEntry(d))
					}
					dir.Loaded = true
				}
			default:
				fmt.Println("Unknown Type.")
		}
	}
}
