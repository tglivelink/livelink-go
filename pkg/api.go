package pkg

import (
	"context"

	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/pkg/errs"
)

type Api struct {
	client client.Client
	opts   []client.Options
}

func NewApi(opts ...client.Options) *Api {
	return &Api{
		client: client.DefaultClient,
		opts:   opts,
	}
}

func (api *Api) Request(ctx context.Context, head *client.Head, opts ...client.Options) error {
	callOpts := make([]client.Options, 0, len(api.opts)+len(opts))
	callOpts = append(callOpts, api.opts...)
	callOpts = append(callOpts, opts...)
	return api.client.Do(ctx, head, callOpts...)
}

func (api *Api) Head(ctx context.Context) (context.Context, *client.Head) {
	return client.EnsureHead(ctx)
}

func (api *Api) CheckUser(ctx context.Context, user client.User) error {

	if user == nil || user.Key() == "" {
		return errs.ErrUserInvalid
	}
	return nil
}
