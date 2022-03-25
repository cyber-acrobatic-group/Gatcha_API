package main

import (
	"fmt"

	"github.com/cyber-acrobatic-group/Gatcha_API/router"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!!")

	router.Setting()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
