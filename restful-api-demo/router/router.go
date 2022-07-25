package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-practice/restful-api-demo/apps"
	"github.com/go-practice/restful-api-demo/apps/host"
)

type Handler struct {
	svc host.Service
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Config() {
	if apps.HostService == nil {
		panic("dependencies host hostservice is nil")
	}
	h.svc = apps.HostService
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.CreateHost)
}
