package router

import (
	"fmt"

	"github.com/cyber-acrobatic-group/Gatcha_API/controller/mysql"
	"github.com/cyber-acrobatic-group/Gatcha_API/controller/service"
	"github.com/cyber-acrobatic-group/Gatcha_API/model"
	"github.com/gin-gonic/gin"
)

func Setting() {
	r := gin.Default()
	// 応答確認用
	r.GET("/ping", func(c *gin.Context) {
		mysql.ConnectDB()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// ユーザー情報登録
	r.POST("/user", func(c *gin.Context) {
		user := model.User{}
		user.Name = c.Query("name")
		user.Password = c.Query("password")
		user.Count = c.Query("count")
		res := service.RegisterUser(&user)
		c.JSON(200, gin.H {
			"message" : res,
		})
	})
	// ユーザー情報取得
	r.GET("/user", func(c *gin.Context) {
		id := c.Query("id")
		user := model.User{}
		res := service.GetUserbyId(id, &user)
		if res != "" {
			fmt.Println(res)
			return
		}
		c.JSON(200, gin.H{
			"name":  user.Name,
			"pass":  user.Password,
			"count": user.Count,
		})
	})

	r.Run(":3000")
}
