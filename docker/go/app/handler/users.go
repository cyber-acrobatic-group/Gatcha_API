package handler

import (
	"fmt"
	"github.com/cyber-acrobatic-group/Gatcha_API/model"
	"github.com/cyber-acrobatic-group/Gatcha_API/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostUsers(ctx *gin.Context) {
	var input *model.Users
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	output, err := service.PostUsers(*input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, output)

}
func GetUsers(ctx *gin.Context) {

	userId, err := strconv.Atoi(ctx.MustGet("Id").(string))
	fmt.Println("", userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	output, err := service.GetUsers(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, output)

}
