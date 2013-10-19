package main

import (
	"time"
)

type file struct {
	path  string
	hash  string
	mtime time.Time
}

type ServerFileTree struct {
	files []file
}

