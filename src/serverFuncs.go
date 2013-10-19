package main
import (
	"net"
	"bufio"
	"encoding/gob"
	"encoding/json"
	"os"
	//"sync"
	"fmt"
)

func buildMasterFiles () {
	configFile, _ := os.Open("/.serverconfig")
	dec := json.NewDecoder(configFile)
	dec.Decode(&masterFiles)
	configFile.Close()
}

func saveMasterFiles () {
	configFile, _ := os.Open("/.serverconfig")
	enc := json.NewEncoder(configFile)
	enc.Encode(&masterFiles)
	configFile.Close()
}

func handleConnection(conn net.Conn) {
	read := bufio.NewReader(conn)

	str, _ := read.ReadString('\n')
	if str != "swagfs" {
		fmt.Println("Not our client. Disconnecting unkown connection.")
		conn.Close()
		return
	}
	fmt.Fprintf(conn, "hashtag\n")
	enc	:= gob.NewEncoder(conn)
	mutex.Lock()
	enc.Encode(masterFiles)
	mutex.Unlock()
	dec := gob.NewDecoder(conn)

	var p Packet

	for {
		dec.Decode(&p)
		pkt <- p
	}

}

func handleIncomingPkts () {
	var p Packet
	for {
		p = <-pkt
		p.Print()
	}
}


