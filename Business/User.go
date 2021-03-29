package Business

import (
	"database/sql"
	"main.go/Database"
	"main.go/Models"
)

//login
func Login(db *sql.DB, username string, pass string) (bool, bool, Models.User) {
	return Database.Login(db, username, pass)
}

//
////register
//func Register(db *sql.DB, user MODELS.RequestRegister) (bool, error) {
//	return DATABASE.Register(db, user)
//}
//
////get all username
//func GetAllUserName(db *sql.DB) []string {
//	return DATABASE.GetAllUserName(db)
//}
//
////get a user or all users
//func GetUsers(db *sql.DB, Id int) []MODELS.USERS {
//	return DATABASE.GetUsers(db, Id)
//}
