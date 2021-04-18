package Business

import (
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
