package router

import (
	"strconv"

	dto "github.com/cyber-acrobatic-group/Gatcha_API/dto/user"
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

	router.GET("/users/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			panic(err)
		}
		userService := service.NewUserService()
		user, err := userService.Show(id)
		if err != nil {
			panic(err)
		}
		ctx.JSON(200, dto.ToShowResponse(user))
	})
	router.Run(":3000")
}
