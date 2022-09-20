package user

import (
	"context"
	"fmt"

	"github.com/mahendrabp/alterra-agmc/day-6/internal/entity"
	"github.com/mahendrabp/alterra-agmc/day-6/internal/middlewares"
)

type userService struct {
	UserRepository entity.UserRepository
}

func NewUserService(ur entity.UserRepository) entity.UserService {
	return &userService{
		UserRepository: ur,
	}
}

func (us *userService) LoginUsers(ctx context.Context, user *entity.User) (*entity.User, error) {
	userFound, err := us.UserRepository.GetUserByEmailAndPassword(user.Email, user.Password)
	if err != nil {
		fmt.Println("called inn line 23", err)
		return nil, err
	}

	fmt.Println(userFound)

	token, err := middlewares.CreateToken(int(user.ID))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	userFound.Token = token

	return userFound, nil

}
