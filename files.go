package kubo_rpc_client

import (
	"context"

	"github.com/ipfs/kubo/client/rpc"
)

type filesAPI rpc.HttpApi

func FilesAPI(api *rpc.HttpApi) *filesAPI {
	return (*filesAPI)(api)
}

/*
	{
	  "Blocks": "<int>",
	  "CumulativeSize": "<uint64>",
	  "Hash": "<string>",
	  "Local": "<bool>",
	  "Mode": "<uint32>",
	  "Mtime": "<int64>",
	  "MtimeNsecs": "<int>",
	  "Size": "<uint64>",
	  "SizeLocal": "<uint64>",
	  "Type": "<string>",
	  "WithLocality": "<bool>"
	}
*/
type FileStatResult struct {
	Blocks         int    `json:"Blocks"`
	CumulativeSize uint64 `json:"CumulativeSize"`
	Hash           string `json:"Hash"`
	Local          bool   `json:"Local"`
	Mode           uint32 `json:"Mode"`
	Mtime          int64  `json:"Mtime"`
	MtimeNsecs     int    `json:"MtimeNsecs"`
	Size           uint64 `json:"Size"`
	SizeLocal      uint64 `json:"SizeLocal"`
	Type           string `json:"Type"`
	WithLocality   bool   `json:"WithLocality"`
}

func (api *filesAPI) Stat(ctx context.Context, path string, opt ...APIOption) (*FileStatResult, error) {
	opts := make([]APIOption, 0, len(opt)+1)
	opts = append(opts, WithArgs(path))
	opts = append(opts, opt...)
	return Request[FileStatResult](ctx, (*rpc.HttpApi)(api), "files/stat", opts...)
}

/*
	{
	  "Entries": [
	    {
	      "Hash": "<string>",
	      "Name": "<string>",
	      "Size": "<int64>",
	      "Type": "<int>"
	    }
	  ]
	}
*/
type FileListEntry struct {
	Hash string `json:"Hash"`
	Name string `json:"Name"`
	Size int64  `json:"Size"`
	Type int    `json:"Type"`
}

type FileListResult struct {
	Entries []FileListEntry `json:"Entries"`
}

/*
List directories in the local mutable namespace.
Arguments
arg [string]: Path to show listing for. Defaults to '/'. Required: no.
long [bool]: Use long listing format. Required: no.
U [bool]: Do not sort; list entries in directory order. Required: no.
*/
func (api *filesAPI) List(ctx context.Context, path string, opt ...APIOption) (*FileListResult, error) {
	opts := make([]APIOption, 0, len(opt)+1)
	opts = append(opts, WithArgs(path))
	opts = append(opts, opt...)
	return Request[FileListResult](ctx, (*rpc.HttpApi)(api), "files/ls", opts...)
}

/*
Add references to IPFS files and directories in MFS (or copy within MFS).
Arguments
arg [string]: Source IPFS or MFS path to copy. Required: yes.
arg [string]: Destination within MFS. Required: yes.
*/
func (api *filesAPI) Copy(ctx context.Context, src, dst string, opt ...APIOption) error {
	opts := make([]APIOption, 0, len(opt)+2)
	opts = append(opts, WithArgs(src), WithArgs(dst))
	opts = append(opts, opt...)
	return Exec(ctx, (*rpc.HttpApi)(api), "files/cp", opts...)
}

/*
Move files.
Arguments
arg [string]: Source file to move. Required: yes.
arg [string]: Destination path for file to be moved to. Required: yes.
*/
func (api *filesAPI) Move(ctx context.Context, src, dst string, opt ...APIOption) error {
	opts := make([]APIOption, 0, len(opt)+2)
	opts = append(opts, WithArgs(src), WithArgs(dst))
	opts = append(opts, opt...)
	return Exec(ctx, (*rpc.HttpApi)(api), "files/mv", opts...)
}
