package create

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/cobra"

	"github.com/k82cn/activiti-client/cmd"
)

func init() {
	cmd.RootCmd.AddCommand(createCmd)
}

type CreateCmd interface {
	Create(args []string)
}

var mapMutex sync.Mutex
var createSubCmdMap = make(map[string]CreateCmd)

func registerCreateCmd(name string, subCmd CreateCmd) {
	mapMutex.Lock()
	defer mapMutex.Unlock()

	createSubCmdMap[name] = subCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("What are you want to create?")
			os.Exit(1)
		}

		createSubCmd, found := createSubCmdMap[args[0]]

		if !found {
			fmt.Printf("%s is not support.", args[0])
			os.Exit(1)
		}

		createSubCmd.Create(args[1:])
	},
}
