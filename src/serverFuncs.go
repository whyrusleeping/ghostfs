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
	configFile, _ := os.Open(".serverconfig")
	dec := json.NewDecoder(configFile)
	dec.Decode(&masterFiles)
	configFile.Close()
}

func saveMasterFiles () {
	configFile, _ := os.Open(".serverconfig")
	enc := json.NewEncoder(configFile)
	enc.Encode(&masterFiles)
	configFile.Close()
}

func handleConnection(conn net.Conn) {
	read := bufio.NewReader(conn)
	str, _ := read.ReadString('\n')
	fmt.Println("Handshake...");
	str = str[:len(str)-1]
	if str != "swagfs" {
		//fmt.Println([]byte(str))
		fmt.Println("Not our client. Disconnecting unkown connection.")
		conn.Close()
		return
	}
	fmt.Fprintf(conn, "hashtag\n")

	mutex.Lock()
	count++
	fmt.Fprintf(conn, "%d\n", count)

	clients = append(clients, client{conn, count})
	mutex.Unlock()

	enc	:= json.NewEncoder(conn)
	//mutex.Lock()

	sft := PktServerFileTree{PKT_SERVER_FILE_TREE, 0, masterFiles}
	//mutex.Unlock()
	enc.Encode(sft)
	fmt.Println("Done")
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
		BroadcastToAll(p.GetClientId(), p)
		p.Print()
	}
}

func BroadcastToAll(id int, p Packet) {
	i := 0
	var toRemove []int
	for i = 0; i < len(clients); i++ {
		if clients[i].id != id {
			enc := gob.NewEncoder(clients[i].conn)
			err := enc.Encode(p)
			if err != nil { //list who has disconnected
				toRemove = append(toRemove, i)
			}
		}
	}

	for i = len(toRemove); i >= 0; i-- {
		clients = append(clients[:i], clients[i + 1:]...)
	}
}


