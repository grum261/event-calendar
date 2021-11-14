package db

import (
	"context"
	"time"
)

type EventInsertParams struct {
	Name      string
	StartDate time.Time
	EndDate   time.Time
	URL       string
}

type EventReturn struct {
	Id        int
	Name      string
	StartDate time.Time
	EndDate   time.Time
	URL       string
}

func (q *Queries) EventInsert(ctx context.Context, p EventInsertParams) (int, error) {
	var id int

	// name, start_date, end_date, url
	if err := q.db.QueryRow(ctx, eventInsert, p.Name, p.StartDate, p.EndDate, p.URL).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (q *Queries) EventUpdate(ctx context.Context, id int, p EventInsertParams) error {
	if _, err := q.db.Exec(ctx, eventUpdate, p.Name, p.StartDate, p.EndDate, p.URL); err != nil {
		return err
	}

	return nil
}

func (q *Queries) EventsSelectByYearMonth(ctx context.Context, year, month int) ([]EventReturn, error) {
	return nil, nil
}

func (q *Queries) EventDelete(ctx context.Context, id int) error {
	if _, err := q.db.Exec(ctx, eventDelete, id); err != nil {
		return err
	}

	return nil
}
