package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"bufio"
	"strconv"
	"encoding/gob"
	"github.com/whyrusleeping/swagfs/sfs_types"
	"log"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	//"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

func handshake(conn net.Conn) (int, error) {
	fmt.Fprintf(conn, "swagfs\n")
	reader := bufio.NewReader(conn)
	response, _ := reader.ReadString('\n')

	response = response[:len(response)-1]
	if response != "hashtag" {
		fmt.Println("[FAIL]")
		return 0, errors.New("We have connected to somebody that isn't our server! Exiting...")
	}
	fmt.Println("Done")
	response, _ = reader.ReadString('\n')
	response = response[:len(response)-1]
	id,_ = strconv.Atoi(response)
	return id, nil
}

var id int

func main() {
	fmt.Println("")

	if len(os.Args) < 3 {
		fmt.Printf("usage: %s server port folder\n", os.Args[0])
		fmt.Println("Not enough args, exiting")
		return
	}

	host := fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println("Connection to server failed. Exiting...")
		return
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

	//Make our filesystem structure
	swag := MakeSwag()

	//Use it to create a file system interface
	nfs := pathfs.NewPathNodeFs(swag, nil)

	//Mount our filesystem
	server, _, err := nodefs.MountRoot(os.Args[3], nfs.Root(), nil)
	if err != nil {
		log.Fatalf("Mount fail: %s\n", err)
	}
	go server.Serve()

	//Get and handle packets
	gob.Register(sfs.DirInfoMessage{})
	gob.Register(&sfs.EntryInfo{})
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
				e := swag.GetEntry(m.RelPath)
				dir,ok := e.(*Dir)
				if !ok {
					fmt.Println("Recieved Dir info for non dir...")
				} else {
					for _,d := range m.Inf.Entries {
						fmt.Println(d.Name)
						dir.AddEntry(MakeEntry(d))
					}
				}
			default:
				fmt.Println("Unknown Type.")
		}
	}
}

