package kubo_rpc_client

import (
	"context"
	"testing"

	"github.com/ipfs/kubo/client/rpc"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/assert"
)

func Test_coreAPI_Cat(t *testing.T) {
	// "Connect" to local node
	addr, err := multiaddr.NewMultiaddr("/ip4/192.168.31.12/tcp/5001") // Using 127.0.0.1 for local testing
	if err != nil {
		assert.NoError(t, err)
		return
	}
	node, err := rpc.NewApi(addr)
	if err != nil {
		assert.NoError(t, err)
		return
	}

	api := CoreAPI(node)
	ctx := context.Background()

	// 1. Add a file to get its CID
	// testContent := "Hello IPFS Cat!"
	// testFileName := "test_cat_file.txt"
	// fileNode := files.NewReaderFile(strings.NewReader(testContent))
	// statelessFile := files.NewMapDirectory(map[string]files.Node{
	// 	testFileName: fileNode,
	// })

	// addedPath, err := api.Add(ctx, statelessFile)
	// assert.NoError(t, err)
	// assert.NotNil(t, addedPath)

	// The Add command wraps the file in a directory if multiple files or a directory is added.
	// If a single file is added, it returns the CID of the file itself.
	// For simplicity, let's assume Add returns the CID of the added content directly
	// or we extract it. Here, we'll use the path from Add which should be the CID of the content.
	// If Add wraps it in a directory, you might need to adjust to get the file's CID.
	// For this example, we assume `addedPath.RootCid()` gives the CID of `testContent`.

	// 2. Cat the file using its CID
	// catPath := addedPath.String() // This should be the CID string like "/ipfs/Qm..."
	contentBytes, err := api.Cat(ctx, "QmSycZbTru7yzdse311XXoH8iMj4nZ8PeaJzNZc9YFP7Cp")
	assert.NoError(t, err)
	assert.Equal(t, "testContent", string(contentBytes))
}
