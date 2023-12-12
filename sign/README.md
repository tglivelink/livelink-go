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