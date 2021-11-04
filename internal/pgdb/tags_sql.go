package pgdb

import "context"

type tagReturn struct {
	Id   int
	Name string
}

func (q *Queries) tagsInsert(ctx context.Context, names []string) ([]int, error) {
	rows, err := q.db.Query(ctx, tagsInsert, names)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int

	for rows.Next() {
		var id int

		if err := rows.Scan(&id); err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func (q *Queries) tagNameUpdate(ctx context.Context, id int, name string) error {
	if _, err := q.db.Exec(ctx, tagNameUpdate, id, name); err != nil {
		return err
	}

	return nil
}

func (q *Queries) tagDelete(ctx context.Context, id int) error {
	if _, err := q.db.Exec(ctx, tagDelete, id); err != nil {
		return err
	}

	return nil
}

func (q *Queries) tagsSelectAll(ctx context.Context) ([]tagReturn, error) {
	rows, err := q.db.Query(ctx, tagsSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []tagReturn

	for rows.Next() {
		var tag tagReturn

		if err := rows.Scan(&tag.Id, &tag.Name); err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}
