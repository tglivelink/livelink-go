package bind

import (
	"context"
	"testing"

	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/pkg/config"
)

// TestGetBoundGameRole xxxx
func TestGetBoundGameRole(t *testing.T) {
	config.ConfigPath = "../livelink.yaml"
	NewBindApi().GetBoundGameRole(context.Background(), &client.ReqParam{
		LivePlatId: "huya",
		GameId:     "cf",
		User:       &client.PlatUser{Userid: "xxxxx"},
		FromGame:   false,
	},
	)
}
