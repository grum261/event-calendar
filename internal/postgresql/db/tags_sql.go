package db

import "context"

type TagReturn struct {
	Id   int
	Name string
}

func (q *Queries) TagsInsert(ctx context.Context, names []string) ([]int, error) {
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

func (q *Queries) TagNameUpdate(ctx context.Context, id int, name string) error {
	if _, err := q.db.Exec(ctx, tagNameUpdate, id, name); err != nil {
		return err
	}

	return nil
}

func (q *Queries) TagDelete(ctx context.Context, id int) error {
	if _, err := q.db.Exec(ctx, tagDelete, id); err != nil {
		return err
	}

	return nil
}

func (q *Queries) TagsSelectAll(ctx context.Context) ([]TagReturn, error) {
	rows, err := q.db.Query(ctx, tagsSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []TagReturn

	for rows.Next() {
		var tag TagReturn

		if err := rows.Scan(&tag.Id, &tag.Name); err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}
