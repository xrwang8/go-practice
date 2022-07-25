package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-practice/restful-api-demo/apps"
	"github.com/go-practice/restful-api-demo/apps/host"
)

type Handler struct {
	svc host.Service
}

var handler = &Handler{}

func (h *Handler) Config() {
	h.svc = apps.GetHostService(host.AppName).(host.Service)
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.CreateHost)
}

func (h *Handler) Name() string {
	return host.AppName
}

func init() {
	apps.RegistryHandler(handler)
}
