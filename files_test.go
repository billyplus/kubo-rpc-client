package kubo_rpc_client

import (
	"context"
	"testing"

	"github.com/ipfs/kubo/client/rpc"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/assert"
)

func Test_filesAPI_Stat(t *testing.T) {
	type args struct {
		ctx  context.Context
		path string
		opt  []APIOption
	}
	// "Connect" to local node
	addr, err := multiaddr.NewMultiaddr("/ip4/192.168.31.12/tcp/5001")
	if err != nil {
		assert.NoError(t, err)
		return
	}
	node, err := rpc.NewApi(addr)
	if err != nil {
		assert.NoError(t, err)
		return
	}
	tests := []struct {
		name    string
		api     *filesAPI
		args    args
		want    *FileStatResult
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "root",
			api:  (FilesAPI(node)),
			args: args{
				ctx:  context.Background(),
				path: "/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.api.Stat(tt.args.ctx, tt.args.path, tt.args.opt...)
			assert.NotNil(t, got)
			assert.NoError(t, err)
		})
	}
}
