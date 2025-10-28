## 辅助计算请求签名

接入方可以参考以下方式生成接口签名，自行发起接口调用

### 拉起小程序签名 
```go
import (
	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/sign"
)

// 构建必需参数结构 
param := sign.MiniProgramReq{
	ReqParam: client.Param{
		GameId:     "cf",
		LivePlatId: "huya",
		User: &client.PlatUser{
			Userid: "hughtest",
		},
		Ext: map[string]string{
			"gameAuthScene": "act_1",
		},
	},
	FaceUrl:  "http://baidu.com",
	NickName: "我",
}

// 生成拉起小程序参数
arg, _ := sign.SignForMiniProgram(&param, &client.Secret{
	SigKey: "your sig_key",
	SecKey: "your sec_key",
})

// 使用encode后的参数拉起小程序 
fmt.Println(arg.Encode())

```

### 常规接口请求签名   

```go 

import (
	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/sign"
)

param := client.Param{
		ActId:      6490,
		GameId:     "cf",
		LivePlatId: "huya",
		User: &client.PlatUser{
			Userid: "xxxx",
		},
		Ext: map[string]string{},
	}

	arg, _ := sign.Sign(&param, &client.Secret{
		SigKey: "your sig_key",
		SecKey: "your sec_key",
	})

// encode后的数据，拼接在请求的url后面即可 
fmt.Println(arg.Encode())

``` 

### 反绑拉起签名【Livelink登录类型】

基于LiveLink签名模式登录web/H5页面，详情[参考](https://livelink.qq.com/doc/activities/pages/bind/web-sdk/lm-livelink.html)  

```go 

import (
	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/sign"
)

// 1. 传入游戏相关参数，计算出对应的code、sign等
param := &LivelinkLogin{
	T:     0,
	Nonce: "",
	User: &client.GameUser{
		GameOpenId:   "xxxxxxxx",
		RoleId:       "xxxxxx",
		Area:         1,
		PlatId:       0,
		Partition:    1130,
		GameNickName: "",
		HeadImg:      "",
		AreaName:     "",
		PlatName:     "",
		RoleName:     "",
		AccType:      "qq",
	},
}
code, sign, err := sign.SignForLivelinkLoginType(param, &client.Secret{
		SigKey: "your sig_key",
		SecKey: "your sec_key",
	})
if err != nil {
	t.Fatalf("err:%v", err)
}

// 2. 将生成好的参数传入
/**
var instance = new LivelinkManager();
// 加密版本（推荐使用）
instance.init({
  loginType: 'livelink', // 登录方式
  actId: 475, // 活动ID，没有可传入为0
  gameId: 'lol', // 游戏ID
  t: param.T, // 秒级时间戳
  code: code,
  sig: sign,
  nonce: param.Nonce,
  onBoundSuccess(data) {
    // 绑定账号成功的回调
    console.log('onBoundSuccess', data);
  },
  onBoundError(errData) {
    // 绑定账号失败的回调
    console.log('onBoundError', errData);
  },
  initialedEvent(data) {
    // 初始化成功的回调
    console.log('initialedEvent', data);
  },
});
*/

```