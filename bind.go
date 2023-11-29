package livelink

import (
	"context"

	"github.com/huangzixiang5/livelink-go/client"
)

// BindApi 绑定相关api
type BindApi interface {
	// GetBoundGameRole 拉取当前用户已绑定的游戏账号信息
	GetBoundGameRole(ctx context.Context, head *client.ReqHead, opts ...client.Options) (rsp GetBindInfoRsp, err error)
}

// NewBindApi xxxx
func NewBindApi() BindApi {
	return &bindApi{
		client: client.DefaultClient,
	}
}

/*************************************************8*/

type bindApi struct {
	client client.Client
}

// GetBindInfoRsp 拉取绑定信息
type GetBindInfoRsp struct {
	client.RspBase
	JData struct {
		IsBind  bool `json:"isBind"` // 是否已绑定
		GameAcc struct {
			Type string `json:"type"` // 游戏账号类型，"qq" or "wx"
		} `json:"gameAcc"`
		GameRole struct { // 游戏角色信息
			RoleName      string `json:"roleName"`
			AreaName      string `json:"areaName"`
			PartitionName string `json:"partitionName"`
			PlatName      string `json:"platName"`
		} `json:"gameRole"`
	} `json:"jData"`
}

// GetBoundGameRole 拉取当前用户已绑定的游戏账号信息
func (ba *bindApi) GetBoundGameRole(ctx context.Context, head *client.ReqHead, opts ...client.Options) (rsp GetBindInfoRsp, err error) {

	head.PathOrApiName = "GetBindInfo"
	if err = ba.client.Do(ctx, head, nil, &rsp, opts...); err != nil {
		return
	}

	return
}
