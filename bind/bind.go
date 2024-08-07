package bind

import (
	"context"
	"fmt"

	"github.com/tglivelink/livelink-go/pkg"
	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/pkg/errs"
)

// BindApi 绑定相关api
type BindApi interface {
	// GetBoundGameRole 拉取当前用户已绑定的游戏账号信息
	GetBoundGameRole(ctx context.Context, param *client.Param, opts ...client.Options) (rsp GetBoundGameRoleRsp, err error)
	// GetBoundGameRoleInAct 拉取当前用户在某个活动的绑定关系（针对不可换绑活动）
	GetBoundGameRoleInAct(ctx context.Context, param *client.Param, opts ...client.Options) (rsp GetBoundGameRoleInActRsp, err error)
	// BindGameRoleInAct 将用户当前游戏账号应用于某个活动 （针对不可换绑活动）
	BindGameRoleInAct(ctx context.Context, param *client.Param, opts ...client.Options) (rsp client.ResponseBase, err error)
	// GetBoundQQ 拉取当前用户绑定的QQ号
	GetBoundQQ(ctx context.Context, param *client.Param, opts ...client.Options) (rsp GetBoundQQRsp, err error)
	// Unbind 解绑游戏账号
	Unbind(ctx context.Context, param *client.Param, req *UnbindReq, opts ...client.Options) (rsp UnbindRsp, err error)
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
			RoleName      string `json:"roleName"`      // 角色名称
			AreaName      string `json:"areaName"`      // 大区名
			PartitionName string `json:"partitionName"` // 小区名
			PlatName      string `json:"platName"`      // iOS 、 Android
		} `json:"gameRole"`
	} `json:"jData"`
}

// GetBoundGameRole
func (ba *bindApi) GetBoundGameRole(ctx context.Context, param *client.Param, opts ...client.Options) (rsp GetBoundGameRoleRsp, err error) {

	if param.LivePlatId == "" {
		err = errs.ErrLivePlatIdInvalid
		return
	}
	if param.GameId == "" {
		err = errs.ErrGameIdInvalid
		return
	}
	if err = ba.api.CheckUser(ctx, param.User); err != nil {
		return
	}

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

	if err = param.Check(); err != nil {
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

	if err = param.Check(); err != nil {
		return
	}

	ctx, head := ba.api.Head(ctx)
	head.PathOrApiName = "ActBind"
	head.Param = param
	head.Rsp = &rsp

	err = ba.api.Request(ctx, head, opts...)

	return
}

/*************************************/

type GetBoundQQRsp struct {
	client.ResponseBase
	Data struct {
		Uin           string `json:"uin"`           // 模糊的qq号，类似 ‘2323****’，为空表示未绑定
		AccountAvatar string `json:"accountAvatar"` // 头像
		AccountName   string `json:"accountName"`   // 昵称
	} `json:"jData"`
}

func (ba *bindApi) GetBoundQQ(ctx context.Context, param *client.Param, opts ...client.Options) (rsp GetBoundQQRsp, err error) {
	if param.LivePlatId == "" {
		err = errs.ErrLivePlatIdInvalid
		return
	}
	if err = ba.api.CheckUser(ctx, param.User); err != nil {
		return
	}
	ctx, head := ba.api.Head(ctx)
	head.PathOrApiName = "GetRegUin"
	head.Param = param
	head.Rsp = &rsp

	err = ba.api.Request(ctx, head, opts...)

	return
}

type UnbindReq struct {
	GameRoleId string `json:"gameRoleId"` // 查询接口时返回的绑定角色唯一标识
}

type UnbindRsp struct {
	client.ResponseBase
	Data struct{} `json:"jData"`
}

func (ba *bindApi) Unbind(ctx context.Context, param *client.Param, req *UnbindReq, opts ...client.Options) (
	rsp UnbindRsp, err error) {
	if param.LivePlatId == "" {
		err = errs.ErrLivePlatIdInvalid
		return
	}
	if err = ba.api.CheckUser(ctx, param.User); err != nil {
		return
	}

	if req.GameRoleId == "" {
		err = fmt.Errorf("gameRoleId is invalid")
		return
	}
	body := make(map[string]interface{})
	body["gameRoleId"] = req.GameRoleId

	ctx, head := ba.api.Head(ctx)
	head.PathOrApiName = "Unbind"
	head.Param = param
	head.Body = body
	head.Rsp = &rsp

	err = ba.api.Request(ctx, head, opts...)

	return
}
