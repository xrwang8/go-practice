package impl

import (
	"context"
	"testing"

	"github.com/go-practice/restful-api-demo/apps/host"
	"github.com/infraboard/mcube/logger/zap"
)

var hs *HostService

func TestCreareHost(t *testing.T) {

	host := host.NewHost()
	host.Name = "test"
	hs.CreateHost(context.Background(), host)

}

func init() {
	zap.DevelopmentSetup()
	hs = NewHostService()
}
