## 辅助计算请求签名

接入方可以参考以下方式生成接口签名，发起接口调用

### 拉起小程序签名 
```go
import "github.com/huangzixiang5/livelink-go/sign"

param := client.ReqParam{
	GameId:     "cf",
	LivePlatId: "huya",
	User: &client.PlatUser{
		Userid: "hughtest",
	},
	Ext: map[string]string{
		"gameAuthScene": "act_1",
		"faceUrl":       "http://baidu.com",
		"nickName":      "我",
	},
}

// 注意，小程序这里需要使用 codec.SignerMd5Fixed 类型签名 
arg, _ := sign.Sign(&param, "xxxx", "xxxxx", client.WithSigner(codec.SignerMd5Fixed))

// 使用encode后的参数拉起小程序 
fmt.Println(arg.Encode())

```

### 接口请求签名 
```go 

import "github.com/huangzixiang5/livelink-go/sign"

param := client.ReqParam{
	GameId:     "cf",
	LivePlatId: "huya",
	User: &client.PlatUser{
		Userid: "xxx",
	},
	Ext: map[string]string{},
}

arg, _ := Sign(&param, "xxxxx", "xxxx")

// encode后的数据，拼接在请求的url后面即可 
fmt.Println(arg.Encode())

```