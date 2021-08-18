package Models

import (
	"main.go/CONST"
	"sort"
	"strings"
	"time"
)

type Issue struct {
	Id          string    `json:"id" bson:"id"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Environment string    `json:"environment" bson:"environment"`
	CheckCode   string    `json:"checkCode" bson:"checkCode"`
	Frame       string    `json:"frame" bson:"frame"`
	Lineno      int       `json:"lineno" bson:"lineNo"`
	Colno       int       `json:"colno" bson:"colNo"`
	Status      string    `json:"status" bson:"status"`
	Assignee    string    `json:"assignee" bson:"assignee"`
	Path        string    `json:"path" bson:"path"`
	DueDate     time.Time `json:"dueDate" bson:"dueDate"`
	Priority    string    `json:"priority" bson:"priority"`
	Detail      string    `json:"detail" bson:"detail"`
	CreateTime  time.Time `json:"createTime" bson:"createTime"`
}

type IssueRequest struct {
	Type  string `json:"type" bson:"type"`
	Issue Issue  `json:"Issue" bson:"issue"`
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
		a, err := time.Parse(time.RFC3339, value)
		if err == nil {
			for _, issue := range listIssue {
				if issue.CreateTime.After(a) {
					result = append(result, issue)
				}
			}
		}
		break
	case "ToDate":
		a, err := time.Parse(time.RFC3339, value)
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

func SortIssueByCreateTime(listIssue []Issue) []Issue {
	sort.Slice(listIssue, func(i, j int) bool {
		return listIssue[i].CreateTime.After(listIssue[j].CreateTime)
	})
	return listIssue
}
