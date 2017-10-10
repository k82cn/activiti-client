package api

type Query struct {
	ProcessInstanceId   *string `json:"processInstanceId,omitempty"`
	ProcessDefinitionId *string `json:"processDefinitionId,omitempty"`
}
