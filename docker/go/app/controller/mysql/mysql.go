package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	USER := "root"
	PASS := "root"
	ADDRESS := "172.29.0.3:3306"
	DBNAME := "gatcha"

	dbconf := USER + ":" + PASS + "@tcp(" + ADDRESS + ")/" + DBNAME + "?charset=utf8mb4"
	db, err := sql.Open("mysql", dbconf)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	// 接続が終了したらクローズする
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗" + dbconf)
	} else {
		fmt.Println("データベース接続成功")
	}
	return db
}
