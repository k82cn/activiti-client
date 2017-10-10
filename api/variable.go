package api

type Variable struct {
	Name  string `json:"name"`
	Scope string `json:"scope"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type VariableData struct {
	ErrMessage

	ID                string
	Variable          Variable
	ProcessInstanceId string
}

type VariableList struct {
	ErrMessage

	Data  []VariableData
	Size  int
	Total int
}
