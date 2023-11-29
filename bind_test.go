package livelink

import (
	"context"
	"testing"

	"github.com/huangzixiang5/livelink-go/client"
)

// TestGetBoundGameRole xxxx
func TestGetBoundGameRole(t *testing.T) {
	NewBindApi().GetBoundGameRole(context.Background(), &client.ReqHead{
		LivePlatId: "huya",
		GameId:     "cf",
		User:       &client.PlatUser{Userid: "xxxxx"},
		FromGame:   false,
	})
}
