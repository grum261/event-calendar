package pgdb

import (
	"context"
	"time"
)

type eventPartInsertParams struct {
	EventId     int
	Name        string
	CityId      int
	Description string
	StartTime   time.Time
	EndTime     time.Time
	Address     string
	Place       string
	Age         int
}

type eventPartUpdateParameters struct {
	Id          int
	Name        string
	CityId      int
	Description string
	StartTime   time.Time
	EndTime     time.Time
	Address     string
	Place       string
	Age         int
}

type eventPartReturn struct {
	Id          int
	Name        string
	CityId      int
	Description string
	StartTime   time.Time
	EndTime     time.Time
	Address     string
	Place       string
	Age         int
}

func (q *Queries) eventPartsInsert(ctx context.Context, p eventPartInsertParams) (int, error) {
	var id int

	// name, address, description, start_time, end_time, event_id, city_id, place
	if err := q.db.QueryRow(
		ctx, eventPartInsert, p.Name, p.Address, p.Description,
		p.StartTime, p.EndTime, p.EventId, p.CityId, p.Place, p.Age,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (q *Queries) eventPartUpdate(ctx context.Context, p eventPartUpdateParameters) error {
	if _, err := q.db.Exec(
		ctx, eventPartUpdate, p.Id, p.Name, p.Address, p.Description,
		p.StartTime, p.EndTime, p.CityId, p.Place, p.Age,
	); err != nil {
		return err
	}

	return nil
}

func (q *Queries) selectEventPartByEventDayCity(ctx context.Context, eventId, cityId, day int) ([]eventPartReturn, error) {
	rows, err := q.db.Query(ctx, eventPartsSelectByEventDayCity, eventId, cityId, day)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventParts []eventPartReturn

	for rows.Next() {
		var ep eventPartReturn

		// id, name, address, description, start_time, end_time, city_id, place
		if err := rows.Scan(
			&ep.Id, &ep.Name, &ep.Address, &ep.Description,
			&ep.StartTime, &ep.EndTime, &ep.CityId, &ep.Place, &ep.Age,
		); err != nil {
			return nil, err
		}

		eventParts = append(eventParts, ep)
	}

	return eventParts, nil
}

func (q *Queries) deleteEventParts(ctx context.Context, id int) error {
	if _, err := q.db.Exec(ctx, eventPartDelete, id); err != nil {
		return err
	}

	return nil
}

func (q *Queries) eventPartsSelectByEventCity(ctx context.Context, eventId, cityId int) ([]eventPartReturn, error) {
	rows, err := q.db.Query(ctx, eventPartsSelectByEventCity, eventId, cityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventParts []eventPartReturn

	for rows.Next() {
		var ep eventPartReturn

		if err := rows.Scan(
			&ep.Id, &ep.Name, &ep.Address, &ep.Description,
			&ep.StartTime, &ep.EndTime, &ep.CityId, &ep.Place, &ep.Age,
		); err != nil {
			return nil, err
		}

		eventParts = append(eventParts, ep)
	}

	return eventParts, nil
}
