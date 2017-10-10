package get

import (
	"fmt"

	"github.com/k82cn/activiti-client/api"
	"github.com/k82cn/activiti-client/cmd"
)

func init() {
	registerGetCmd("process", &GetProcessCmd{
		processHisURL:     "history/historic-process-instances",
		acitivityQueryURL: "query/historic-activity-instances",
		variableQueryURL:  "query/historic-variable-instances",
	})
}

type GetProcessCmd struct {
	processHisURL     string
	acitivityQueryURL string
	variableQueryURL  string
}

func (g *GetProcessCmd) Get() error {
	outputFormat := "%-10s%-50s%-30s%-30s\n"

	var pl api.ProcessList
	err := cmd.Client.Get(g.processHisURL+"?size=10000", &pl)
	cmd.CheckErr(err, &pl)

	fmt.Printf(outputFormat, "ID", "BusinessKey", "StartTime", "EndTime")
	for _, p := range pl.Data {
		fmt.Printf(outputFormat, p.ID, p.BusinessKey,
			cmd.FormatTime(p.StartTime), cmd.FormatTime(p.EndTime))
	}

	return nil
}

func (g *GetProcessCmd) Describe(pid string) error {
	var p api.Process
	url := g.processHisURL + "/" + pid
	err := cmd.Client.Get(url, &p)
	cmd.CheckErr(err, &p)

	fmt.Printf("%-22s: %s\n", "ID", p.ID)
	fmt.Printf("%-22s: %s\n", "Business Key", p.BusinessKey)
	fmt.Printf("%-22s: %s\n", "Process Definition Id", p.ProcessDefinitionId)
	fmt.Printf("%-22s: %s\n", "Start Time", p.StartTime)
	fmt.Printf("%-22s: %s\n", "End Time", p.EndTime)
	fmt.Printf("%-22s: \n\n", "Variables")

	var vl api.VariableList

	err = cmd.Client.Post(g.variableQueryURL, &api.Query{
		ProcessInstanceId: pid,
	}, &vl)

	cmd.CheckErr(err, &vl)

	varFmt := "%-35s%-15s%-15s%s\n"
	fmt.Printf(varFmt, "Name", "Scope", "Type", "Value")

	for _, v := range vl.Data {
		fmt.Printf(varFmt, v.Variable.Name, v.Variable.Scope, v.Variable.Type, v.Variable.Value)
	}

	fmt.Printf("\n")

	fmt.Printf("%-22s: \n\n", "Tasks")

	var al api.ActivityList

	err = cmd.Client.Post(g.acitivityQueryURL, &api.Query{
		ProcessInstanceId: pid,
	}, &al)

	cmd.CheckErr(err, &al)

	activityFmt := "%-10s%-45s%-15s%-30s%-20s\n"
	fmt.Printf(activityFmt, "ID", "Activity Id", "Assignee", "Start Time", "End Time")

	for _, a := range al.Data {
		fmt.Printf(activityFmt, a.ID, a.ActivityID, a.Assignee,
			cmd.FormatTime(a.StartTime), cmd.FormatTime(a.EndTime))
	}

	return nil
}
