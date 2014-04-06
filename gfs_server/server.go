package main

import (
	"fmt"
	"net"
	"reflect"
	"os"
	"encoding/gob"
	"github.com/whyrusleeping/ghostfs/gfs_types"
	"crypto/md5"
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

func HandleFileDataRequest(fdr *gfs.FileDataRequest) {
	fi,err := os.Open(fdr.Path)
	resp := new(gfs.FileDataResponse)
	defer func() {
		fi.Close()
		fdr.GetCallback() <- resp
	}()
	if err != nil {
		resp.Error = err.Error()
		return
	}
	_,err = fi.Seek(fdr.Offset, os.SEEK_SET)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	data := make([]byte, fdr.Size)
	_,err = fi.Read(data)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	resp.Path = fdr.Path
	resp.Data = data
	resp.Hash = md5.Sum(data)
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
							dim := new(gfs.DirInfoMessage)
							dim.Inf = ent.GetDirInfo()
							dim.RelPath = in.Path
							s.Broadcast <- dim
							fmt.Println("Broadcasted dirinfo message.")
						}()
					case *gfs.FileDataRequest:
						go HandleFileDataRequest(in)
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
	rootinf := new(gfs.DirInfoMessage)
	rootinf.Inf = s.TreeRoot.GetDirInfo()
	rootinf.RelPath = ""
	cl.SendMessage(rootinf)
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
