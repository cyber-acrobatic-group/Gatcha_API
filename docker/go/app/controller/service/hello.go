package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 処理のロジックを書く（DBとデータのやりとりを行う部分はmysqlパッケージに書く）
func GetMessage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
