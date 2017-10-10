package main

import (
	"fmt"
	"os"

	"github.com/k82cn/activiti-client/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
