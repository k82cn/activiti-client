package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	InitClient("admin", "test", "http://activiti.druid.io/activiti-app/api")
}

var RootCmd = &cobra.Command{
	Use:   "acli",
	Short: "Activiti CLI is a command line tool of Activiti",
	Long:  `A command line of Activiti by REST API.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
