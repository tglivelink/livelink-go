package client

import (
	"fmt"
	"time"

	"github.com/tglivelink/livelink-go/pkg/errs"
	"github.com/tglivelink/livelink-go/pkg/util"
)

// Param 请求的url参数
type Param struct {
	ActId      int64             // 活动id
	LivePlatId string            // 平台code
	GameId     string            // 游戏id
	User       User              // 用户信息, 必须是PlatUser或者GameUser
	Ext        map[string]string // 扩展字段
}

// FixToKVs 转换成请求需要的kv形式，并且会补充上时间戳、随机盐值、版本等字段
func (r *Param) FixToKVs() map[string]string {

	kvs := make(map[string]string, 8+len(r.Ext))
	kvs["livePlatId"] = r.LivePlatId
	kvs["actId"] = fmt.Sprintf("%d", r.ActId)
	kvs["gameId"] = r.GameId
	kvs["gameIdList"] = r.GameId
	kvs["t"] = fmt.Sprintf("%d", time.Now().Unix())
	kvs["nonce"] = fmt.Sprintf("%x", util.RandBytes(3))
	kvs["v"] = "2.0"
	if _, ok := r.User.(*GameUser); ok {
		kvs["fromGame"] = "1"
	}
	for k, v := range r.Ext {
		kvs[k] = v
	}
	return kvs
}

func (r *Param) AddExt(k, v string) {
	if r.Ext == nil {
		r.Ext = make(map[string]string)
	}
	r.Ext[k] = v
}

func (p *Param) Check() error {
	if p.ActId <= 0 {
		return errs.ErrActIdInvalid
	}
	if p.LivePlatId == "" {
		return errs.ErrLivePlatIdInvalid
	}
	if p.GameId == "" {
		return errs.ErrGameIdInvalid
	}
	if p.User == nil || p.User.Key() == "" {
		return errs.ErrUserInvalid
	}
	return nil
}

// PlatUser 平台用户
type PlatUser struct {
	Userid string `json:"userid"` // 用户uid，必填
}

func (p *PlatUser) Key() string {
	if p == nil {
		return ""
	}
	return p.Userid
}

func (*PlatUser) user() {}

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
	AccType      string `json:"accType,omitempty"` // qq、wx
}

func (g *GameUser) Key() string {
	return g.GameOpenId
}

func (*GameUser) user() {}

type User interface {
	// 私有方法，禁止业务自己实现， 固定只能是 PlatUser/GameUser
	user()
	Key() string
}
