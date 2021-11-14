package db

const (
	tagsInsert    = `INSERT INTO tags (name) (SELECT unnest($1::varchar[])) RETURNING id`
	tagNameUpdate = `UPDATE tags SET name = $2 WHERE id = $1`
	tagsSelect    = `SELECT id, name FROM tags`
	tagDelete     = `DELETE FROM tags WHERE id = $1`
)
