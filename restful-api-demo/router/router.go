package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-practice/restful-api-demo/apps/host"
)

type Handler struct {
	host.Service
}

func NewHandler(svc host.Service) *Handler {
	return &Handler{
		svc,
	}
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.CreateHost)
}
