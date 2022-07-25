package impl

import (
	"database/sql"
	"github.com/go-practice/restful-api-demo/apps/host"
	"github.com/go-practice/restful-api-demo/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var _ host.Service = (*HostService)(nil)

type HostService struct {
	l  logger.Logger
	db *sql.DB
}

func NewHostService() *HostService {
	return &HostService{
		zap.L().Named("Host"),
		conf.C().MySQL.GetDB(),
	}
}
