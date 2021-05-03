package Business

import (
	"main.go/Database"
	"main.go/General"
	"main.go/Models"
)

func CreateProject(project Models.Project) (bool, General.ErrorCode) {
	project.Id = General.CreateUUID()

	project.UserList = append(project.UserList, Models.UserRole{
		Email: project.CreateUser,
		Role:  "owner",
	})
	result, ErrCode := Database.CreateProject(project)
	if result {
		//create project success --> add projectlist of member
		//update project-role of user
		for _, userlist := range project.UserList {
			CurrentUser, Err := Database.GetUsers(userlist.Email)
			if Err != nil {
				return result, ErrCode
			}
			CurrentUser[0].ProjectList = append(CurrentUser[0].ProjectList, Models.ProjectList{
				ProjectId: project.Id,
				Role:      userlist.Role,
			})
			Database.UpdateProjectList(CurrentUser[0])
		}
	}

	return result, ErrCode
}
func ChangeStatusProject(project Models.Project) bool {
	return Database.ChangeStatusProject(project)
}

func SearchProject(filter string) ([]Models.Project, error) {
	return Database.SearchProject(filter)
}
func GetProjects(email string, Id string) ([]Models.Project, error) {
	if Id == "" && len(email) > 0 {
		user, err := Database.GetUsers(email)
		if err == nil && len(user) > 0 {
			listProject := user[0].ProjectList
			var listProjectResult []Models.Project
			for _, projectRole := range listProject {
				project, _ := Database.GetProject(projectRole.ProjectId)
				listProjectResult = append(listProjectResult, project[0])
			}
			return listProjectResult, nil
		}
	}
	return Database.GetProject(Id)
}
func AddMember(email string, project Models.Project) bool {
	Project, _ := GetProjects("", project.Id)
	if len(Project) > 0 {

		if !validateMember(project.UserList) {
			return false
		}
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
		if !validateMember(ParamProject.UserList) {
			return false
		}
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
		if !validateMember(ParamProject.UserList) {
			return false
		}
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

func validateMember(listmember []Models.UserRole) bool {
	//check mail is registered
	//if any not , return fail
	for _, userlist := range listmember {
		_, Err := Database.GetUsers(userlist.Email)
		if Err != nil {
			return false
		}
	}
	return true
}
