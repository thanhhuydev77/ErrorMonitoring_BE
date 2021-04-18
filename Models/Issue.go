package Models

import "time"

type Issue struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Environment string    `json:"environment"`
	Frame       string    `json:"frame"`
	Status      string    `json:"status"`
	Assignee    string    `json:"assignee"`
	DueDate     time.Time `json:"dueDate"`
	Priority    string    `json:"priority"`
	Detail      string    `json:"detail"`
	CreateTime  time.Time `json:"createTime"`
}
