package act

import (
	"context"
	"testing"

	"github.com/tglivelink/livelink-go/pkg/client"
)

func init() {
	client.DefaultClient = client.New(client.Secret{
		SigKey: "your sig_key",
		SecKey: "your sec_key",
	})
}

func TestCallFlow(t *testing.T) {
	rsp := client.Response{}
	err := NewActApi().CallFlow(context.Background(), &client.Param{
		ActId:      6351,
		LivePlatId: "huya",
		GameId:     "yxzj",
		User:       &client.PlatUser{Userid: "xx"},
	}, &CallFlowReq{FlowId: "274224f1"}, &rsp)
	t.Logf("%v %v", rsp, err)
}

func TestReceiveAward(t *testing.T) {
	rsp, err := NewActApi().ReceiveAward(context.Background(), &client.Param{
		ActId:      6571,
		LivePlatId: "huya",
		GameId:     "yxzj",
		User:       &client.PlatUser{Userid: "xxxxx"},
	}, &ReceiveAwardReq{FlowId: "b364e211", OrderId: "12345678901234567"})
	t.Logf("%v %v", rsp, err)
}

func TestGetActList(t *testing.T) {
	rsp, err := NewActApi().GetActList(context.Background(), &GetActListReq{
		Page:       1,
		Size:       10,
		LivePlatId: "huya",
		GameId:     "yxzj",
		IsOnline:   true,
	})
	t.Logf("%v %v", rsp, err)
}
