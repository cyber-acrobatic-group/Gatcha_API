package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Setting() {
	fmt.Println("router")
	router := gin.Default()
	router.POST("/photo", func(context *gin.Context) {
		fmt.Println("hogehoge")

	})
	router.Run(":3000")
}
