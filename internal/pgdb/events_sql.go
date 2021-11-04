package pgdb

import (
	"context"
	"time"
)

type eventInsertParams struct {
	Name      string
	StartDate time.Time
	EndDate   time.Time
	URL       string
}

type eventReturn struct {
	Id        int
	Name      string
	StartDate time.Time
	EndDate   time.Time
	URL       string
}

func (q *Queries) eventInsert(ctx context.Context, p eventInsertParams) (int, error) {
	var id int

	// name, start_date, end_date, url
	if err := q.db.QueryRow(ctx, eventInsert, p.Name, p.StartDate, p.EndDate, p.URL).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (q *Queries) eventUpdate(ctx context.Context, id int, p eventInsertParams) error {
	if _, err := q.db.Exec(ctx, eventUpdate, p.Name, p.StartDate, p.EndDate, p.URL); err != nil {
		return err
	}

	return nil
}

func (q *Queries) eventsSelectByCityYearMonth(ctx context.Context, cityId, year, month int) ([]eventReturn, error) {
	return nil, nil
}

func (q *Queries) eventDelete(ctx context.Context, id int) error {
	if _, err := q.db.Exec(ctx, eventDelete, id); err != nil {
		return err
	}

	return nil
}
