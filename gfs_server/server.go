package main

import (
	"fmt"
	"net"
	"reflect"
	"os"
	"encoding/gob"
	"github.com/whyrusleeping/ghostfs/gfs_types"
)

type SfsServer struct {
	TreeRoot *Node
	Root string
	Clients []*Client

	NewClients chan *Client
	Broadcast chan gfs.Message
	Incoming chan gfs.Message
}

func NewServer(root string) *SfsServer {
	ss := new(SfsServer)
	ss.Root = root
	ss.TreeRoot = new(Node)
	ss.TreeRoot.BuildTree(root)
	ss.Incoming = make(chan gfs.Message)
	ss.NewClients = make(chan *Client)
	ss.Broadcast = make(chan gfs.Message)
	return ss
}

func (s *SfsServer) ServeSwag(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			 continue
		}
		go s.AddClient(conn)
	}
}

func (s *SfsServer) SyncChan() {
	for {
		select {
			case nc := <-s.NewClients:
				s.Clients = append(s.Clients, nc)
			case mes := <-s.Broadcast:
				for _,c := range s.Clients {
					c.OutGoing <-mes
				}
			case in := <-s.Incoming:
				switch in := in.(type) {
					case *gfs.DirInfoRequest:
						fmt.Printf("dir requested: %s\n", in.Path)
						ent := s.TreeRoot.Find(in.Path)
						go func() {
							s.Broadcast <- &gfs.DirInfoMessage{ent.GetDirInfo(),in.Path}
							fmt.Println("Broadcasted dirinfo message.")
						}()
					default:
						fmt.Println("Unrecognized message type...")
						fmt.Println(reflect.TypeOf(in))
				}
		}
	}
}

func (s *SfsServer) AddClient(c net.Conn) {
	cl := s.NewClient(c)
	go cl.Start()
	cl.SendMessage(&gfs.DirInfoMessage{s.TreeRoot.GetDirInfo(),""})
	s.NewClients <- cl
}

func (s *SfsServer) NewClient(c net.Conn) *Client {
	cl := new(Client)
	cl.Con = c
	cl.Dec = gob.NewDecoder(c)
	cl.Enc = gob.NewEncoder(c)
	cl.OutGoing = make(chan gfs.Message)
	cl.ServCom = s.Incoming
	return cl
}

func main () {
	if len(os.Args) < 2 {
		fmt.Println("Must specify root path")
		return
	}
	rootpath := os.Args[1]
	s := NewServer(rootpath)
	go s.SyncChan()
	s.ServeSwag(":8080")
}
