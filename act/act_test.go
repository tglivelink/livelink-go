package act

import (
	"context"
	"testing"

	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/pkg/config"
)

func init() {
	config.ConfigPath = "../livelink.yaml"
}

func TestCallFlow(t *testing.T) {
	NewActApi().CallFlow(context.Background(), &client.ReqParam{
		ActId:      6512,
		LivePlatId: "huya",
		GameId:     "yxzj",
		User:       &client.PlatUser{},
		FromGame:   false,
	}, map[string]interface{}{
		"flowId": "xxx",
	})
}
