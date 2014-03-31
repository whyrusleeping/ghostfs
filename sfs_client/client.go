package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"bufio"
	"strconv"
	"github.com/whyrusleeping/swagfs/sfs_types"
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

var clients []types.Cclient
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

	fmt.Printf("%-30s", "Handshake...");
	id, err = handshake(conn);
	if err != nil	{
		fmt.Println(err);
		return
	}

	fmt.Println("Id: ", id);
	fmt.Printf("[OK]\n");

	//Get and handle packets
	for {
		types.GetPacket(conn)
	}
}

