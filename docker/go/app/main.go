package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dbconf := "root:root@tcp(172.27.0.3:3306)/gatcha?charset=utf8mb4"

	db, err := sql.Open("mysql", dbconf)

	// 接続が終了したらクローズする
	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功")
	}
	// fmt.Println("Hello, World!!")

	// //router.Setting()

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.POST("/user", func(c *gin.Context) {
	// 	service.RespUser(c)
	// })

	// r.GET("/user", func(c *gin.Context) {
	// 	id := c.Query("id")
	// 	res := service.GetUserbyId(id)
	// 	c.JSON(200, gin.H{
	// 		"result": res,
	// 	})
	// })

	// r.Run(":3000")
}

// func main() {
//     engine:= gin.Default()
//     engine.GET("/", func(c *gin.Context) {
//         c.JSON(http.StatusOK, gin.H{
//             "message": "hello world",
//         })
//     })
//     engine.Run(":3000")
// }
