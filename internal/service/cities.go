package service

import (
	"context"

	"github.com/grum261/event-calendar/internal/models"
)

type CityRepository interface {
	Create(ctx context.Context, name string, timezone int) (int, error)
	Update(ctx context.Context, id, timezone int, name string) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]models.City, error)
}

type City struct {
	repo CityRepository
}

func newCity(repo CityRepository) *City {
	return &City{
		repo: repo,
	}
}

func (c *City) Create(ctx context.Context, name string, timezone int) (int, error) {
	return c.repo.Create(ctx, name, timezone)
}

func (c *City) Update(ctx context.Context, id, timezone int, name string) error {
	return c.repo.Update(ctx, id, timezone, name)
}

func (c *City) Delete(ctx context.Context, id int) error {
	return c.repo.Delete(ctx, id)
}
func (c *City) GetAll(ctx context.Context) ([]models.City, error) {
	return c.repo.GetAll(ctx)
}
