package sign

import (
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/tglivelink/livelink-go/pkg/client"
)

var secret = client.Secret{}

func init() {
	secret = client.Secret{
		SigKey: "your sigKey",
		SecKey: "your secKey",
	}
	client.DefaultClient = client.New(secret)
}

// TestSignForMiniProgram 测试拉起小程序
func TestSignForMiniProgram(t *testing.T) {
	param := MiniProgramReq{
		Param: client.Param{
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
	arg, _ := SignForMiniProgram(&param, &secret)
	t.Logf("arg:%s", arg.Encode())

	// {
	// 	client.DefaultClient.Do(context.Background(), &client.Head{
	// 		PathOrApiName: "/api/h5/loginPlatUserInH5",
	// 		Param: &client.Param{
	// 			ActId:      0,
	// 			LivePlatId: "",
	// 			GameId:     "",
	// 			User:       nil,
	// 			Ext:        map[string]string{},
	// 		},
	// 		Body: map[string]string{
	// 			"rawUrl": arg.Encode(),
	// 		},
	// 	})
	// }
}

// TestSignForWeb xxxx
func TestSignForWeb(t *testing.T) {
	param := client.Param{
		ActId:      6490,
		GameId:     "cf",
		LivePlatId: "huya",
		User: &client.PlatUser{
			Userid: "xxxx",
		},
		Ext: map[string]string{},
	}

	arg, _ := Sign(&param, &secret)
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
	// 	client.DefaultClient.Do(context.Background(), &client.Head{
	// 		PathOrApiName: "toRequest",
	// 		Param: &client.Param{
	// 			ActId:      0,
	// 			LivePlatId: "",
	// 			GameId:     "",
	// 			User:       nil,
	// 			Ext: map[string]string{
	// 				"c": "Api",
	// 			},
	// 		},
	// 		Body: data,
	// 	})
	// }
}

// TestSignForLivelinkType xxxx
func TestSignForLivelinkType(t *testing.T) {
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
	code, sign, err := SignForLivelinkLoginType(param, &secret)
	if err != nil {
		t.Fatalf("err:%v", err)
	}
	t.Logf("code:%s sign:%s, t:%d, nonce:%s", code, sign, param.T, param.Nonce)
	// 直接用计算后的数据发起请求
	// {
	// 	rsp, err := http.Post("https://s1.livelink.qq.com/api/reverseBind/login",
	// 		"application/json",
	// 		strings.NewReader(
	// 			fmt.Sprintf(`{"game":"yxzj","type":4,"livelink":{"t":%d,"nonce":"%s","code":"%s","sig":"%s"}}`,
	// 				param.T, param.Nonce, code, sign)),
	// 	)
	// 	if err != nil {
	// 		log.Fatalf("Post err:%v %v", err, param)
	// 	}
	// 	defer rsp.Body.Close()

	// 	bs, _ := io.ReadAll(rsp.Body)
	// 	t.Logf("resp:%s  %v", bs, param)
	// }
}
