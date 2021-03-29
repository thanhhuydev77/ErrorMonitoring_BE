package Database

import (
	"database/sql"
)

var Db *sql.DB = nil

//get Database instance
func GetDbInstance() *sql.DB {
	return nil
}

//connect to database
func connectdatabase() (*sql.DB, error) {
	return nil, nil
}
