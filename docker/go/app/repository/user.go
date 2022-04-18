package repository

import (
	"github.com/cyber-acrobatic-group/Gatcha_API/model"
	"github.com/cyber-acrobatic-group/Gatcha_API/mysql"
)

type User interface {
	Create(user model.User) error
	Show(id uint64) (model.User, error)
}

func Create(user model.User) error {
	db, err := mysql.GormConnect()
	if err != nil {
		return err
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func Show(id uint64) (model.User, error) {
	db, err := mysql.GormConnect()
	if err != nil {
		return model.User{}, err
	}
	var targetUser model.User
	if err := db.First(&targetUser, id).Error; err != nil {
		return model.User{}, err
	}
	return targetUser, nil
}
