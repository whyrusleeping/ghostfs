package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"bufio"
	"strconv"
)

func handshake(conn net.Conn) (int, error) {
	fmt.Fprintf(conn, "swagfs\n")
	reader := bufio.NewReader(conn)
	response, _ := reader.ReadString('\n')

	response = response[:len(response)-1]
	if response != "hashtag" {
		fmt.Println("[FAIL]")
		//fmt.Println(response)
		//fmt.Println([]byte(response))
		return 0, errors.New("We have connected to somebody that isn't our server! Exiting...")
	}
	fmt.Println("Done")
	response, _ = reader.ReadString('\n')
	response = response[:len(response)-1]
	id,_ = strconv.Atoi(response); 
	return id, nil
}

var clients []cclient
var id int
func main() {
	fmt.Println("")

	if len(os.Args) < 3 {
		fmt.Println("%s : server port folder", os.Args[0])
		fmt.Println("Not enough args, exiting")
		return
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]))

	if err != nil {
		fmt.Println("Connection to server failed. Exiting...")
		return
	}

	fmt.Printf("%-30s", "Handshake...");
	id, err = handshake(conn);
	fmt.Println("Id: ", id);
	if err != nil	{
		fmt.Println(err);
		return
	}

	fmt.Printf("[OK]\n");

	for {
		GetPacket(conn)
	}
}

