package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-practice/restful-api-demo/apps/host"
)

// IOC  管理所有的服务实例对象

var (
	HostService host.Service
	hsMap       = map[string]HSService{}
	handlerMap  = map[string]HandlerService{}
)

type HSService interface {
	Config()
	Name() string
}

func Registry(svc HSService) {
	if _, ok := hsMap[svc.Name()]; ok {
		panic(fmt.Sprintf("Service %s has registered", svc.Name()))
	}
	hsMap[svc.Name()] = svc
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

func GetHostService(name string) interface{} {
	for k, v := range hsMap {
		if k == name {
			return v
		}
	}
	return nil
}

func InitHS() {
	for _, v := range hsMap {
		v.Config()
	}
}

type HandlerService interface {
	Registry(r gin.IRouter)
	Config()
	Name() string
}

func RegistryHandler(svc HandlerService) {
	// 服务实例注册到svcs map当中
	if _, ok := handlerMap[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}
	handlerMap[svc.Name()] = svc
}

func InitHandler(r gin.IRouter) {
	for _, v := range handlerMap {
		v.Config()
	}
	for _, v := range handlerMap {
		v.Registry(r)
	}
}
