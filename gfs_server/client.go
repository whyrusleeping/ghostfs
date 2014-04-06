package main

import (
	"fmt"
	"net"
	"encoding/gob"
	"github.com/whyrusleeping/ghostfs/gfs_types"
)

type Client struct {
	Con net.Conn
	ServCom chan gfs.Message
	OutGoing chan gfs.Message

	Dec *gob.Decoder
	Enc *gob.Encoder
}

func (c *Client) SendMessage(m gfs.Message) {
	fmt.Println("Sending:")
	fmt.Println(m)
	c.OutGoing <- m
}

func (c *Client) Start() {
	go c.sync()

	var m gfs.Message
	for {
		err := c.Dec.Decode(&m)
		fmt.Println("Got message from client.")
		if err != nil {
			fmt.Println("Client Read Loop:")
			fmt.Println(err)
			return
		}
		m.SetCallback(c.OutGoing)
		c.ServCom <- m
		fmt.Println("Message relayed to server.")
	}
}

func (c *Client) sync() {
	for {
		select {
			case out := <-c.OutGoing:
				err := c.Enc.Encode(&out)
				if err != nil {
					fmt.Println(err)
				}
		}
	}
}
