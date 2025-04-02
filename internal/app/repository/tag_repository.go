package repository

import (
	entity "gtodo/internal/app/entity"

	"gorm.io/gorm"
)

type TagRepository interface {
	CreateTag(tag *entity.Tag) error
	GetAllTags() ([]entity.Tag, error)
	DeleteTag(id string) error
}

type TagDataBaseInteraction struct {
	DB *gorm.DB
}

func (r *TagDataBaseInteraction) CreateTag(tag *entity.Tag) error {
	if err := r.DB.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

func (r *TagDataBaseInteraction) GetAllTags() ([]entity.Tag, error) {
	var tags []entity.Tag
	if err := r.DB.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *TagDataBaseInteraction) DeleteTag(id string) error {
	if err := r.DB.Delete(&entity.Tag{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &TagDataBaseInteraction{
		DB: db,
	}
}
