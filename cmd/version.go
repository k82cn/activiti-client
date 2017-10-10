package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Activiti Client",
	Long:  `All software has versions. This is acli's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Command line for Activiti v0.1 -- HEAD")
	},
}
