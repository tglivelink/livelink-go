package livelink

import (
	"fmt"
	"testing"
	"time"

	"github.com/huangzixiang5/livelink-go/client"
)

func TestArgsForMiniProgram(t *testing.T) {
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
	fmt.Println(arg)
}

func TestArgsForIframe(t *testing.T) {
	param := IframeReq{
		GameId:     "cf",
		Timestamp:  time.Now().Unix(),
		LivePlatId: "huya",
		ActId:      123,
		User: client.PlatUser{
			Userid: "hughhuangtest",
		},
	}
	arg, _ := ArgsForIframe(&param, "1111222233334444", "1111222233334444")
	fmt.Println(arg)
}
