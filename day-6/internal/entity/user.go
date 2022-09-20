package entity

import (
	"context"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Token    string `json:"token"`
}

type UserRepository interface {
	GetUsers() (interface{}, error)
	GetUserByID(id string) (interface{}, error)
	CreateUsers(users User) (interface{}, error)
	UpdateUsers(user User, id string) (interface{}, error)
	DeleteUserByID(id string) (interface{}, error)
	GetUserByEmailAndPassword(email string, password string) (*User, error)
}

type UserService interface {
	LoginUsers(ctx context.Context, user *User) (*User, error)
}
