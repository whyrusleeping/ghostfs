package main
import (
	"net"
	"bufio"
	"encoding/gob"
	"encoding/json"
	"os"
	"sync"
)

func buildMasterFiles () {
	configFile, err := os.Open("/.serverconfig")
	dec := json.NewDecoder(configFile)
	dec.Decode(&masterFiles)
	configFile.Close()
}

func saveMasterFiles () {
	configFile, err := os.Open("/.serverconfig")
	enc := json.NewEncoder(configFile)
	enc.Encode(&masterFiles)
	configFile.Close()
}

func handleConnection(conn net.conn) {
	read := bufio.NewReader(conn)
	
	str, err := read.ReadString('\n')
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
	
	for ;

}

func updateFile(
	


