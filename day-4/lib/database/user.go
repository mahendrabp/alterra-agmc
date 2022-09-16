package database

import (
	"github.com/mahendrabp/alterra-agmc/day-4/config"
	"github.com/mahendrabp/alterra-agmc/day-4/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUserByID(id string) (interface{}, error) {
	var users models.User

	if e := config.DB.Where("id = ?", id).First(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func CreateUsers(users models.User) (interface{}, error) {
	if e := config.DB.Create(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func UpdateUsers(user models.User, id string) (interface{}, error) {
	var users models.User

	err := config.DB.First(&users, id).Error
	if err != nil {
		return users, err
	}

	dberr := config.DB.
		Model(&users).
		Set("gorm:save_associations", false).
		Updates(map[string]interface{}{
			"name":     user.Name,
			"email":    user.Email,
			"password": user.Password,
		}).Error

	if dberr != nil {
		return users, err
	}

	return users, nil

}

func DeleteUserByID(id string) (interface{}, error) {
	var users models.User

	err := config.DB.First(&users, id).Error
	if err != nil {
		return users, err
	}

	if e := config.DB.Where("id = ?", id).Delete(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}
