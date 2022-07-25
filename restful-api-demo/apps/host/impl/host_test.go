package impl

import (
	"context"
	"fmt"
	"github.com/go-practice/restful-api-demo/conf"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/go-practice/restful-api-demo/apps/host"
	"github.com/infraboard/mcube/logger/zap"
)

var hs *HostService

func TestCreate(t *testing.T) {
	should := assert.New(t)
	ins := host.NewHost()
	ins.Id = "ins-01"
	ins.Name = "test"
	ins.Region = "cn-hangzhou"
	ins.Type = "sm1"
	ins.CPU = 1
	ins.Memory = 2048
	ins, err := hs.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}
}

func init() {
	// 测试用例的配置文件
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	// 需要初始化全局Logger,
	// 为什么不设计为默认打印, 因为性能
	fmt.Println(zap.DevelopmentSetup())
	hs = NewHostService()
}
