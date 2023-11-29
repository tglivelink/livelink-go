package client

import (
	"fmt"
	"time"

	"github.com/huangzixiang5/livelink-go/pkg/util"
)

// ReqHead 请求
type ReqHead struct {
	PathOrApiName string // 请求路径
	ReqParam             // 其他请求参数
}

// ReqParam 请求参数
type ReqParam struct {
	ActId      int64             // 活动id
	LivePlatId string            // 平台code
	GameId     string            // 游戏id
	User       User              // 用户信息, PlatUser or GameUser
	FromGame   bool              // 是否来源游戏
	Ext        map[string]string // 扩展字段
}

// FixToKVs 转换成请求需要的kv形式，并且会补充上时间戳、随机盐值、版本等字段
func (r *ReqParam) FixToKVs() map[string]string {

	kvs := make(map[string]string, 16)
	kvs["livePlatId"] = r.LivePlatId
	kvs["actId"] = fmt.Sprintf("%d", r.ActId)
	kvs["gameId"] = r.GameId
	kvs["gameIdList"] = r.GameId
	kvs["t"] = fmt.Sprintf("%d", time.Now().Unix())
	kvs["nonce"] = fmt.Sprintf("%x", util.RandBytes(3))
	kvs["v"] = "2.0"
	if r.FromGame {
		kvs["fromGame"] = "1"
	}
	for k, v := range r.Ext {
		kvs[k] = v
	}
	return kvs
}

// PlatUser 平台用户
type PlatUser struct {
	Userid   string `json:"userid"`             // 用户uid，必填
	ClientIp string `json:"clientIp,omitempty"` // 用户ip 0.0.0.0
}

func (*PlatUser) nocopy() {}

// GameUser 如果是游戏发起的请求，需要携带如下字段
type GameUser struct {
	GameOpenId   string `json:"gameOpenId"`
	RoleId       string `json:"roleId,omitempty"`
	Area         int    `json:"area,omitempty"`
	PlatId       int    `json:"platId,omitempty"`
	Partition    int    `json:"partition,omitempty"`
	GameNickName string `json:"gameNickName,omitempty"`
	HeadImg      string `json:"headImg,omitempty"`
	AreaName     string `json:"areaName,omitempty"`
	PlatName     string `json:"platName,omitempty"`
	RoleName     string `json:"roleName,omitempty"`
}

func (*GameUser) nocopy() {}

type User interface {
	// nocopy
	nocopy()
}
