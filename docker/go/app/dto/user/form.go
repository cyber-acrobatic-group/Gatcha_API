package dto

import "github.com/cyber-acrobatic-group/Gatcha_API/model"

type UserCreateForm struct {
	Name     string `json:"name" binding:"required,max=100"`
	Password string `json:"password" binding:"required,max=100"`
}

func (u UserCreateForm) ToUser() model.User {
	return model.User{
		Name:     u.Name,
		Password: u.Password,
	}
}
