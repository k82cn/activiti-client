package api

type Activity struct {
	ErrMessage

	ID         string
	ActivityID string
	Assignee   string
	StartTime  string
	EndTime    string
}

type ActivityList struct {
	ErrMessage

	Data  []Activity
	Size  int
	Total int
}
