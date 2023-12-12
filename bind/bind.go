package bind

import (
	"context"
	"fmt"

	"github.com/tglivelink/livelink-go/pkg"
	"github.com/tglivelink/livelink-go/pkg/client"
)

// BindApi 绑定相关api
type BindApi interface {
	// GetBoundGameRole 拉取当前用户已绑定的游戏账号信息
	GetBoundGameRole(ctx context.Context, head *client.Param, opts ...client.Options) (rsp GetBoundGameRoleRsp, err error)
	// GetBoundGameRoleInAct 拉取当前用户在某个活动的绑定关系（针对不可换绑活动）
	GetBoundGameRoleInAct(ctx context.Context, param *client.Param, opts ...client.Options) (rsp GetBoundGameRoleInActRsp, err error)
	// BindGameRoleInAct 将用户当前游戏账号应用于某个活动 （针对不可换绑活动）
	BindGameRoleInAct(ctx context.Context, param *client.Param, opts ...client.Options) (rsp client.ResponseBase, err error)
}

// NewBindApi xxxx
func NewBindApi(opts ...client.Options) BindApi {
	return &bindApi{
		api: pkg.NewApi(opts...),
	}
}

/**************************************************/

type bindApi struct {
	api *pkg.Api
}

// GetBoundGameRoleRsp 拉取绑定信息
type GetBoundGameRoleRsp struct {
	client.ResponseBase
	Data struct {
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

// GetBoundGameRole
func (ba *bindApi) GetBoundGameRole(ctx context.Context, param *client.Param, opts ...client.Options) (rsp GetBoundGameRoleRsp, err error) {

	ctx, head := ba.api.Head(ctx)
	head.PathOrApiName = "GetBindInfo"
	head.Param = param
	head.Rsp = &rsp

	err = ba.api.Request(ctx, head, opts...)

	return
}

/*************************/

type GetBoundGameRoleInActRsp struct {
	client.ResponseBase
	Data struct {
		ChangBindTime string `json:"changBindTime"` // 下次可更换绑定的时间,"2006-01-02 00:00:00", 为空表示该活动后续都不可更换绑定
		GameAcc       struct {
			Type    string `json:"type"` // qq or wx
			Faceurl string `json:"faceurl"`
			Nick    string `json:"nick"`
		} `json:"gameAcc"`
		GameRole struct {
			RoleName      string `json:"roleName"`
			AreaName      string `json:"areaName"`
			PartitionName string `json:"partitionName"`
			PlatName      string `json:"platName"`
		} `json:"gameRole"`
	} `json:"jData"`
}

// GetBoundGameRoleInAct
func (ba *bindApi) GetBoundGameRoleInAct(ctx context.Context, param *client.Param, opts ...client.Options) (rsp GetBoundGameRoleInActRsp, err error) {

	if param.ActId <= 0 {
		err = fmt.Errorf("缺少活动ID")
		return
	}

	ctx, head := ba.api.Head(ctx)
	head.PathOrApiName = "GetActBind"
	head.Param = param
	head.Rsp = &rsp

	err = ba.api.Request(ctx, head, opts...)

	return
}

/**********************************************/
func (ba *bindApi) BindGameRoleInAct(ctx context.Context, param *client.Param, opts ...client.Options) (rsp client.ResponseBase, err error) {

	if param.ActId <= 0 {
		err = fmt.Errorf("缺少活动ID")
		return
	}

	ctx, head := ba.api.Head(ctx)
	head.PathOrApiName = "ActBind"
	head.Param = param
	head.Rsp = &rsp

	err = ba.api.Request(ctx, head, opts...)

	return
}
