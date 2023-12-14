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
├── privacy 与隐私授权相关接口 
└── sign 辅助计算签名的工具 
```

## 调用示例
```go

import (
	"context"
	
	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/act"
	"github.com/tglivelink/livelink-go/bind"
)

func init() {
	// step.1 建议在调用前，采用如下方式初始化全局秘钥,这时只需要设置一次即可，后续请求默认都会使用这个客户端 
	client.DefaultClient = client.New(client.Secret{
		SigKey: "your sig_key",
		SecKey: "your sec_key",
	})
}

func main() {
	// eg. 直接调用拉取绑定的游戏角色信息
	bind.NewBindApi().GetBoundGameRole(context.Background(), &client.Param{
		LivePlatId: "huya",
		GameId:     "cf",
		User:       &client.PlatUser{Userid: "xxxxx"},
	})

	// eg. 调用发货流程
	rsp, err := act.NewActApi().ReceiveAward(context.Background(), &client.Param{
		ActId:      6571,
		LivePlatId: "huya",
		GameId:     "yxzj",
		User:       &client.PlatUser{Userid: "xxxxx"},
	}, &ReceiveAwardReq{FlowId: "b364e211", OrderId: "12345678901234567"})
	
	// eg. 如果需要临时使用其他秘钥，可以这样设置，只会影响当前api,不会影响全局秘钥
	api := bind.NewBindApi(client.WithSecret(client.Secret{
		SigKey: "other sig_key",
		SecKey: "other sec_key",
	}))
	api.GetBoundGameRole(context.Background(), &client.Param{
		LivePlatId: "huya",
		GameId:     "cf",
		User:       &client.PlatUser{Userid: "xxxxx"},
	})
}

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