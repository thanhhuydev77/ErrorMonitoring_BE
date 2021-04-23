package Business

import (
	"main.go/Database"
	"main.go/General"
	"main.go/Models"
)

func CreateProject(project Models.Project) (bool, General.ErrorCode) {
	return Database.CreateProject(project)
}
func ChangeStatusProject(project Models.Project) bool {
	return Database.ChangeStatusProject(project)
}

func GetProjects(email string, Id string) ([]Models.Project, error) {
	return Database.GetProjects(email, Id)
}
func AddMember(email string, project Models.Project) bool {
	Project, _ := GetProjects("", project.Id)
	if len(Project) > 0 {
		for _, userlist := range project.UserList {
			Project[0].UserList = append(Project[0].UserList, userlist)
		}
		return Database.UpdateUserList(Project[0])
	}
	return false
}
