package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/go-practice/restful-api-demo/apps"
	_ "github.com/go-practice/restful-api-demo/apps/host/all"
	"github.com/go-practice/restful-api-demo/conf"
	"github.com/go-practice/restful-api-demo/router"
	"github.com/spf13/cobra"
)

var confFile string

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 restful-api-demo 后端API",
	Long:  "启动 restful-api-demo 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := conf.LoadConfigFromToml(confFile); err != nil {
			panic(err)
		}
		apps.Init()
		handler := router.NewHandler()
		handler.Config()
		g := gin.Default()
		handler.Registry(g)
		return g.Run(conf.C().App.HttpAddr())

	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "restful-api-demo配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
