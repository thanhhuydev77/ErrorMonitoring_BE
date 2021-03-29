package Database

import (
	"database/sql"
	_ "database/sql"
	_ "log"
	"main.go/Models"
	_ "strconv"
)

//login
func Login(db *sql.DB, username string, pass string) (bool, bool, Models.User) {
	exsist := false
	passOK := false
	a := Models.User{}

	return exsist, passOK, a
}

//register a new user
//func Register(db *sql.DB, user MODELS.RequestRegister) (bool, error) {
//
//	if db == nil {
//		log.Print("can not connect to database!")
//		return false, fmt.Errorf("can not connect to database")
//	}
//	passhash, _ := hashPassword(user.Pass)
//	rows, err := db.Query(`insert into USERS(userName,Pass,FullName,Address,Role,Sex,Province,Email) values(?,?,?,?,?,?,?,?)`, user.UserName, passhash, user.FullName, user.Address, user.Role, user.Sex, user.Province, user.Email)
//	if err != nil {
//		fmt.Print("loi:" + err.Error())
//		return false, err
//	}
//
//	defer rows.Close()
//	return true, nil
//}
//
////get all user name
//func GetAllUserName(db *sql.DB) []string {
//	var Allusername []string
//	//db, err := connectdatabase()
//	// Query all users
//	if db == nil {
//
//		log.Print("can not connect to database!")
//		return nil
//	}
//	//defer db.Close()
//
//	rows, err := db.Query("select username from USERS")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for rows.Next() {
//		var username string
//		err := rows.Scan(&username)
//		if err != nil {
//			log.Fatal(err)
//		}
//		Allusername = append(Allusername, username)
//	}
//	defer rows.Close()
//	return Allusername
//}
//
////hash password by bycript
//func hashPassword(password string) (string, error) {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
//	return string(bytes), err
//}
//
////check password hash
//func checkPasswordHash(password, hash string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//	return err == nil
//}
//
////get a user or all user(id = -1)
//func GetUsers(db *sql.DB, Id int) []MODELS.USERS {
//
//	//db, err := connectdatabase()
//	//// Query all users
//	if db == nil {
//
//		log.Print("can not connect to database!")
//		return nil
//	}
//	//defer db.Close()
//	query := ""
//	list := []MODELS.USERS{}
//
//	if Id == -1 {
//		query = "select * from USERS"
//	} else {
//		query = "select * from USERS where Id = " + strconv.Itoa(Id)
//	}
//
//	rows, err := db.Query(query)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for rows.Next() {
//		var user MODELS.USERS
//		err := rows.Scan(&user.Id, &user.UserName, &user.Pass, &user.FullName, &user.IdentifyFront, &user.IdentifyBack, &user.DateBirth, &user.Address,
//			&user.Role, &user.Sex, &user.Job, &user.WorkPlace, &user.TempReg, &user.Province, &user.Email, &user.Avatar, &user.PhoneNumber)
//		if err != nil {
//			log.Fatal(err)
//		}
//		list = append(list, user)
//	}
//	return list
//}
