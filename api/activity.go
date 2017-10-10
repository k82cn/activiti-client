package api

import (
	"time"
)

type Activity struct {
	ErrMessage

	ID         string
	ActivityID string
	Assignee   string
	StartTime  *time.Time
	EndTime    *time.Time
}

type ActivityList struct {
	ErrMessage

	Data  []Activity
	Size  int
	Total int
}
