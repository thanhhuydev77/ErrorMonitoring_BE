package Database

import (
	"database/sql"
)

const (
	APP_KEY = "thisissecreckeyyesitisreallyofcourcetrustmeitiskeyofthisapphahaha"
	//CONNECTION_STRING = "thanhhuydz123:_doan123@tcp(db4free.net:3306)/rooms_con?parseTime=true"
	CONNECTION_STRING = "root:s2hautjeuthu@tcp(127.0.0.1:3306)/room_schema?parseTime=true"
	DRIVER_NAME       = "mysql"
	REDISURL          = "redis://localhost:6379"
)

var Db *sql.DB = nil

//get Database instance
func GetDbInstance() *sql.DB {
	if Db == nil {
		Db, _ = connectdatabase()
	}
	return Db
}

//connect to database
func connectdatabase() (*sql.DB, error) {
	db, err := sql.Open(DRIVER_NAME, CONNECTION_STRING)
	if err != nil {
		return nil, nil
	}
	if err := db.Ping(); err != nil {
		return nil, nil
	}
	return db, nil
}
