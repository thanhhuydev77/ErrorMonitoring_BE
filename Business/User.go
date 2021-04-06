package Business

import (
	"main.go/Database"
	"main.go/Models"
)

//login
func Login(username string, pass string) (bool, bool) {
	return Database.Login(username, pass)
}

//
////register
func Register(user Models.User) (bool, Database.ErrorCode) {
	return Database.Register(user)
}

//
////get all username
//func GetAllUserName(db *sql.DB) []string {
//	return DATABASE.GetAllUserName(db)
//}
//
////get a user or all users
func GetUsers(Id string) ([]Models.User, error) {
	return Database.GetUsers(Id)
}
