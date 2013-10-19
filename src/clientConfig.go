package main 
import (
	"fmt"
)

type clientConfig struct {
	pathToRoot string
	localFiles file[]
	serverIP connecInfo
	clientIPs []connecInfo
}


