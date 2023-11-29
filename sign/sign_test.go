package sign

import (
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/huangzixiang5/livelink-go/pkg/client"
	"github.com/huangzixiang5/livelink-go/pkg/codec"
)

// TestSignForMiniProgram 测试拉起小程序
func TestSignForMiniProgram(t *testing.T) {
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

	// 注意，小程序这里需要使用 codec.SignerMd5Fixed 签名方式
	arg, _ := Sign(&param, "xxxxxxx", "xxxxxx", client.WithSigner(codec.SignerMd5Fixed))
	t.Logf("arg:%s", arg.Encode())

	// {
	// 	config.ConfigPath = "../livelink.yaml"
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
			Userid: "2211471928",
		},
		Ext: map[string]string{},
	}

	arg, _ := Sign(&param, "ea58755ce4320a2c", "09e645299b1a7e28")
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
