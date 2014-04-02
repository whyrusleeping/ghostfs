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
		fmt.Printf("usage: %s server folder\n", os.Args[0])
		fmt.Println("Not enough args, exiting")
		return
	}

	cli := new(SfsCli)
	cli.Start(os.Args[1], os.Args[2])
}

