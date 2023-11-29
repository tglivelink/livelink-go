package client

type ReqHead struct {
	PathOrApiName string            // 请求路径
	ActId         int64             // 活动id
	LivePlatId    string            // 平台code
	GameId        string            // 游戏id
	User          User              // 用户信息, PlatUser or GameUser
	FromGame      bool              // 是否来源游戏
	Ext           map[string]string // 扩展字段
}

// PlatUser 平台用户
type PlatUser struct {
	Userid string `json:"userid"`
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
