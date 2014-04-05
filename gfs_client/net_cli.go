package main

import (
	"fmt"
	"net"
	"encoding/gob"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/whyrusleeping/ghostfs/gfs_types"
	"log"
)

type GfsCli struct {
	Callbacks map[string]chan struct{}
	Outgoing chan gfs.Message
	Incoming chan gfs.Message
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


func (ss *SwagSystem) BuildAndMount(mountpoint string, cli *GfsCli) error {
	//Make our filesystem structure
	swag := MakeSwag(cli)
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

func NewGfsCli() *GfsCli {
	s := new(GfsCli)
	s.Callbacks = make(map[string]chan struct{})
	s.Incoming = make(chan gfs.Message)
	s.Outgoing = make(chan gfs.Message)
	s.DirRequests = make(chan *dirInfoCallback)
	return s
}

func (s *GfsCli) Start(host, mount string) error {
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
	ss.BuildAndMount(mount, s)
	s.ss = ss

	//Get and handle packets
	dec := gob.NewDecoder(conn)
	s.Enc = gob.NewEncoder(conn)
	var m gfs.Message
	go s.SyncChan()
	for {
		fmt.Println("Wait for message...")
		err := dec.Decode(&m)
		if err != nil {
			panic(err)
		}
		fmt.Println("Got message...")
		//fmt.Println(m)
		s.Incoming <- m
	}
}

func (s *GfsCli) SyncChan() {
	for {
		select {
		case dir := <-s.DirRequests:
			fmt.Printf("Processing request for: %s\n", dir.Path)
			s.Callbacks[dir.Path] = dir.Reply
			drm := new(gfs.DirInfoRequest)
			drm.Path = dir.Path
			go func() {
				s.Outgoing <- drm
				fmt.Println("Placed request on outgoing queue.")
			}()
		case out := <-s.Outgoing:
			fmt.Println("Sending message to server...")
			err := s.Enc.Encode(&out)
			if err != nil {
				fmt.Println(err)
			}
		case in := <-s.Incoming:
			switch m := in.(type) {
				case *gfs.DirInfoMessage:
					fmt.Printf("DirInfoMessage: '%s'\n", m.RelPath)
					e := s.ss.fs.GetEntry(m.RelPath, false)
					if e == nil {
						fmt.Println("Nil entry returned...")
					}
					dir,ok := e.(*Dir)
					if !ok {
						fmt.Println("Recieved Dir info for non dir...")
					} else {
						for _,d := range m.Inf.Entries {
							//fmt.Println(d.Name)
							dir.AddEntry(MakeEntry(d))
						}
						dir.Loaded = true

						//If someone was waiting on this info, tell them
						resp,ok := s.Callbacks[m.RelPath]
						if ok {
							if resp == nil {
								fmt.Printf("Nil channel for path: %s\n", m.RelPath)
							}
							delete(s.Callbacks, m.RelPath)
							fmt.Println("Replying to callback!")
							go func() {
								resp <- struct{}{}
							}()
						}
					}
				default:
					fmt.Println("Unknown Type.")
			}
		}
	}
}

func (s *GfsCli) RequestDirInfo(path string) {
	fmt.Printf("Requesting: %s\n", path)
	resp := make(chan struct{})
	dir := new(dirInfoCallback)
	dir.Path = path
	dir.Reply = resp
	s.DirRequests <- dir
	<-resp
	fmt.Printf("Request for '%s' completed!\n", path)
}
