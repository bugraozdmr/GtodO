package repository

import (
	user "gtodo/internal/app/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *user.UserRegister) error
	FindUserByUserName(username string) (*user.UserRegister, error)
	FindUserIdByUserName(username string) (string, error)
}

type UserDataBaseInteraction struct {
	DB *gorm.DB
}

func (u *UserDataBaseInteraction) CreateUser(user *user.UserRegister) error {
	result := u.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserDataBaseInteraction) FindUserByUserName(username string) (*user.UserRegister, error) {
	var user *user.UserRegister
	if err := u.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDataBaseInteraction) FindUserIdByUserName(username string) (string, error) {
	var user struct {
		ID string `gorm:"column:id"`
	}

	if err := u.DB.Table("user").Select("id").Where("username = ?", username).Take(&user).Error; err != nil {
		return "", err
	}

	return user.ID, nil
}


func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserDataBaseInteraction{
		DB: db,
	}
}
