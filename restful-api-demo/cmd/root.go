package cmd

import (
	"github.com/spf13/cobra"
)

var vers bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "restful-api-demo",
	Short: "restful-api-demo 后端API",
	Long:  "restful-api-demo 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			// fmt.Println(version.FullVersion())
			return nil
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "print restful-api-demo version")
}
