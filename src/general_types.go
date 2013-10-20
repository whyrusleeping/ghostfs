package main

import (
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
