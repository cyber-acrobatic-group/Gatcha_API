package service

import (
	"net/http"

	"github.com/cyber-acrobatic-group/Gatcha_API/dto"
	"github.com/cyber-acrobatic-group/Gatcha_API/repository"
	"github.com/gin-gonic/gin"
)

// 処理のロジックを書く（DBとデータのやりとりを行う部分はmysqlパッケージに書く）
func GetMessage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

type UserService interface {
	Create(dto.UserCreateForm) error
}

type User struct {
	userRepository repository.User
}

func NewUserService() UserService {
	return &User{}
}

func (u *User) Create(form dto.UserCreateForm) error {
	user := form.ConvertToUser()
	if err := repository.Create(user); err != nil {
		return err
	}
	return nil
}
