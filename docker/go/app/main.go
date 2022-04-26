package main

import (
	"fmt"

	"github.com/cyber-acrobatic-group/Gatcha_API/service"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!!")

	//router.Setting()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/user", func(c *gin.Context) {
		service.RespUser(c)
	})

	r.GET("/user", func(c *gin.Context) {
		id := c.Query("id")
		res := service.GetUserbyId(id)
		c.JSON(200, gin.H{
			"result": res,
		})
	})

	r.Run(":3000")
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
