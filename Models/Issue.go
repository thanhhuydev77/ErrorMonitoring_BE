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
	StartDate   time.Time `json:"startDate" bson:"startDate"`
	DueDate     time.Time `json:"dueDate" bson:"dueDate"`
	Priority    string    `json:"priority" bson:"priority"`
	Detail      string    `json:"detail" bson:"detail"`
	CreateTime  time.Time `json:"createTime" bson:"createTime"`
	Reviewer    string    `json:"reviewer" bson:"reviewer"`
	Comment     []Comment `json:"comment" bson:"comment"`
}

type IssueForCal struct {
	Id          string
	Environment float64
	Status      float64
	Assignee    string
	StartDate   time.Time
	DueDate     time.Time
	Priority    float64
}

type Comment struct {
	Email string `json:"email" bson:"email"`
	Text  string `json:"text" bson:"text"`
	Like  int    `json:"like" bson:"like"`
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
type UserTask struct {
	Email    string  `json:"email"`
	TaskRate float64 `json:"taskRate" bson:"taskRate"`
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

//func ConvertListIssueForCalc(issues []Issue) []IssueForCal{
//	var result []IssueForCal
//	for _,issue := range issues {
//		result = append(result, ConvertIssueForCalc(issue))
//	}
//	return result
//}
func ConvertIssueForCalc(Issue Issue) IssueForCal {
	var IssueCal IssueForCal
	IssueCal.Id = Issue.Id
	IssueCal.Assignee = Issue.Assignee
	IssueCal.StartDate = Issue.StartDate
	IssueCal.DueDate = Issue.DueDate

	switch Issue.Status {
	case "unresolved":
		IssueCal.Status = 0
		break
	case "processing":
		IssueCal.Status = 50
		break
	case "reviewing":
		IssueCal.Status = 90
		break
	case "resolved":
		IssueCal.Status = 100
		break
	}

	switch Issue.Environment {
	case "development":
		IssueCal.Environment = 1
		break
	case "staging":
		IssueCal.Environment = 1.5
		break
	case "production":
		IssueCal.Environment = 2
		break
	}

	switch Issue.Priority {
	case "high":
		IssueCal.Priority = 3
		break
	case "medium":
		IssueCal.Priority = 2
		break
	case "low":
		IssueCal.Priority = 1
		break
	}

	return IssueCal
}
