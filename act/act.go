package act

import (
	"context"
	"fmt"

	"github.com/huangzixiang5/livelink-go/pkg/client"
)

type ActApi interface {
	// CallFlow 调用活动流程
	CallFlow(ctx context.Context, head *client.ReqParam, req map[string]interface{}, opts ...client.Options) (rsp client.RspHead, err error)
}

// NewActApi
func NewActApi() ActApi {
	return &actApi{
		client: client.DefaultClient,
	}
}

/****************/

type actApi struct {
	client client.Client
}

func (aa *actApi) CallFlow(ctx context.Context, head *client.ReqParam, req map[string]interface{}, opts ...client.Options) (rsp client.RspHead, err error) {

	if req == nil || req["flowId"] == nil {
		err = fmt.Errorf("缺少flowId参数")
		return
	}
	if head.ActId <= 0 {
		err = fmt.Errorf("缺少活动ID")
		return
	}

	h := &client.ReqHead{
		PathOrApiName: "apiRequest",
		ReqParam:      *head,
	}

	err = aa.client.Do(ctx, h, req, &rsp, opts...)

	return
}
