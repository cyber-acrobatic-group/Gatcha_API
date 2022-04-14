package router

import (
	"github.com/cyber-acrobatic-group/Gatcha_API/dto"
	"github.com/cyber-acrobatic-group/Gatcha_API/service"
	"github.com/gin-gonic/gin"
)

func Setting() {
	router := gin.Default()

	router.POST("/users", func(ctx *gin.Context) {
		var form dto.UserCreateForm
		if err := ctx.Bind(&form); err != nil {
			panic(err)
		}
		userService := service.NewUserService()
		if err := userService.Create(form); err != nil {
			panic(err)
		}
		ctx.JSON(200, gin.H{
			"result": "success",
		})
	})
	router.Run(":3000")
}
