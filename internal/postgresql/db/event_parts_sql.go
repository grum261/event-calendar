package db

import (
	"context"
	"time"
)

type EventPartInsertParams struct {
	EventId int
	EventPartCommonParams
}

type EventPartUpdateParameters struct {
	Id int
	EventPartCommonParams
}

type EventPartCommonParams struct {
	Name        string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	Address     string
	Place       string
	Age         int
}

type EventPartReturn struct {
	Id int
	EventPartCommonParams
}

func (q *Queries) EventPartsInsert(ctx context.Context, p EventPartInsertParams) (int, error) {
	var id int

	// name, address, description, start_time, end_time, event_id, place
	if err := q.db.QueryRow(
		ctx, eventPartInsert, p.Name, p.Address, p.Description,
		p.StartTime, p.EndTime, p.EventId, p.Place, p.Age,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (q *Queries) EventPartUpdate(ctx context.Context, p EventPartUpdateParameters) error {
	if _, err := q.db.Exec(
		ctx, eventPartUpdate, p.Id, p.Name, p.Address, p.Description,
		p.StartTime, p.EndTime, p.Place, p.Age,
	); err != nil {
		return err
	}

	return nil
}

func (q *Queries) EventPartSelectByEventDay(ctx context.Context, eventId, day int) ([]EventPartReturn, error) {
	rows, err := q.db.Query(ctx, eventPartsSelectByEventDayCity, eventId, day)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventParts []EventPartReturn

	for rows.Next() {
		var ep EventPartReturn

		// id, name, address, description, start_time, end_time, place
		if err := rows.Scan(
			&ep.Id, &ep.Name, &ep.Address, &ep.Description,
			&ep.StartTime, &ep.EndTime, &ep.Place, &ep.Age,
		); err != nil {
			return nil, err
		}

		eventParts = append(eventParts, ep)
	}

	return eventParts, nil
}

func (q *Queries) EventPartsDelete(ctx context.Context, id int) error {
	if _, err := q.db.Exec(ctx, eventPartDelete, id); err != nil {
		return err
	}

	return nil
}

func (q *Queries) EventPartsSelectByEvent(ctx context.Context, eventId int) ([]EventPartReturn, error) {
	rows, err := q.db.Query(ctx, eventPartsSelectByEventCity, eventId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventParts []EventPartReturn

	for rows.Next() {
		var ep EventPartReturn

		if err := rows.Scan(
			&ep.Id, &ep.Name, &ep.Address, &ep.Description,
			&ep.StartTime, &ep.EndTime, &ep.Place, &ep.Age,
		); err != nil {
			return nil, err
		}

		eventParts = append(eventParts, ep)
	}

	return eventParts, nil
}
