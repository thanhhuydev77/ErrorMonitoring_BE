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
func ModifyMember(email string, project Models.Project) bool {
	Project, _ := GetProjects("", project.Id)
	if len(Project) > 0 {
		for i, userlist := range Project[0].UserList {
			if userlist.Email == project.UserList[0].Email {
				Project[0].UserList[i].Role = project.UserList[0].Role
				return Database.UpdateUserList(Project[0])
			}
		}
		//return Database.UpdateUserList(Project[0])
	}
	return false
}
func RemoveMember(email string, project Models.Project) bool {
	Project, _ := GetProjects("", project.Id)
	if len(Project) > 0 {
		for i, userlist := range Project[0].UserList {
			if userlist.Email == project.UserList[0].Email {
				Project[0].UserList[i] = Project[0].UserList[len(Project[0].UserList)-1]
				Project[0].UserList = Project[0].UserList[:len(Project[0].UserList)-1]
				return Database.UpdateUserList(Project[0])
			}
		}
		//return Database.UpdateUserList(Project[0])
	}
	return false
}
