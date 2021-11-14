package postgresql

import (
	"context"

	"github.com/grum261/event-calendar/internal/models"
	"github.com/grum261/event-calendar/internal/postgresql/db"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Tag struct {
	q *db.Queries
}

func newTag(pool *pgxpool.Pool) *Tag {
	return &Tag{
		q: db.NewQueries(pool),
	}
}

func (t *Tag) Create(ctx context.Context, names []string) ([]int, error) {
	ids, err := t.q.TagsInsert(ctx, names)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (t *Tag) Update(ctx context.Context, id int, name string) error {
	return t.q.TagNameUpdate(ctx, id, name)
}

func (t *Tag) Delete(ctx context.Context, id int) error {
	return t.q.TagDelete(ctx, id)
}

func (t *Tag) GetAll(ctx context.Context) ([]models.Tag, error) {
	tags, err := t.q.TagsSelectAll(ctx)
	if err != nil {
		return nil, err
	}

	var out []models.Tag

	for _, t := range tags {
		out = append(out, models.Tag{
			Id:   t.Id,
			Name: t.Name,
		})
	}

	return out, nil
}
