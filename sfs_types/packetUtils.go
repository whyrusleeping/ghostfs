package types

import (
	"net"
	//"encoding/json"
	"encoding/gob"
	"fmt"
	"bufio"
	"strconv"
)

type ConnecInfo struct {
	ipaddr string //port and addr
}

/* file info */

func GetPacket (conn net.Conn) {
/*	b := make([]byte, 1)
	for {
		conn.Read(b)
		fmt.Println(string(b))
	}
	*/

	decode := gob.NewDecoder(conn)
	fmt.Println("GetPacket")
	reader := bufio.NewReader(conn)
	var p Packet
	response, _ := reader.ReadString('\n')
	response = response[:len(response)-1]
	id,_ := strconv.Atoi(response);
	fmt.Println("ID:", id)
	switch id {
		case PKT_FILE_CREATE:
			p = &PktFileCreate{}
		case PKT_FILE_UPDATE:
			p = &PktFileUpdate{}
		case PKT_FILE_DELETE:
			p = &PktFileDelete{}
		case PKT_FILE_REQUEST_CHUNK:
			p = &PktFileRequestChunk{}
		case PKT_FILE_REQUEST_CHUNK_MAP:
			p = &PktFileRequestChunkMap{}
		case PKT_FILE_CHUNK:
			p = &PktFileChunk{}
		case PKT_FILE_CHUNK_MAP:
			p = &PktFileChunkMap{}
		case PKT_CLIENT_INFO:
			p = &PktClientInfo{}
		case PKT_SERVER_FILE_TREE:
			p = &PktServerFileTree{}
	}

	err := decode.Decode(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Have Packet")
	fmt.Println(p)
	p.Print()
	p.Handle()
}
