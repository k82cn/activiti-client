package api

type Execution struct {
	ErrMessage

	ID         string
	URL        string
	ActivityID string
}

type ExecutionList struct {
	ErrMessage

	Data  []Execution
	Size  int
	Total int
}
