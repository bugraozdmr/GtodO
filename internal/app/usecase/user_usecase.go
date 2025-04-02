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
	RegisterUser(ctx context.Context, user *app.UserRegister) (int, error)
	Login(login *app.UserLogin) (*LoginOutput, error)
}

type UserInteraction struct {
	UserRepository repository.UserRepository
}

type LoginOutput struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// if its 2 this means validation error
func (u *UserInteraction) RegisterUser(ctx context.Context, user *app.UserRegister) (int, error) {
	if err := validation.ValidateUserRegister(user); err != nil {
		return 2, err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return 1, err
	}

	user.Password = hashedPassword

	err = u.UserRepository.CreateUser(user)
	if err != nil {
		return 1, err
	}

	return 1, nil
}

func (u *UserInteraction) Login(login *app.UserLogin) (*LoginOutput, error) {
	user, err := u.UserRepository.FindUserByUserName(login.Username)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(login.Password, user.Password) {
		return nil, errors.New("password error")
	}
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &LoginOutput{
		Message: "Login successful",
		Token:   token,
	}, nil
}

func UseCase(repo repository.UserRepository) UserUseCase {
	return &UserInteraction{
		UserRepository: repo,
	}
}
