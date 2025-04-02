package app

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tag struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Name  string    `json:"name" gorm:"not null;unique"`
	Color string    `json:"color" gorm:"default:blue"`
}

// UUID ve Timestamp Ayarları
func (t *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}
