package postgresql

import (
	"context"

	"github.com/grum261/event-calendar/internal/models"
	"github.com/grum261/event-calendar/internal/postgresql/db"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Event struct {
	q *db.Queries
}

func newEvent(pool *pgxpool.Pool) *Event {
	return &Event{
		q: db.NewQueries(pool),
	}
}

func (e *Event) Create(ctx context.Context, p models.EventInsertParameters) (int, []int, error) {
	tx, err := e.q.Begin(ctx)
	if err != nil {
		return 0, nil, err
	}
	defer tx.Rollback(ctx)

	eventId, err := e.q.WithTx(tx).EventInsert(ctx, db.EventInsertParams{
		Name:      p.Name,
		StartDate: p.StartDate,
		EndDate:   p.EndDate,
		URL:       p.URL,
	})
	if err != nil {
		return 0, nil, err
	}

	var partsIds []int

	for _, ep := range p.EventParts {
		partId, err := e.q.WithTx(tx).EventPartsInsert(ctx, db.EventPartInsertParams{
			EventId: eventId,
			EventPartCommonParams: db.EventPartCommonParams{
				Name:        ep.Name,
				Description: ep.Description,
				StartTime:   ep.StartTime,
				EndTime:     ep.EndTime,
				Address:     ep.Address,
				Place:       ep.Place,
			},
		})
		if err != nil {
			return 0, nil, err
		}

		partsIds = append(partsIds, partId)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, nil, err
	}

	return eventId, partsIds, nil
}

func (e *Event) Delete(ctx context.Context, id int) error {
	return e.q.EventDelete(ctx, id)
}

func (e *Event) Update(ctx context.Context, eventId int, p models.EventUpdateParameters) error {
	tx, err := e.q.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := e.q.WithTx(tx).EventUpdate(ctx, eventId, db.EventInsertParams{
		Name:      p.Name,
		StartDate: p.StartDate,
		EndDate:   p.EndDate,
		URL:       p.URL,
	}); err != nil {
		return err
	}

	if len(p.EventParts) != 0 {
		// TODO: возможно сделать в горутинах, если медленно работать будет
		for _, ep := range p.EventParts {
			if err := e.q.WithTx(tx).EventPartUpdate(
				ctx, db.EventPartUpdateParameters{
					Id: ep.Id,
					EventPartCommonParams: db.EventPartCommonParams{
						Name:        ep.Name,
						Description: ep.Description,
						StartTime:   ep.StartTime,
						EndTime:     ep.EndTime,
						Address:     ep.Address,
						Place:       ep.Place,
						Age:         ep.Age,
					},
				},
			); err != nil {
				return err
			}
		}
	}

	return nil
}
