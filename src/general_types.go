package main

import (
	"time"
	"net"
)

type file struct {
	path  string
	hash  string
	mtime time.Time
}

type ServerFileTree struct {
	files []file
}

type client struct {
	conn net.Conn
	id int
}

