package pgdb

import (
	"context"

	"github.com/grum261/event-calendar/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type City struct {
	q *Queries
}

func newCity(db *pgxpool.Pool) *City {
	return &City{
		q: newQueries(db),
	}
}

func (c *City) Create(ctx context.Context, name string, timezone int) (int, error) {
	return c.q.cityInsert(ctx, name, timezone)
}

func (c *City) Update(ctx context.Context, id, timezone int, name string) error {
	return c.q.cityUpdate(ctx, id, timezone, name)
}

func (c *City) Delete(ctx context.Context, id int) error {
	return c.q.cityDelete(ctx, id)
}

func (c *City) GetAll(ctx context.Context) ([]models.City, error) {
	cities, err := c.q.citiesSelectAll(ctx)
	if err != nil {
		return nil, err
	}

	var out []models.City

	for _, c := range cities {
		out = append(out, models.City{
			Id:       c.Id,
			Name:     c.Name,
			Timezone: c.Timezone,
		})
	}

	return out, nil
}
