package main

import (
	"fmt"
	"net"
	"encoding/gob"
	"github.com/whyrusleeping/swagfs/sfs_types"
)

type Client struct {
	Con net.Conn
	ServCom chan sfs.Message
	OutGoing chan sfs.Message

	Dec *gob.Decoder
	Enc *gob.Encoder
}

func (c *Client) SendMessage(m sfs.Message) {
	fmt.Println("Sending:")
	fmt.Println(m)
	c.OutGoing <- m
}

func (c *Client) Start() {
	go c.sync()

	var m sfs.Message
	for {
		err := c.Dec.Decode(m)
		if err != nil {
			fmt.Println("Client Read Loop:")
			fmt.Println(err)
		}
		c.ServCom <- m
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
