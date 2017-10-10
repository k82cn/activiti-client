package get

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/cobra"

	"github.com/k82cn/activiti-client/cmd"
)

func init() {
	cmd.RootCmd.AddCommand(getCmd)
}

type Get interface {
	Get() error
	Describe(id string) error
}

var mapMutex sync.Mutex
var getSubCmdMap = make(map[string]Get)

func registerGetCmd(name string, subCmd Get) {
	mapMutex.Lock()
	defer mapMutex.Unlock()

	getSubCmdMap[name] = subCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("What are you want to get?")
			os.Exit(1)
		}

		getSubCmd, found := getSubCmdMap[args[0]]

		if !found {
			fmt.Printf("%s is not support.", args[0])
			os.Exit(1)
		}

		if len(args) > 1 {
			for _, id := range args[1:] {
				getSubCmd.Describe(id)
			}
		} else {
			getSubCmd.Get()
		}
	},
}
