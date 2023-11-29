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
