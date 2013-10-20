package main
import (
	"net"
	"bufio"
	"encoding/gob"
	//"encoding/json"
	"os"
	//"sync"
	"fmt"
)

func buildMasterFiles () {
	configFile, _ := os.Open(".serverconfig")
	dec := gob.NewDecoder(configFile)
	dec.Decode(&MasterFiles)
	configFile.Close()
}

func saveMasterFiles () {
	configFile, _ := os.Open(".serverconfig")
	enc := gob.NewEncoder(configFile)
	enc.Encode(&MasterFiles)
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

	enc	:= gob.NewEncoder(conn)
	//mutex.Lock()
/*
	for i:=0; i<len(MasterFiles.Files); i++ {
		fmt.Println(MasterFiles.Files[i])
	}
	*/
	sft := &PktServerFileTree{PKT_SERVER_FILE_TREE, 0, MasterFiles.Files}
	fmt.Println(sft)
	var pp Packet
	pp = sft;
	pp.Print();
	//mutex.Unlock()
	fmt.Fprintf(conn,"%d\n", PKT_SERVER_FILE_TREE);
	enc.Encode(sft)
	fmt.Println("Done")
	dec := gob.NewDecoder(conn)

	BroadcastToAll(count, &PktClientInfo{PKT_CLIENT_INFO, count, conn.LocalAddr().String()})

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
	fmt.Println("Count: ", len(clients));
	for i = 0; i < len(clients); i++ {
		if clients[i].id != id {
			fmt.Println("Sending to client ", i)
			enc := gob.NewEncoder(clients[i].conn)
			err := enc.Encode(p)
			if err != nil { //list who has disconnected
				fmt.Println("ERROR, KILL THE CLIENT!")
				fmt.Println(err)
				toRemove = append(toRemove, i)
			}
		}
	}

	for i = len(toRemove)-1; i >= 0; i-- {
		fmt.Println("RMEOVE")
		clients = append(clients[:i], clients[i + 1:]...)
	}
}


