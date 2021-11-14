package postgresql

import (
	"context"

	"github.com/grum261/event-calendar/internal/models"
	"github.com/grum261/event-calendar/internal/postgresql/db"
	"github.com/jackc/pgx/v4/pgxpool"
)

type City struct {
	q *db.Queries
}

func newCity(pool *pgxpool.Pool) *City {
	return &City{
		q: db.NewQueries(pool),
	}
}

func (c *City) Create(ctx context.Context, name string, timezone int) (int, error) {
	return c.q.CityInsert(ctx, name, timezone)
}

func (c *City) Update(ctx context.Context, id, timezone int, name string) error {
	return c.q.CityUpdate(ctx, id, timezone, name)
}

func (c *City) Delete(ctx context.Context, id int) error {
	return c.q.CityDelete(ctx, id)
}

func (c *City) GetAll(ctx context.Context) ([]models.City, error) {
	cities, err := c.q.CitiesSelectAll(ctx)
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
