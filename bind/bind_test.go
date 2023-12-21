package bind

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

// TestGetBoundGameRole xxxx
func TestGetBoundGameRole(t *testing.T) {
	rsp, err := NewBindApi().GetBoundGameRole(context.Background(), &client.Param{
		LivePlatId: "huya",
		GameId:     "yxzj",
		User:       &client.PlatUser{Userid: "xxxx"},
	})
	t.Logf("%v %v", rsp, err)
}

func TestGetBoundGameRoleInAct(t *testing.T) {
	rsp, err := NewBindApi().GetBoundGameRoleInAct(context.Background(), &client.Param{
		ActId:      6032,
		LivePlatId: "huya",
		GameId:     "yxzj",
		User:       &client.PlatUser{Userid: "xxxx"},
	})
	t.Logf("%v %v", rsp, err)
}

func TestBindGameRoleInAct(t *testing.T) {
	rsp, err := NewBindApi().BindGameRoleInAct(context.Background(), &client.Param{
		ActId:      6032,
		LivePlatId: "huya",
		GameId:     "yxzj",
		User:       &client.PlatUser{Userid: "xxxx"},
	})
	t.Logf("%v %v", rsp, err)
}

func TestGetBoundQQ(t *testing.T) {
	rsp, err := NewBindApi().GetBoundQQ(context.Background(), &client.Param{
		LivePlatId: "huya",
		User:       &client.PlatUser{Userid: "xxx"},
	})
	t.Logf("%v %v", rsp, err)
}
