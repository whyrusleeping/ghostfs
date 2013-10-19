package main
import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func handshake(conn net.Conn){
	fmt.Fprintf(conn, "swagfs\n")
	reader := bufio.NewReader(conn)
	response, _ := reader.ReadString('\n')

	if response != "hashtag" {
		fmt.Println("We have connected to somebody that isn't our server! Exiting...")
		return
	}
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

	handshake(conn);

	buffer := make([]byte, 1024)

	for {
		conn.Read(buffer)

		fmt.Printf(string(buffer))

	}
}

