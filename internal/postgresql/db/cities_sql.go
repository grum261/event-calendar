package db

import "context"

type CityReturn struct {
	Id       int
	Name     string
	Timezone int
}

func (q *Queries) CityInsert(ctx context.Context, name string, timezone int) (int, error) {
	var id int

	if err := q.db.QueryRow(ctx, cityInsert, name, timezone).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (q *Queries) CityUpdate(ctx context.Context, id, timezone int, name string) error {
	if _, err := q.db.Exec(ctx, cityUpdate, id, name, timezone); err != nil {
		return err
	}

	return nil
}

func (q *Queries) CityDelete(ctx context.Context, id int) error {
	if _, err := q.db.Exec(ctx, cityDelete, id); err != nil {
		return err
	}

	return nil
}

func (q *Queries) CitiesSelectAll(ctx context.Context) ([]CityReturn, error) {
	rows, err := q.db.Query(ctx, citiesSelectAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []CityReturn

	for rows.Next() {
		var city CityReturn

		if err := rows.Scan(&city.Id, &city.Name, &city.Timezone); err != nil {
			return nil, err
		}

		cities = append(cities, city)
	}

	return cities, nil
}
