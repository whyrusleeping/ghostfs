package main 
import (
)

type clientConfig struct {
	pathToRoot string
	localFiles []file 
	serverIP connecInfo
	clientIPs []connecInfo
}
