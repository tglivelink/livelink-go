# livelink-go
提供livelink常用接口的封装

## 目录结构
```
.
├── act 活动相关接口 
├── bind 绑定相关接口 
├── pkg 内部封装，一般情况下无需关注 
│   ├── client 请求客户端
│   ├── codec 序列化、签名实现、加密实现等 
│   ├── config 配置信息 
│   ├── log 请求/响应日志
│   └── util
└── sign 辅助计算签名的工具 
```

## 调用示例
```go

import (
	"context"

	"github.com/huangzixiang5/livelink-go/pkg/client"
	"github.com/huangzixiang5/livelink-go/bind"
)

// 拉取绑定的游戏角色信息
bind.NewBindApi().GetBoundGameRole(context.Background(), &client.ReqParam{
	LivePlatId: "huya",
	GameId:     "cf",
	User:       &client.PlatUser{Userid: "xxxxx"},
	FromGame:   false,
})

```

## 配置信息
```yaml
server:
  domain: "https://s1.livelink.qq.com"
client:
  appid: "your appid" # 请求方标识
  sig_key: "your sig_key" # 计算sig需要的key 
  sec_key: "your sec_key" # 计算用户code需要的key,用户敏感信息是通过密文传输

```

## 自定义功能
pkg目录下提供了相关能力的默认实现，包括配置、签名等，如果需要自定义可以使用下面的方式修改
```go
// 设置自己的配置加载器，需要实现pkg/config/ConfigLoader接口
config.DefaultConfigLoader = MyConfigLoader 
// 或者直接设置
config.SetGlobalConfig(&config.Config{
	Server: &config.ServerConfig{
		Domain: "https://s1.livelink.qq.com",
	},
	Client: &config.ClientConfig{
		Appid:  "huya",
		SigKey: "xxxx",
		SecKey: "xxxx",
	},
})

// 设置自己的日志打印,需要实现pkg/log/Logger接口 
log.DefaultLogger = Logger

```

