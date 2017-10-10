package api

import (
	"time"
)

type Process struct {
	ErrMessage

	ID                  string
	BusinessKey         string
	StartTime           *time.Time
	EndTime             *time.Time
	ProcessDefinitionId string
}

type ProcessList struct {
	ErrMessage

	Data  []Process
	Size  int
	Total int
}

type ProcessDef struct {
	ErrMessage

	ID           string
	Key          string
	Name         string
	Version      int
	DeploymentID string
}

type ProcessDefList struct {
	ErrMessage

	Data  []ProcessDef
	Size  int
	Total int
}
