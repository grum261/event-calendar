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

type EventUpdateParams struct {
	Id int
	EventInsertParams
}

type EventReturn struct {
	Id        int
	EventDate time.Time
	Name      string
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

func (q *Queries) EventUpdate(ctx context.Context, id int, p EventUpdateParams) error {
	if _, err := q.db.Exec(ctx, eventUpdate, p.Id, p.Name, p.StartDate, p.EndDate, p.URL); err != nil {
		return err
	}

	return nil
}

func (q *Queries) EventsSelectByYearMonth(ctx context.Context, year, month int) ([]EventReturn, error) {
	rows, err := q.db.Query(ctx, eventsSelectByYearMonth, year, month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var er []EventReturn

	for rows.Next() {
		var e EventReturn

		if err := rows.Scan(&e.Id, &e.Name, &e.URL, &e.EventDate); err != nil {
			return nil, err
		}

		er = append(er, e)
	}

	return er, nil
}

func (q *Queries) EventDelete(ctx context.Context, id int) error {
	if _, err := q.db.Exec(ctx, eventDelete, id); err != nil {
		return err
	}

	return nil
}
