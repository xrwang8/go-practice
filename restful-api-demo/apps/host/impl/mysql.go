package impl

import (
	"database/sql"
	"github.com/go-practice/restful-api-demo/apps"
	"github.com/go-practice/restful-api-demo/apps/host"
	"github.com/go-practice/restful-api-demo/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var hostService = &HostService{}

type HostService struct {
	l  logger.Logger
	db *sql.DB
}

func (hs *HostService) Config() {
	hs.l = zap.L().Named("Host")
	hs.db = conf.C().MySQL.GetDB()
}

func (hs *HostService) Name() string {
	return host.AppName
}
func init() {
	apps.Registry(hostService)

}
