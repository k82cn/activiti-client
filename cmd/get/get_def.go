package get

import (
	"fmt"
	"strconv"

	"github.com/k82cn/activiti-client/api"
	"github.com/k82cn/activiti-client/cmd"
)

func init() {
	registerGetCmd("def", &GetProcessDefCmd{
		processDefHisURL: "repository/process-definitions",
		processQueryURL:  "query/historic-process-instances",
	})
}

type GetProcessDefCmd struct {
	processDefHisURL string
	processQueryURL  string
}

func (g *GetProcessDefCmd) Get() {
	outputFormat := "%-30s%-30s%-20s%-10s%-10s\n"

	var pl api.ProcessDefList
	err := cmd.Client.Get(g.processDefHisURL+"?size=10000", &pl)
	cmd.CheckErr(err, &pl)

	fmt.Printf(outputFormat, "ID", "Key", "Name", "Version", "Deployment ID")
	pmap := make(map[string]api.ProcessDef)
	vmap := make(map[string]int)
	for _, p := range pl.Data {
		if p.Version > vmap[p.Key] {
			vmap[p.Key] = p.Version
			pmap[p.Key] = p
		}
	}

	for _, p := range pmap {
		fmt.Printf(outputFormat, p.ID, p.Key, p.Name, strconv.Itoa(p.Version), p.DeploymentID)
	}
}

func (g *GetProcessDefCmd) Describe(pid string) {
	var pd api.ProcessDef
	err := cmd.Client.Get(g.processDefHisURL+"/"+pid, &pd)
	cmd.CheckErr(err, &pd)

	fmt.Printf("%-13s : %s\n", "ID", pd.ID)
	fmt.Printf("%-13s : %s\n", "Key", pd.Key)
	fmt.Printf("%-13s : %s\n", "Name", pd.Name)
	fmt.Printf("%-13s : %d\n", "Version", pd.Version)
	fmt.Printf("%-13s : %s\n", "Deployment ID", pd.DeploymentID)
	fmt.Printf("%-13s : \n\n", "Process")

	var pl api.ProcessList

	err = cmd.Client.Post(g.processQueryURL, &api.Query{
		ProcessDefinitionId: &pid,
	}, &pl)

	cmd.CheckErr(err, &pl)

	outputFormat := "%-10s%-50s%-30s%-30s\n"
	fmt.Printf(outputFormat, "ID", "BusinessKey", "StartTime", "EndTime")
	for _, p := range pl.Data {
		fmt.Printf(outputFormat, p.ID, p.BusinessKey,
			cmd.FormatTime(p.StartTime), cmd.FormatTime(p.EndTime))
	}

}
