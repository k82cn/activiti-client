package api

type Variable struct {
	ErrMessage

	Name  string
	Scope string
	Type  string
	Value string
}

type VariableList struct {
	ErrMessage

	Data  []Variable
	Size  int
	Total int
}
