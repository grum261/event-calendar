package db

import (
	"context"
	"time"
)

type EventPositionInsertParams struct {
	EventId   int
	EventDate time.Time
	Position  int
}

func (q *Queries) EventPositionInsert(ctx context.Context, eventId, position int, eventDate time.Time) error {
	if _, err := q.db.Exec(ctx, eventPositionInsert, eventId, eventDate, position); err != nil {
		return err
	}

	return nil
}

func (q *Queries) EventPositionUpdate(ctx context.Context, eventId, position int, eventDate time.Time) error {
	if _, err := q.db.Exec(ctx, eventPositionUpdate, eventId, eventDate, position); err != nil {
		return err
	}

	return nil
}

func (q *Queries) EventPositionDelete(ctx context.Context, eventId int, eventDate time.Time) error {
	if _, err := q.db.Exec(ctx, eventPositionUpdate, eventId, eventDate); err != nil {
		return err
	}

	return nil
}
