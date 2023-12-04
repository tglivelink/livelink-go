package sign

import (
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/pkg/config"
)

func init() {
	config.SetGlobalConfig(&config.Config{
		Server: &config.ServerConfig{
			Domain: "https://s1.livelink.qq.com",
		},
		Client: &config.ClientConfig{
			Appid:  "huya",
			SigKey: "your sig_key",
			SecKey: "your sec_key",
		},
	})
}

// TestSignForMiniProgram 测试拉起小程序
func TestSignForMiniProgram(t *testing.T) {
	param := MiniProgramReq{
		ReqParam: client.ReqParam{
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
	arg, _ := SignForMiniProgram(&param, config.GlobalConfig().Client)
	t.Logf("arg:%s", arg.Encode())

	// {
	// 	client.DefaultClient.Do(context.Background(), &client.ReqHead{
	// 		PathOrApiName: "/api/h5/loginPlatUserInH5",
	// 		ReqParam: client.ReqParam{
	// 			ActId:      0,
	// 			LivePlatId: "",
	// 			GameId:     "",
	// 			User:       nil,
	// 			FromGame:   false,
	// 			Ext:        map[string]string{},
	// 		},
	// 	}, map[string]string{
	// 		"rawUrl": arg.Encode(),
	// 	}, nil)
	// }
}

// TestSignForWeb xxxx
func TestSignForWeb(t *testing.T) {
	param := client.ReqParam{
		ActId:      6490,
		GameId:     "cf",
		LivePlatId: "huya",
		User: &client.PlatUser{
			Userid: "xxxx",
		},
		Ext: map[string]string{},
	}

	arg, _ := Sign(&param, config.GlobalConfig().Client)
	t.Logf("arg:%s", arg.Encode())

	// 直接用计算后的数据发起请求
	{
		rsp, err := http.Post("https://s1.livelink.qq.com/livelink?a=apiRequest&"+arg.Encode(),
			"application/json", strings.NewReader(`{"flowId":""}`))
		if err != nil {
			log.Fatalf("Post err:%v", err)
		}
		defer rsp.Body.Close()

		bs, _ := io.ReadAll(rsp.Body)
		t.Logf("resp:%s", bs)
	}

	// {
	// 	data := make(map[string]string)
	// 	for k, v := range arg {
	// 		data[k] = url.QueryEscape(v[0])
	// 	}
	// 	config.ConfigPath = "../livelink.yaml"
	// 	client.DefaultClient.Do(context.Background(), &client.ReqHead{
	// 		PathOrApiName: "toRequest",
	// 		ReqParam: client.ReqParam{
	// 			ActId:      0,
	// 			LivePlatId: "",
	// 			GameId:     "",
	// 			User:       nil,
	// 			FromGame:   false,
	// 			Ext: map[string]string{
	// 				"c": "Api",
	// 			},
	// 		},
	// 	}, data, nil)
	// }
}
