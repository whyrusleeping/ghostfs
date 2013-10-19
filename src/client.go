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
		fmt.Println(response)
		fmt.Println([]byte(response))
		return errors.New("We have connected to somebody that isn't our server! Exiting...")
	}

	return nil
}

func main() {
	fmt.Println("")

	if len(os.Args) < 2 {
		fmt.Println("Not enough args, exiting")
		return
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]))

	if err != nil {
		fmt.Println("Connection to server failed. Exiting...")
		return
	}

	fmt.Printf("%-30s", "Handshake...");
	err = handshake(conn);

	if err != nil	{
		fmt.Println(err);
		return
	}

	fmt.Printf("[OK]\n");

	for {
		GetPacket(conn)
	}
}

