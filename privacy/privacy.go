package privacy

import (
	"context"

	"github.com/tglivelink/livelink-go/pkg"
	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/pkg/errs"
)

// PrivacyApi 隐私授权相关api
type PrivacyApi interface {
	// IsAuthorized 判断用户是否已经授权
	IsAuthorized(ctx context.Context, param *client.Param, req *IsAuthorizedReq, opts ...client.Options) (rsp IsAuthorizedRsp, err error)
}

func NewPrivacyApi(opts ...client.Options) PrivacyApi {
	return &privacyApi{
		api: pkg.NewApi(opts...),
	}
}

/**********************************************/

type privacyApi struct {
	api *pkg.Api
}

type IsAuthorizedReq struct {
	Scene string `json:"scene"`
}

type IsAuthorizedRsp struct {
	// Deprecated: use data.IsGranted
	IsGranted bool `json:"isGranted"`
	client.ResponseBase
	Data struct {
		IsGranted bool `json:"isGranted"`
	} `json:"jData"`
}

func (pa *privacyApi) IsAuthorized(ctx context.Context, param *client.Param, req *IsAuthorizedReq,
	opts ...client.Options) (rsp IsAuthorizedRsp, err error) {

	if param.LivePlatId == "" {
		err = errs.ErrLivePlatIdInvalid
		return
	}
	if err = pa.api.CheckUser(ctx, param.User); err != nil {
		return
	}

	ctx, head := pa.api.Head(ctx)
	head.Param = param
	head.Body = req
	head.Rsp = &rsp
	head.PathOrApiName = "GetGameGrantInfo"
	err = pa.api.Request(ctx, head, opts...)
	return
}
