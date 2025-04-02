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
	Tags        []Tag     `json:"tags" gorm:"many2many:todo_tags;"`
	Completed   bool      `json:"completed" gorm:"default:false"`
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
