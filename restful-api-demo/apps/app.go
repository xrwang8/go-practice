package apps

import (
	"fmt"
	"github.com/go-practice/restful-api-demo/apps/host"
)

// IOC  管理所有的服务实例对象

var (
	HostService host.Service
	svcs        = map[string]Service{}
)

type Service interface {
	Config()
	Name() string
}

func Registry(svc Service) {
	if _, ok := svcs[svc.Name()]; ok {
		panic(fmt.Sprintf("Service %s has registered", svc.Name()))
	}
	svcs[svc.Name()] = svc
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

func Init() {
	for _, v := range svcs {
		v.Config()
	}
}
