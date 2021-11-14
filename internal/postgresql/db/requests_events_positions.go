package db

const (
	eventPositionInsert = `INSERT INTO events_positions (event_id, event_date, position) VALUES ($1, $2, $3)`
	eventPositionUpdate = `UPDATE events_position SET position = $3 WHERE event_id = $1 and event_date = $3`
	eventPositionDelete = `DELETE FROM events_positions WHERE event_id = $1 and event_date = $2`
)
