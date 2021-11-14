package db

const (
	cityInsert      = `INSERT INTO cities (name, timezone) VALUES ($1, $2) RETURNING id`
	cityUpdate      = `UPDATE cities SET name = $2, timezone = $3 WHERE id = $1`
	cityDelete      = `DELETE FROM cities WHERE id = $1`
	citiesSelectAll = `SELECT id, name, timezone FROM cities`
)
