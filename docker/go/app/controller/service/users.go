package service

import (
	"fmt"

	"github.com/cyber-acrobatic-group/Gatcha_API/controller/mysql"
	"github.com/cyber-acrobatic-group/Gatcha_API/model"
	_ "github.com/go-sql-driver/mysql"
)

func GetUserbyId(id string, user *model.User) string {
	sql := "SELECT name, password, count FROM USERS WHERE id = ?"
	db := mysql.ConnectDB()
	if db == nil {
		return "DB接続に失敗しました"
	}
	err := db.QueryRow(sql, id).Scan(&user.Name, &user.Password, &user.Count)
	if err != nil {
		fmt.Println(err.Error())
		return "クエリ実行に失敗しました"
	}
	return ""
}

func RegisterUser(user *model.User) string {
	db := mysql.ConnectDB()
	defer db.Close()
	if db == nil {
		return "DB接続に失敗しました"
	}
	ins, err := db.Prepare("INSERT INTO users(name, password, count) VALUES(?,?,?)")
	if err != nil {
		return "クエリ実行に失敗しました(" + err.Error() + ")"
	}
	ins.Exec(&user.Name, user.Password, user.Count)
	return "success"
}
