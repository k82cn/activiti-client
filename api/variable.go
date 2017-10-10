package api

type Variable struct {
	Name  string
	Scope string
	Type  string
	Value string
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
