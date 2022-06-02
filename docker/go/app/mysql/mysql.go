package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GormConnect() (*gorm.DB, error) {
	DBMS := "mysql"
	USER := "docker"
	PASS := "docker"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "gatcha"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	return db, nil
}
