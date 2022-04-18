package service

import (
	dto "github.com/cyber-acrobatic-group/Gatcha_API/dto/user"
	"github.com/cyber-acrobatic-group/Gatcha_API/model"
	"github.com/cyber-acrobatic-group/Gatcha_API/repository"
)

type UserService interface {
	Create(form dto.UserCreateForm) error
	Show(id uint64) (model.User, error)
}

type User struct {
	userRepository repository.User
}

func NewUserService() UserService {
	return &User{}
}

func (u *User) Create(form dto.UserCreateForm) error {
	user := form.ToUser()
	if err := repository.Create(user); err != nil {
		return err
	}
	return nil
}

func (u *User) Show(id uint64) (model.User, error) {
	user, err := repository.Show(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
