package main
import (
	"github.com/whyrusleeping/swagfs/sfs_types"
)

type clientConfig struct {
	pathToRoot string
	localFiles []*types.File
	serverIP types.ConnecInfo
	clientIPs []types.ConnecInfo
}
