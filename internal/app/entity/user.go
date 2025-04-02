package app

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRegister struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Username  string    `json:"username" gorm:"not null;unique"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
