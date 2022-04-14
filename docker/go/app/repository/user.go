package repository

import (
	"github.com/cyber-acrobatic-group/Gatcha_API/model"
	"github.com/cyber-acrobatic-group/Gatcha_API/mysql"
)

type User interface {
	Create(user model.User) error
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
