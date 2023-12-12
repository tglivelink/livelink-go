# livelink-go

提供livelink常用接口（golang）的封装,具体可参考[接入文档](https://livelink.qq.com/doc/activities/)

## 目录结构
```
.
├── act 与活动相关接口 
├── bind 与绑定相关接口 
├── pkg 内部封装，一般情况下无需关注,默认即可 
│   ├── client 实际请求客户端
│   ├── codec 序列化、签名实现、加密实现等 
│   ├── errs 错误码定义 
│   ├── log 请求/响应日志
│   └── util
└── sign 辅助计算签名的工具 
```

## 调用示例
```go

import (
	"context"

	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/bind"
)

// step.1 初始化全局客户端,只需要设置一次即可,默认所有的请求都会使用这个客户端
client.DefaultClient = client.New(client.Secret{
	SigKey: "your sig_key",
	SecKey: "your sec_key",
})


// eg. 直接调用拉取绑定的游戏角色信息
NewBindApi().GetBoundGameRole(context.Background(), &client.Param{
	LivePlatId: "huya",
	GameId:     "cf",
	User:       &client.PlatUser{Userid: "xxxxx"},
},
)

// eg. 如果部分api请求需要使用不同的签名信息，可以这样设置，不会影响全局
api := NewBindApi(client.WithSecret(client.Secret{
	SigKey: "other sig_key",
	SecKey: "other sec_key",
}))
api.GetBoundGameRole(context.Background(), &client.Param{
	LivePlatId: "huya",
	GameId:     "cf",
	User:       &client.PlatUser{Userid: "xxxxx"},
},
)
```

## 直接使用底层发起调用
```go 
import (
	"github.com/tglivelink/livelink-go/pkg/client"
)

// 创建一个全局请求客户端
cli := client.New(client.Secret{
	SigKey: "your sig_key",
	SecKey: "your sec_key",
})

// 使用客户端发起请求
rsp := make(map[string]interface{})
err := cli.Do(context.Background(), &client.Head{
	PathOrApiName: "GetBindInfo", // 后端接口 
	Param: &client.Param{
		LivePlatId: "huya",
		GameId:     "yxzj",
		User:       &client.PlatUser{Userid: "xxx"},
	},
	Rsp: &rsp,
})
if err != nil {
	return
}

```



### 自定义日志输出
```go
// 设置自己的日志输出,需要实现pkg/log/Logger接口,默认会将日志打印到标准输出 
// 设置为nil则关闭日志 
log.DefaultLogger = Logger
```