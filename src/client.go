package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"bufio"
)

func handshake(conn net.Conn) error{
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

	response, _ := reader.ReadString('\n')
	response = response[:len(response)-1]
	return int(id), nil
}

id int
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

	if err != nil	{
		fmt.Println(err);
		return
	}

	fmt.Printf("[OK]\n");

	for {
		GetPacket(conn)
	}
}

