package Models

import (
	"main.go/CONST"
	"strings"
	"time"
)

type Issue struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Environment string    `json:"environment"`
	Frame       string    `json:"frame"`
	Status      string    `json:"status"`
	Assignee    string    `json:"assignee"`
	Path        string    `json:"path"`
	DueDate     time.Time `json:"dueDate"`
	Priority    string    `json:"priority"`
	Detail      string    `json:"detail"`
	CreateTime  time.Time `json:"createTime"`
}

type IssueRequest struct {
	Type  string `json:"type"`
	Issue Issue  `json:"Issue"`
}

type IssueFilter struct {
	ProjectId   string    `json:"projectId"`
	Assignee    string    `json:"assignee"`
	Environment string    `json:"environment"`
	FromDate    time.Time `json:"fromDate"`
	ToDate      time.Time `json:"toDate"`
}

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {

	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(CONST.TIMEFORMAT, s)
	return
}

type CustomTime struct {
	time.Time
}

func FilterIssue(listIssue []Issue, field string, value string) []Issue {
	var result []Issue
	switch field {
	case "Assignee":
		for _, issue := range listIssue {
			if issue.Assignee == value {
				result = append(result, issue)
			}
		}
		break
	case "Environment":
		for _, issue := range listIssue {
			if issue.Environment == value {
				result = append(result, issue)
			}
		}
		break
	case "FromDate":
		a, err := time.Parse(CONST.TIMEFORMAT, value)
		if err == nil {
			for _, issue := range listIssue {
				if issue.CreateTime.After(a) {
					result = append(result, issue)
				}
			}
		}
		break
	case "ToDate":
		a, err := time.Parse(CONST.TIMEFORMAT, value)
		if err == nil {
			for _, issue := range listIssue {
				if issue.CreateTime.Before(a) {
					result = append(result, issue)
				}
			}
		}
		break
	}
	return result
}
