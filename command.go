package kubo_rpc_client

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"github.com/ipfs/kubo/client/rpc"
)

type Option struct {
	args []string
	opts map[string]any
}

type APIOption func(*Option)

func WithArgs(args ...string) APIOption {
	return func(o *Option) {
		o.args = append(o.args, args...)
	}
}

func WithOption(key string, val interface{}) APIOption {
	return func(o *Option) {
		o.opts[key] = val
	}
}

func Request[R any](ctx context.Context, ipfsAPI *rpc.HttpApi, cmd string, opt ...APIOption) (*R, error) {
	o := Option{
		opts: map[string]any{},
	}
	for _, oo := range opt {
		oo(&o)
	}

	req := ipfsAPI.Request(cmd, o.args...)
	for k, v := range o.opts {
		req = req.Option(k, v)
	}

	res, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}
	if res.Output == nil {
		return nil, errors.New("no output")
	}
	defer res.Output.Close()

	dec := json.NewDecoder(res.Output)
	var out R
	err = dec.Decode(&out)
	if err != nil {
		if err != io.EOF {
			return nil, err
		}
		return nil, nil
	}

	return &out, nil
}

func RequestRaw(ctx context.Context, ipfsAPI *rpc.HttpApi, cmd string, opt ...APIOption) ([]byte, error) {
	o := Option{
		opts: map[string]any{},
	}
	for _, oo := range opt {
		oo(&o)
	}

	req := ipfsAPI.Request(cmd, o.args...)
	for k, v := range o.opts {
		req = req.Option(k, v)
	}

	res, err := req.Send(ctx)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}
	if res.Output == nil {
		return nil, errors.New("no output")
	}
	defer res.Output.Close()

	return io.ReadAll(res.Output)
}

func Exec(ctx context.Context, ipfsAPI *rpc.HttpApi, cmd string, opt ...APIOption) error {
	o := Option{
		opts: map[string]any{},
	}
	for _, oo := range opt {
		oo(&o)
	}

	req := ipfsAPI.Request(cmd, o.args...)
	for k, v := range o.opts {
		req = req.Option(k, v)
	}

	res, err := req.Send(ctx)
	if err != nil {
		return err
	}
	if res.Error != nil {
		return res.Error
	}

	return nil
}
