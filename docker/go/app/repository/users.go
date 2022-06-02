package repository

import (
	"github.com/cyber-acrobatic-group/Gatcha_API/model"
	"github.com/cyber-acrobatic-group/Gatcha_API/mysql"
)

func PostUsers(user model.Users) (*model.Users, error) {

	db, err := mysql.GormConnect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func GetUsers(id int) (model.Users, error) {
	db, err := mysql.GormConnect()
	if err != nil {
		return model.Users{}, err
	}
	var targetUser model.Users
	if err := db.First(&targetUser, id).Error; err != nil {
		return model.Users{}, err
	}
	return targetUser, nil
}
