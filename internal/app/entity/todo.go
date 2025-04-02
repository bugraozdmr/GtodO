package app

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	TagID       uuid.UUID `json:"tag_id" gorm:"type:uuid;not null;index"` // Tek bir tag
	Tag         Tag       `json:"tag" gorm:"foreignKey:TagID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DueDate     *time.Time `json:"due_date" gorm:"default:null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}
