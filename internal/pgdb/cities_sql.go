package pgdb

import "context"

type cityReturn struct {
	Id       int
	Name     string
	Timezone int
}

func (q *Queries) cityInsert(ctx context.Context, name string, timezone int) (int, error) {
	var id int

	if err := q.db.QueryRow(ctx, cityInsert, name, timezone).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (q *Queries) cityUpdate(ctx context.Context, id, timezone int, name string) error {
	if _, err := q.db.Exec(ctx, cityUpdate, id, name, timezone); err != nil {
		return err
	}

	return nil
}

func (q *Queries) cityDelete(ctx context.Context, id int) error {
	if _, err := q.db.Exec(ctx, cityDelete, id); err != nil {
		return err
	}

	return nil
}

func (q *Queries) citiesSelectAll(ctx context.Context) ([]cityReturn, error) {
	rows, err := q.db.Query(ctx, citiesSelectAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []cityReturn

	for rows.Next() {
		var city cityReturn

		if err := rows.Scan(&city.Id, &city.Name, &city.Timezone); err != nil {
			return nil, err
		}

		cities = append(cities, city)
	}

	return cities, nil
}
