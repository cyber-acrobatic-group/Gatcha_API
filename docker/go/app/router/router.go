package router

import (
	"fmt"
	"github.com/cyber-acrobatic-group/Gatcha_API/handler"
	"github.com/gin-gonic/gin"
)

func Setting() {
	fmt.Println("r")
	r := gin.Default()
	r.POST("/photo", func(context *gin.Context) {
		fmt.Println("hogehoge")
	})
	r.POST("/user", handler.PostUsers)
	r.GET("/users/:Id", RequirePathParamStr("Id"), handler.GetUsers)
	r.Run(":3000")

}

func RequirePathParamStr(param string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(param, c.Param(param))
		c.Next()
	}
}
