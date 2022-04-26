package service

import (
	"github.com/gin-gonic/gin"
)

func RespUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "success",
	})
}

func GetUserbyId(id string) string {
	return id + "1"
}
