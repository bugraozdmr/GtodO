package usecase

import (
	"context"
	"errors"
	entity "gtodo/internal/app/entity"
	"gtodo/internal/app/repository"
)

type TagUseCase interface {
	CreateTag(ctx context.Context, tag *entity.Tag) error
	GetAllTags(ctx context.Context) ([]entity.Tag, error)
	DeleteTag(ctx context.Context, id string) error
}

type TagInteraction struct {
	TagRepository repository.TagRepository
}

func (t *TagInteraction) CreateTag(ctx context.Context, tag *entity.Tag) error {
	err := t.TagRepository.CreateTag(tag)
	if err != nil {
		return err
	}

	return nil
}

func (t *TagInteraction) GetAllTags(ctx context.Context) ([]entity.Tag, error) {
	tags, err := t.TagRepository.GetAllTags()
	if err != nil {
		return nil, errors.New("could not retrieve tags")
	}

	return tags, nil
}

func (t *TagInteraction) DeleteTag(ctx context.Context, id string) error {
	err := t.TagRepository.DeleteTag(id)
	if err != nil {
		return errors.New("failed to delete tag")
	}

	return nil
}

func UseCaseTag(repo repository.TagRepository) TagUseCase {
	return &TagInteraction{
		TagRepository: repo,
	}
}
