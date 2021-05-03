package Business

import (
	"errors"
	"main.go/Database"
	"main.go/General"
	"main.go/Models"
)

//login
func Login(username string, pass string) (bool, bool) {
	return Database.Login(username, pass)
}

//
////register
func Register(user Models.User) (bool, General.ErrorCode) {
	return Database.Register(user)
}

//update
func Update(user Models.User) bool {
	return Database.Update(user)

}

////get a user or all users
func GetUsers(Id string) ([]Models.User, error) {
	return Database.GetUsers(Id)
}

func CheckUserExsist(email string) bool {
	user, _ := Database.GetUsers(email)
	if len(user) == 0 {
		return false
	}
	return true
}
func GetUsersByProjectId(ProjectId string) ([]Models.User, error) {
	ProjectList, Err := GetProjects("", ProjectId)
	var CurProject Models.Project
	if Err != nil || len(ProjectList) == 0 {
		return nil, errors.New("Project Id is not valid!")
	}
	CurProject = ProjectList[0]
	var listUser []Models.User
	for _, user := range CurProject.UserList {
		foundUser, Err := Database.GetUsers(user.Email)
		if Err != nil || len(foundUser) == 0 {
			return nil, errors.New("User Id is not valid!")
		}
		listUser = append(listUser, foundUser[0])
	}
	return listUser, nil
}
