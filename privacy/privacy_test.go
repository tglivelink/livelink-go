package privacy

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

func TestIsAuthorized(t *testing.T) {
	rsp, err := NewPrivacyApi().IsAuthorized(context.Background(), &client.Param{
		LivePlatId: "huya",
		GameId:     "nz",
		User:       &client.PlatUser{Userid: "xxxx"},
	}, &IsAuthorizedReq{Scene: "act_6585"})
	t.Logf("%v %v", rsp, err)
}
