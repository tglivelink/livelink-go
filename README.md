# livelink-go

提供livelink常用接口（golang）的封装,具体可参考[接入文档](https://livelink.qq.com/doc/activities/)

## 目录结构
```
.
├── act 与活动相关接口 
├── bind 与绑定相关接口 
├── pkg 内部封装，一般情况下无需关注,默认即可 
│   ├── client 请求客户端
│   ├── codec 序列化、签名实现、加密实现等 
│   ├── config 配置信息 
│   ├── log 请求/响应日志
│   └── util
└── sign 辅助计算签名的工具 
```

## 配置信息
> 接口调用需要使用到的配置信息如下
```yaml
server:
  domain: "https://s1.livelink.qq.com"
client:
  appid: "your appid" # 请求方标识
  sig_key: "your sig_key" # 计算sig需要的key 
  sec_key: "your sec_key" # 计算用户code需要的key,用户敏感信息是通过密文传输
```


## 调用示例
```go

import (
	"context"

	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/bind"
)

// 指定配置文件加载路径 
config.ConfigPath = "../livelink.yaml"

// 拉取绑定的游戏角色信息
bind.NewBindApi().GetBoundGameRole(context.Background(), &client.ReqParam{
	LivePlatId: "huya",
	GameId:     "cf",
	User:       &client.PlatUser{Userid: "xxxxx"},
	FromGame:   false,
})

// 执行活动流程调用 
act.NewActApi().CallFlow(context.Background(), &client.ReqParam{
	ActId:      6512,
	LivePlatId: "huya",
	GameId:     "yxzj",
	User:       &client.PlatUser{},
	FromGame:   false,
}, map[string]interface{}{
	"flowId": "xxx",
})
```


## 自定义功能
pkg目录下提供了相关能力（配置、签名、请求客户端等）的默认实现，如果有特殊需求，可以通过其中暴露的接口修改  

### 自定义配置加载
```go
// 1. 可以自定义配置文件的加载路径,默认值为 “./livelink.yaml”，注意：需要在第一次发起调用前指定 
config.ConfigPath = "../livelink.yaml"

// 2. 可以设置自己的配置加载器，需要实现pkg/config/ConfigLoader接口
config.DefaultConfigLoader = MyConfigLoader 

// 3. 可以直接调用接口进行设置
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

```

### 自定义日志输出
```go
// 设置自己的日志输出,需要实现pkg/log/Logger接口,默认会将日志打印到标准输出 
log.DefaultLogger = Logger
```