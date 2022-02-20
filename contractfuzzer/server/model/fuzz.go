package model

import "time"

type FuzzStartRequest struct {
	Contracts []string      `json:"contracts"`
	Duration  time.Duration `json:"duration"`
}

type FuzzStartResponse struct {
	TaskId string `json:"taskId"`
}
