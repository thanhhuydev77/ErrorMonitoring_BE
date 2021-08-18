package Business

import (
	"github.com/pkg/errors"
	"main.go/CONST"
	"main.go/Database"
	"main.go/General"
	"main.go/Models"
	"time"
)

func CreateIssue(ProjectId string, issue Models.Issue) (bool, string) {
	issue.Id = General.CreateUUID()
	result := false
	ErrMessage := ""
	project, Err := Database.GetProjectWithIssue(ProjectId)
	if Err != nil || len(project) == 0 {
		return false, "ProjectId is invalid"
	}
	if !checkIssueExisted(project[0].Issues, issue) {
		project[0].Issues = append(project[0].Issues, issue)
		result = Database.UpdateIssueList(project[0])
		//integrate Trello and Slack
		if !result {
			return result, "Update Issue failed"
		}
		if project[0].EnableTrello {
			TrelloMsg := ""
			_, TrelloMsg = General.TrelloCreateCard(project[0].TrelloInfo.AppToken, project[0].TrelloInfo.UserID, project[0].TrelloInfo.BoardID, project[0].TrelloInfo.ListID, issue)
			if TrelloMsg != "" {
				ErrMessage += TrelloMsg
			}
		}
		if project[0].EnableSlack {
			SlackMsg := ""
			_, SlackMsg = General.SlackCreateNortification(project[0].SlackInfo.BotToken, project[0].SlackInfo.ChanelId, issue)
			if SlackMsg != "" {
				ErrMessage += "," + SlackMsg
			}
		}
	} else {
		return false, "Duplicate Issue"
	}

	return result, ErrMessage
}

func checkIssueExisted(issuelist []Models.Issue, issue Models.Issue) bool {
	for _, Issue := range issuelist {
		if Issue.CheckCode == issue.CheckCode {
			if Issue.Status == CONST.PROCESSING || Issue.Status == CONST.UNRESOLVED {
				return true
			}
		}
	}
	return false
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

func UpdateIssue(ProjectId string, issue Models.Issue) bool {
	result := false
	haveRec := false
	project, Err := Database.GetProjectWithIssue(ProjectId)
	if Err != nil || len(project) == 0 {
		return false
	}
	for i := range project[0].Issues {
		if project[0].Issues[i].Id == issue.Id {
			project[0].Issues[i].Assignee = issue.Assignee
			project[0].Issues[i].DueDate = issue.DueDate
			project[0].Issues[i].Priority = issue.Priority
			project[0].Issues[i].Status = issue.Status
			haveRec = true
			break
		}
	}
	if haveRec {
		result = Database.UpdateIssueList(project[0])
	}
	return result
}

func GetIssue(ProjectId string, Id string) (Models.Issue, bool) {

	project, Err := Database.GetProjectWithIssue(ProjectId)
	if Err != nil || len(project) == 0 {
		return Models.Issue{}, false
	}
	for _, iss := range project[0].Issues {
		if iss.Id == Id {
			return iss, true
		}
	}
	return Models.Issue{}, false
}
