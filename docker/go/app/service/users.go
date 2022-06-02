package service

import (
	"github.com/cyber-acrobatic-group/Gatcha_API/model"
	"github.com/cyber-acrobatic-group/Gatcha_API/repository"
)

//
func PostUsers(input model.Users) (*model.Users, error) {

	return repository.PostUsers(input)
}

func GetUsers(userId int) (model.Users, error) {

	return repository.GetUsers(userId)
}
