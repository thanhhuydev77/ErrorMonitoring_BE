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

			//update project-role of user
			CurrentUser, Err := Database.GetUsers(userlist.Email)
			if Err != nil {
				return false
			}
			CurrentUser[0].ProjectList = append(CurrentUser[0].ProjectList, Models.ProjectList{
				ProjectId: Project[0].Id,
				Role:      userlist.Role,
			})
			Database.UpdateProjectList(CurrentUser[0])
		}
		return Database.UpdateUserList(Project[0])
	}
	return false
}
func ModifyMember(email string, ParamProject Models.Project) bool {
	CurrentProject, _ := GetProjects("", ParamProject.Id)
	if len(CurrentProject) > 0 {
		for i := range CurrentProject[0].UserList {
			for _, paramUserList := range ParamProject.UserList {
				if paramUserList.Email == CurrentProject[0].UserList[i].Email {
					CurrentProject[0].UserList[i].Role = paramUserList.Role

					//update project-role of user
					CurrentUser, Err := Database.GetUsers(CurrentProject[0].UserList[i].Email)
					if Err != nil {
						return false
					}
					CurrentUser[0].ProjectList = Models.ModifyProjectRole(CurrentUser[0].ProjectList, Models.ProjectList{
						ProjectId: CurrentProject[0].Id,
						Role:      CurrentProject[0].UserList[i].Role,
					})
					Database.UpdateProjectList(CurrentUser[0])
				}
			}

		}
		return Database.UpdateUserList(CurrentProject[0])
	}
	return false
}
func RemoveMember(email string, ParamProject Models.Project) bool {
	CurrentProject, _ := GetProjects("", ParamProject.Id)
	if len(CurrentProject) > 0 {
		for _, userlist := range ParamProject.UserList {
			if Models.Contain(CurrentProject[0].UserList, userlist) {
				CurrentProject[0].UserList = Models.RemoveUserRole(CurrentProject[0].UserList, userlist)
			}

			//update project-role of user
			CurrentUser, Err := Database.GetUsers(userlist.Email)
			if Err != nil {
				return false
			}
			CurrentUser[0].ProjectList = Models.RemoveProjectRole(CurrentUser[0].ProjectList, Models.ProjectList{ProjectId: ParamProject.Id})
			Database.UpdateProjectList(CurrentUser[0])
			//if userlist.Email == project.UserList[0].Email {
			//	userlist = CurrentProject[0].UserList[len(CurrentProject[0].UserList)-1]
			//	CurrentProject[0].UserList = CurrentProject[0].UserList[:len(CurrentProject[0].UserList)-1]
			//return Database.UpdateUserList(Project[0])
		}
		return Database.UpdateUserList(CurrentProject[0])
	}
	return false
}
