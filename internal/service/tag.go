package service

import (
	"context"

	"github.com/grum261/event-calendar/internal/models"
)

type TagRepository interface {
	Create(ctx context.Context, names []string) ([]int, error)
	Update(ctx context.Context, id int, name string) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]models.Tag, error)
}

type Tag struct {
	repo TagRepository
}

func newTagRepo(repo TagRepository) *Tag {
	return &Tag{
		repo: repo,
	}
}

func (t *Tag) Create(ctx context.Context, names []string) ([]int, error) {
	return t.repo.Create(ctx, names)
}

func (t *Tag) Update(ctx context.Context, id int, name string) error {
	return t.repo.Update(ctx, id, name)
}

func (t *Tag) Delete(ctx context.Context, id int) error {
	return t.repo.Delete(ctx, id)
}

func (t *Tag) GetAll(ctx context.Context) ([]models.Tag, error) {
	return t.repo.GetAll(ctx)
}
