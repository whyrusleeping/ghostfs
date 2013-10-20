package main

import (
	"fmt"
)

func (p* PktClientInfo) Handle() {
	fmt.Println("Handle ClientInfo")
}

func (p* PktServerFileTree) Handle() {
	fmt.Println("Handle ServerFileTree")

	for i:=0; i < len(p.Files); i++ {
		fmt.Println(p.Files[i])
	}
}

func (p* PktFileCreate) Handle() {
	// create a shadow
	fmt.Println("Handle FileCreate")
}

func (p* PktFileUpdate) Handle() {
	fmt.Println("Handle FileUpdate")
}

func (p* PktFileDelete) Handle() {
	fmt.Println("Handle FileDelete")
}

func (p* PktFileRequestChunk) Handle() {
	fmt.Println("Handle FileRequestChunk")
}

func (p* PktFileRequestChunkMap) Handle() {
	fmt.Println("Handle FileRequestChunkMap")
}

func (p* PktFileChunk) Handle() {
	fmt.Println("Handle FileChunk")
}

func (p* PktFileChunkMap) Handle() {
	fmt.Println("Handle FileChunkMap")
}
