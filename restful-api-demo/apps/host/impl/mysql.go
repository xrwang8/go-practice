package impl

import (
	"github.com/go-practice/restful-api-demo/apps/host"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var _ host.Service = (*HostService)(nil)

type HostService struct {
	l logger.Logger
}

func NewHostService() *HostService {
	return &HostService{
		zap.L().Named("Host"),
	}
}
