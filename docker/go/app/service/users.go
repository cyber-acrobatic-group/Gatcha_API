package service

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func RespUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "alterd",
	})
}

func GetUserbyId(id string) string {
	//return id + "1"

	return "aaa"
}
