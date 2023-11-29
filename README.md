# livelink-go
提供livelink常用接口的封装

## 调用示例
```go
    // 拉取绑定的游戏角色信息
	NewBindApi().GetBoundGameRole(context.Background(), &client.ReqHead{
		LivePlatId: "huya",
		GameId:     "cf",
		User:       &client.PlatUser{Userid: "xxxxx"},
		FromGame:   false,
	})

    // 生成拉起小程序需要的参数
    param := MiniProgramReq{
		GameIdList: "cf",
		LivePlatId: "huya",
		T:          time.Now().Unix(),
		User: client.PlatUser{
			Userid: "hughhuangtest",
		},
		FaceUrl:  "http://baidu.com",
		NickName: "我",
		Ext: map[string]string{
			"gameAuthScene": "act_1",
		},
	}

	arg, _ := ArgsForMiniProgram(&param, "1111222233334444", "1111222233334444")
```

## 配置
```yaml
domain: "https://s1.livelink.qq.com" # livelink后端域名
appid: "huya" # 请求方标识
sig_key: "xxxxxxx" # 计算sig需要的key 
sec_key: "xxxxxx" # 计算用户code需要的key,用户敏感信息是通过密文传输
```

## 自定义功能
```go
// 设置自己的配置加载器，需要实现config.ConfigLoader接口
config.DefaultConfigLoader = MyConfigLoader 

```