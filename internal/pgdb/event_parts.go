package pgdb

import (
	"context"

	"github.com/grum261/event-calendar/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type EventPart struct {
	q *Queries
}

func newEventPart(db *pgxpool.Pool) *EventPart {
	return &EventPart{
		q: newQueries(db),
	}
}

// func (ep *EventPart) Create(ctx context.Context, args []models.EventPartCreateParameters) ([]int, error) {
// 	var ids []int

// 	for _, p := range args {
// 		id, err := ep.q.eventPartsInsert(ctx, eventPartInsertParams{
// 			EventId:     p.EventId,
// 			Name:        p.Name,
// 			CityId:      p.CityId,
// 			Description: p.Description,
// 			StartTime:   p.StartTime,
// 			EndTime:     p.EndTime,
// 			Address:     p.Address,
// 			Place:       p.Place,
// 			Age:         p.Age,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}

// 		ids = append(ids, id)
// 	}

// 	return ids, nil
// }

func (e *EventPart) Update(ctx context.Context, p models.EventPartUpdateParameters) error {
	return e.q.eventPartUpdate(ctx, eventPartUpdateParameters{
		Id:          p.Id,
		Name:        p.Name,
		CityId:      p.CityId,
		Description: p.Description,
		StartTime:   p.StartTime,
		EndTime:     p.EndTime,
		Address:     p.Address,
		Place:       p.Place,
		Age:         p.Age,
	})
}

func (e *EventPart) Delete(ctx context.Context, id int) error {
	return e.q.deleteEventParts(ctx, id)
}
