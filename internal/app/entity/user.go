package app

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRegister struct {
	ID        uuid.UUID `json:"-" gorm:"type:uuid;primaryKey;"`
	Username  string    `json:"username" gorm:"not null;unique"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"-" gorm:"not null;unique"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserRegister) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

func (UserRegister) TableName() string {
	return "user"
}