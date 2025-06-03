package kubo_rpc_client

import (
	"context"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
	caopts "github.com/ipfs/kubo/core/coreiface/options"
)

type coreAPI rpc.HttpApi

func CoreAPI(api *rpc.HttpApi) *coreAPI {
	return (*coreAPI)(api)
}

/*
	{
	  "Bytes": "<int64>",
	  "Hash": "<string>",
	  "Mode": "<string>",
	  "Mtime": "<int64>",
	  "MtimeNsecs": "<int>",
	  "Name": "<string>",
	  "Size": "<string>"
	}
*/
type AddResult struct {
	Bytes      int64  `json:"Bytes"`
	Hash       string `json:"Hash"`
	Mode       string `json:"Mode"`
	Mtime      int64  `json:"Mtime"`
	MtimeNsecs int    `json:"MtimeNsecs"`
	Name       string `json:"Name"`
	Size       string `json:"Size"`
}

/*
## /api/v0/add

Add a file or directory to IPFS.

### Arguments

- `quiet` [bool]: Write minimal output. Required: no.
- `quieter` [bool]: Write only final hash. Required: no.
- `silent` [bool]: Write no output. Required: no.
- `progress` [bool]: Stream progress data. Required: no.
- `trickle` [bool]: Use trickle-dag format for dag generation. Required: no.
- `only-hash` [bool]: Only chunk and hash - do not write to disk. Required: no.
- `wrap-with-directory` [bool]: Wrap files with a directory object. Required: no.
- `chunker` [string]: Chunking algorithm, size-[bytes], rabin-[min]-[avg]-[max] or buzhash.UnixFSChunker. Required: no.
- `raw-leaves` [bool]: Use raw blocks for leaf nodes.UnixFSRawLeaves. Required: no.
- `max-file-links` [int]: Limit the maximum number of links in UnixFS file nodes to this value. (experimental)UnixFSFileMaxLinks. Required: no.
- `max-directory-links` [int]: Limit the maximum number of links in UnixFS basic directory nodes to this value.UnixFSDirectoryMaxLinks. WARNING: experimental, Import.UnixFSHAMTThreshold is a safer alternative. Required: no.
- `max-hamt-fanout` [int]: Limit the maximum number of links of a UnixFS HAMT directory node to this (power of 2, multiple of 8).UnixFSHAMTDirectoryMaxFanout WARNING: experimental, see Import.UnixFSHAMTDirectorySizeThreshold as well. Required: no.
- `nocopy` [bool]: Add the file using filestore. Implies raw-leaves. (experimental). Required: no.
- `fscache` [bool]: Check the filestore for pre-existing blocks. (experimental). Required: no.
- `cid-version` [int]: CID version. Defaults to 0 unless an option that depends on CIDv1 is passed. Passing version 1 will cause the raw-leaves option to default to true.CidVersion. Required: no.
- `hash` [string]: Hash function to use. Implies CIDv1 if not sha2-256.HashFunction. Required: no.
- `inline` [bool]: Inline small blocks into CIDs. (experimental). Required: no.
- `inline-limit` [int]: Maximum block size to inline. (experimental). Default: `32`. Required: no.
- `pin` [bool]: Pin locally to protect added files from garbage collection. Default: `true`. Required: no.
- `to-files` [string]: Add reference to Files API (MFS) at the provided path. Required: no.
- `preserve-mode` [bool]: Apply existing POSIX permissions to created UnixFS entries. Disables raw-leaves. (experimental). Required: no.
- `preserve-mtime` [bool]: Apply existing POSIX modification time to created UnixFS entries. Disables raw-leaves. (experimental). Required: no.
- `mode` [uint]: Custom POSIX file mode to store in created UnixFS entries. Disables raw-leaves. (experimental). Required: no.
- `mtime` [int64]: Custom POSIX modification time to store in created UnixFS entries (seconds before or after the Unix Epoch). Disables raw-leaves. (experimental). Required: no.
- `mtime-nsecs` [uint]: Custom POSIX modification time (optional time fraction in nanoseconds). Required: no.

### Request Body

Argument path is of file type. This endpoint expects one or several files (depending on the command) in the body of the request as 'multipart/form-data'.
*/
func (api *coreAPI) Add(ctx context.Context, f files.Node, opts ...caopts.UnixfsAddOption) (path.ImmutablePath, error) {
	return (*rpc.UnixfsAPI)(api).Add(ctx, f, opts...)
}

/*
	{
		"Objects": [
		{
			"Hash": "<string>",
			"Links": [
			{
				"Hash": "<string>",
				"ModTime": "<timestamp>",
				"Mode": "<uint32>",
				"Name": "<string>",
				"Size": "<uint64>",
				"Target": "<string>",
				"Type": "<int32>"
			}
			]
		}
		]
	}
*/
type ListEntry struct {
	Hash    string `json:"Hash"`
	ModTime string `json:"ModTime"`
	Mode    uint32 `json:"Mode"`
	Name    string `json:"Name"`
	Size    uint64 `json:"Size"`
	Target  string `json:"Target"`
	Type    int32  `json:"Type"`
}

type ListObjectData struct {
	Hash  string      `json:"Hash"`
	Links []ListEntry `json:"Links"`
}

type ListResult struct {
	Objects []ListObjectData `json:"Objects"`
}

func (api *coreAPI) List(ctx context.Context, hash string, opt ...APIOption) (*ListResult, error) {
	opts := make([]APIOption, 0, len(opt)+1)
	opts = append(opts, WithArgs(hash))
	opts = append(opts, opt...)
	return Request[ListResult](ctx, (*rpc.HttpApi)(api), "ls", opts...)
}

/*
	{
	  "Commit": "<string>",
	  "Golang": "<string>",
	  "Repo": "<string>",
	  "System": "<string>",
	  "Version": "<string>"
	}
*/
type VersionResult struct {
	Commit  string `json:"Commit"`
	Golang  string `json:"Golang"`
	Repo    string `json:"Repo"`
	System  string `json:"System"`
	Version string `json:"Version"`
}

func (api *coreAPI) Version(ctx context.Context, opt ...APIOption) (*VersionResult, error) {
	return Request[VersionResult](ctx, (*rpc.HttpApi)(api), "version", opt...)
}
