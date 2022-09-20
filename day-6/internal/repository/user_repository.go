package repository

import (
	"github.com/mahendrabp/alterra-agmc/day-6/internal/entity"
	"gorm.io/gorm"
)

type user struct {
	Db *gorm.DB
}

func NewUser(db *gorm.DB) entity.UserRepository {
	return &user{
		db,
	}
}

func (u *user) GetUsers() (interface{}, error) {
	var users []entity.User

	if e := u.Db.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func (u *user) GetUserByID(id string) (interface{}, error) {
	var users entity.User

	if e := u.Db.Where("id = ?", id).First(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func (u *user) GetUserByEmailAndPassword(email string, password string) (*entity.User, error) {
	var user *entity.User

	if err := u.Db.Where("email = ? AND password = ?", &email, &password).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *user) CreateUsers(users entity.User) (interface{}, error) {
	if e := u.Db.Create(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func (u *user) UpdateUsers(user entity.User, id string) (interface{}, error) {
	var users entity.User

	err := u.Db.First(&users, id).Error
	if err != nil {
		return users, err
	}

	dberr := u.Db.
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

func (u *user) DeleteUserByID(id string) (interface{}, error) {
	var users entity.User

	err := u.Db.First(&users, id).Error
	if err != nil {
		return users, err
	}

	if e := u.Db.Where("id = ?", id).Delete(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}
