package Business

import (
	"github.com/pkg/errors"
	"main.go/Database"
	"main.go/General"
	"main.go/Models"
	"time"
)

func CreateIssue(ProjectId string, issue Models.Issue) bool {
	issue.Id = General.CreateUUID()

	project, Err := Database.GetProjectWithIssue(ProjectId)
	if Err != nil || len(project) == 0 {
		return false
	}
	project[0].Issues = append(project[0].Issues, issue)

	result := Database.UpdateIssueList(project[0])

	return result
}

func FilterIssue(filter Models.IssueFilter) ([]Models.Issue, error) {
	var listIssue []Models.Issue
	if filter.ProjectId == "" {
		return nil, errors.Errorf("ProjectID empty")
	}
	projectList, err := Database.GetProjectWithIssue(filter.ProjectId)
	if err != nil || len(projectList) == 0 {
		return nil, errors.Errorf("ProjectID invalid")
	}
	listIssue = projectList[0].Issues

	if filter.Assignee != "" {
		listIssue = Models.FilterIssue(listIssue, "Assignee", filter.Assignee)
	}
	if filter.Environment != "" {
		listIssue = Models.FilterIssue(listIssue, "Environment", filter.Environment)
	}
	defaultTime := time.Time{}
	if filter.FromDate != defaultTime {
		listIssue = Models.FilterIssue(listIssue, "FromDate", filter.FromDate.Format(time.RFC3339))
	}
	if filter.ToDate != defaultTime {
		listIssue = Models.FilterIssue(listIssue, "ToDate", filter.ToDate.Format(time.RFC3339))
	}

	return listIssue, nil
}
