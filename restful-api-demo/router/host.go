package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-practice/restful-api-demo/apps/host"
	"github.com/infraboard/mcube/http/response"
)

func (h *Handler) CreateHost(c *gin.Context) {
	host := host.NewHost()
	// 用户传递过来的参数进行解析, 实现了一个json 的unmarshal
	if err := c.Bind(host); err != nil {
		response.Failed(c.Writer, err)
		return
	}

	host, err := h.svc.CreateHost(c.Request.Context(), host)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	// 成功, 把对象实例返回给HTTP API调用方
	response.Success(c.Writer, host)
}
