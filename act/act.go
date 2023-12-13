package act

import (
	"context"
	"fmt"

	"github.com/tglivelink/livelink-go/pkg"
	"github.com/tglivelink/livelink-go/pkg/client"
)

type ActApi interface {
	// CallFlow 调用活动流程
	CallFlow(ctx context.Context, param *client.Param, req *CallFlowReq, rsp interface{}, opts ...client.Options) (err error)
	// ReceiveAward 领取礼包，需要携带幂等号uniq
	ReceiveAward(ctx context.Context, param *client.Param, req *ReceiveAwardReq, opts ...client.Options) (rsp ReceiveAwardRsp, err error)
	// GetActList 拉取活动列表
	GetActList(ctx context.Context, req *GetActListReq) (rsp GetActListRsp, err error)
}

// NewActApi
func NewActApi(opts ...client.Options) ActApi {
	return &actApi{
		api: pkg.NewApi(opts...),
	}
}

/****************/

type actApi struct {
	api *pkg.Api
}

/************************************************* 调用活动流程 **********/

type CallFlowReq struct {
	FlowId string
	Other  map[string]interface{}
}

func (aa *actApi) CallFlow(ctx context.Context, param *client.Param, req *CallFlowReq,
	rsp interface{}, opts ...client.Options) (err error) {

	if req.FlowId == "" {
		err = fmt.Errorf("flowId is empty")
		return
	}
	if param.ActId <= 0 {
		err = fmt.Errorf("actId is invalid")
		return
	}

	body := req.Other
	if body == nil {
		body = make(map[string]interface{})
	}
	body["flowId"] = req.FlowId

	ctx, head := aa.api.Head(ctx)
	head.PathOrApiName = "apiRequest"
	head.Param = param
	head.Body = body
	head.Rsp = rsp

	err = aa.api.Request(ctx, head, opts...)

	return
}

/****************************************************** 调用发货接口 **********/

type ReceiveAwardReq struct {
	FlowId  string
	OrderId string // 唯一订单号，一般是 16~32字节
	Other   map[string]interface{}
}

type ReceiveAwardRsp struct {
	client.ResponseBase
	Data struct {
		Message          string   `json:"message"`           // 提示信息
		PackageId        int      `json:"packageId"`         // 前活动下礼包组ID，可用于标识唯一礼包
		PackageName      string   `json:"packageName"`       // 已发放的礼包组中文名称
		PackageNum       int      `json:"packageNum"`        // 当前发放的礼包个数
		PackageRealFlag  string   `json:"sPackageRealFlag"`  // 该是否为实物。1表示该道具为实物道具,0为游戏虚拟道具
		PackageOtherInfo string   `json:"sPackageOtherInfo"` // 预留字段，礼包补充信息
		CDKey            string   `json:"sCDKey"`            // 如果礼包为cdkey，此处会填充cdkey，合作平台需要设计弹出框为用户弹出此字段
		Ext              struct { // 发货道具扩展字段,特殊情况会携带
			PrizeExchange struct {
				FlowName string `json:"flowName"`
				FlowID   string `json:"flowID"`  // 下一步需要调用的flowId
				NeedUin  bool   `json:"needUin"` // 是否需要绑定QQ
			} `json:"prizeExchange"`
		} `json:"ext"`
	} `json:"jData"`
}

func (aa *actApi) ReceiveAward(ctx context.Context, param *client.Param, req *ReceiveAwardReq,
	opts ...client.Options) (rsp ReceiveAwardRsp, err error) {
	if param.User == nil || param.User.Key() == "" {
		err = fmt.Errorf("user is empty")
		return
	}
	if req.OrderId == "" {
		err = fmt.Errorf("OrderId is empty")
		return
	}

	body := req.Other
	if body == nil {
		body = make(map[string]interface{})
	}
	body["serialCode"] = req.OrderId

	err = aa.CallFlow(ctx, param, &CallFlowReq{FlowId: req.FlowId, Other: body}, &rsp, opts...)
	return
}

/**************************************** 拉取活动列表 **************/

type GetActListReq struct {
	Page       int    `json:"page"`
	Size       int    `json:"size"`
	LivePlatId string `json:"plat"`
	GameId     string `json:"game"`
	IsOnline   bool   `json:"isOnline"`
}

type GetActListRsp struct {
	client.ResponseBase
	Data struct {
		Total   int `json:"total"`
		ActList []struct {
			Id        uint64 `json:"id"`        // 活动id
			BeginTime string `json:"beginTime"` // 开始时间 2006-01-02 00:00:00
			EndTime   string `json:"endTime"`   // 结束时间 2006-01-02 00:00:00
			Game      string `json:"game"`      // 游戏
			Plat      string `json:"plat"`      // 平台
			ActName   string `json:"actName"`   // 活动名称
		} `json:"actList"`
	} `json:"jData"`
}

func (aa *actApi) GetActList(ctx context.Context, req *GetActListReq) (rsp GetActListRsp, err error) {
	if req.LivePlatId == "" {
		err = fmt.Errorf("LivePlatId is empty")
		return
	}
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Size < 1 {
		req.Size = 10
	}
	param := client.Param{
		LivePlatId: req.LivePlatId,
		User:       &client.PlatUser{Userid: "00"},
	}
	body := map[string]interface{}{
		"page":     req.Page,
		"size":     req.Size,
		"game":     req.GameId,
		"isOnline": req.IsOnline,
	}
	ctx, head := aa.api.Head(ctx)
	head.Param = &param
	head.Body = body
	head.PathOrApiName = "ApiActList"
	head.Rsp = &rsp

	err = aa.api.Request(ctx, head)
	return
}

/***********************************************/
