package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/k82cn/activiti-client/api"
)

func init() {
	RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "process":
			getProcess(args[1:])
		}
	},
}

func getProcess(pids []string) {
	if len(pids) == 0 {
		outputFormat := "%-10s%-50s%-30s%-30s\n"

		var pl api.ProcessList
		err := client.Get("history/historic-process-instances?size=10000", &pl)
		CheckErr(err, &pl)

		fmt.Printf(outputFormat, "ID", "BusinessKey", "StartTime", "EndTime")
		for _, p := range pl.Data {
			endTime, startTime := "", ""

			if p.EndTime != nil {
				endTime = p.EndTime.Format(DefaultTimeFormat)
			}

			if p.StartTime != nil {
				startTime = p.StartTime.Format(DefaultTimeFormat)
			}

			fmt.Printf(outputFormat, p.ID, p.BusinessKey, startTime, endTime)
		}
	} else {
		for _, pid := range pids {
			describeProcess(pid)
		}
	}
}

func describeProcess(pid string) {
	var p api.Process
	url := "history/historic-process-instances/" + pid
	err := client.Get(url, &p)
	CheckErr(err, &p)

	fmt.Printf("%-22s: %s\n", "ID", p.ID)
	fmt.Printf("%-22s: %s\n", "Business Key", p.BusinessKey)
	fmt.Printf("%-22s: %s\n", "Process Definition Id", p.ProcessDefinitionId)
	fmt.Printf("%-22s: %s\n", "Start Time", p.StartTime)
	fmt.Printf("%-22s: %s\n", "End Time", p.EndTime)
	fmt.Printf("%-22s: \n", "Variables")

}
