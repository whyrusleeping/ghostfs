package main

import (
	"os"
	"time"
	"net"
)

type file struct {
	Path  string
	Hash  string
	Mtime time.Time
}

type ServerFileTree struct {
	Files []file
}

type Node struct {
	Entries []*Node
	Type os.FileMode
	Name string
}

type client struct {
	conn net.Conn
	id int
}

type cclient struct {
	conn net.Conn
	id int
	live bool
	addr string
}
