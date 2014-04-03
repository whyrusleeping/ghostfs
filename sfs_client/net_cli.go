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

type SfsCli struct {
	Callbacks map[string]chan struct{}
	Outgoing chan sfs.Message
	Incoming chan sfs.Message
	DirRequests chan *dirInfoCallback

	Enc *gob.Encoder
	Dec *gob.Decoder

	ss *SwagSystem
}

type SwagSystem struct {
	fs *swagfs
}

type dirInfoCallback struct {
	Path string
	Reply chan struct{}
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
		s.Incoming <- m
	}
}

func (s *SfsCli) SyncChan() {
	for {
		select {
		case dir := <-s.DirRequests:
			s.Callbacks[dir.Path] = dir.Reply
			drm := new(sfs.DirInfoRequest)
			drm.Path = dir.Path
			go func() {s.Outgoing <- drm}()
		case out := <-s.Outgoing:
			err := s.Enc.Encode(out)
			if err != nil {
				fmt.Println(err)
			}
		case in := <-s.Incoming:
			switch m := in.(type) {
				case sfs.DirInfoMessage:
					fmt.Println("DirInfoMessage:")
					e := s.ss.fs.GetEntry(m.RelPath)
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
}

func (s *SfsCli) RequestDirInfo(path string) {
	resp := make(chan struct{})
	dir := new(dirInfoCallback)
	dir.Path = path
	dir.Reply = resp
	s.DirRequests <- dir
	<-resp
}
