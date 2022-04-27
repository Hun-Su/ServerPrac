package db

import (
	"database/sql"
	"echo/Config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetDb() sql.DB {
	config := Config.LoadConfig()
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", config.Db.User+":"+config.Db.Pwd+"@tcp("+config.Db.Port+")/"+config.Db.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return *db
}
