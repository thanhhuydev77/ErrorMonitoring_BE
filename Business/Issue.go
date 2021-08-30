package Business

import (
	"github.com/pkg/errors"
	"log"
	"main.go/CONST"
	"main.go/Database"
	"main.go/General"
	"main.go/Models"
	"strconv"
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
		//get list of issues
		//suggest assignee
		if project[0].AutoSuggestPerson {
			listMember := project[0].UserList
			//add author - 0.5K
			for i, val := range listMember {
				if val.NameInProduct == issue.Assignee {
					listMember[i].TimeEstimate -= val.Ability * 0.5
				}
			}
			//get min T
			assignee := Models.GetMemberMinT(listMember)
			//set assignee
			issue.Assignee = assignee
		} else {
			issue.Assignee = ""
		}

		//or not
		//suggest how to fix

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
	if project[0].EnableMailNotification {
		MailToAdminAndOwner(project[0])
		if issue.Assignee != "" {
			MailToAssignee(project[0], issue)
		}
	}
	return result, ErrMessage
}

func MailToAdminAndOwner(project Models.Project) {
	log.Print("Start Mailing to admin and owner")
	Object := "[" + project.Name + "]" + "New Issue Created"
	Text := Models.AppConfig.UILink + Models.AppConfig.IssuesPath
	for _, val := range project.UserList {
		if val.Role == "owner" || val.Role == "admin" {
			Name := ""
			user, err := Database.GetUsers(val.Email)
			if err != nil || len(user) == 0 {
				Name = "Sir/Madam"
			} else {
				Name = user[0].FullName
			}
			result := General.SendMail(val.Email, Object, Text, Name, "MailToAdminAndOwner.html")

			log.Print("mail to " + val.Email + " result:" + strconv.FormatBool(result))
		}
	}
	log.Print("finish mailing to admin and owner")
}

func MailToAssignee(project Models.Project, issue Models.Issue) {
	log.Print("Start mailing to assignee")
	Object := "[" + project.Name + "]" + "New Issue Created"
	Text := Models.AppConfig.UILink + Models.AppConfig.IssuesPath
	Name := ""
	user, err := Database.GetUsers(issue.Assignee)
	if err != nil || len(user) == 0 {
		Name = "Sir/Madam"
	} else {
		Name = user[0].FullName
	}
	result := General.SendMail(issue.Assignee, Object, Text, Name, "MailToAssignee.html")
	log.Print("mail to " + issue.Assignee + " result:" + strconv.FormatBool(result))
	log.Print("finish mail to assignee ")
}

func checkIssueExisted(issues []Models.Issue, issue Models.Issue) bool {
	for _, Issue := range issues {
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
	oldAssignee := ""
	newAssignee := issue.Assignee
	for i := range project[0].Issues {
		if project[0].Issues[i].Id == issue.Id {
			project[0].Issues[i].StartDate = issue.StartDate
			project[0].Issues[i].Assignee = issue.Assignee
			project[0].Issues[i].DueDate = issue.DueDate
			project[0].Issues[i].Priority = issue.Priority
			project[0].Issues[i].Status = issue.Status
			oldAssignee = project[0].Issues[i].Assignee
			haveRec = true
			break
		}
	}
	if haveRec {
		result = Database.UpdateIssueList(project[0])
		if oldAssignee != newAssignee {
			go UpdateKAndT(project[0], oldAssignee)
		}
		go UpdateKAndT(project[0], newAssignee)

	}
	return result
}

func UpdateKAndT(project Models.Project, assignee string) {
	//update K
	Database.UpdateAbility(project, assignee)
	//update T
	Database.UpdateTimeEstimate(project, assignee)
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
