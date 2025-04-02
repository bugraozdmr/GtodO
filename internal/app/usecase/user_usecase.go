package usecase

import (
	"context"
	"errors"
	app "gtodo/internal/app/entity"
	"gtodo/internal/app/repository"
	validation "gtodo/internal/app/validations/user"
	"gtodo/internal/utils"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, user *app.UserRegister) (int,error)
	Login(login *app.UserLogin) (*app.UserRegister, error)
}

type UserInteraction struct {
	UserRepository repository.UserRepository
}

// if its 2 this means validation error
func (u *UserInteraction) RegisterUser(ctx context.Context, user *app.UserRegister) (int,error) {
	if err := validation.ValidateUserRegister(user); err != nil {
		return 2,err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return 1,err
	}

	user.Password = hashedPassword

	err = u.UserRepository.CreateUser(user)
	if err != nil {
		return 1,err
	}

	return 1,nil
}

func (u *UserInteraction) Login(login *app.UserLogin) (*app.UserRegister, error) {
	user, err := u.UserRepository.FindUserByUserName(login.Username)
	if err != nil {
		return nil, err
	}
	if login.Password != user.Password {
		return nil, errors.New("password error")
	}
	return user, nil
}

func UseCase(repo repository.UserRepository) UserUseCase {
	return &UserInteraction{
		UserRepository: repo,
	}
}
