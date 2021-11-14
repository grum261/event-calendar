package db

const (
	eventInsert  = `INSERT INTO events (name, start_date, end_date, url) VALUES ($1, $2, $3, $4)`
	eventUpdate  = `UPDATE events SET name = $2, start_date = $3, end_date = $4, url = $5 WHERE id = $1`
	eventDelete  = `DELETE FROM events WHERE id = $1`
	eventsSelect = `
	SELECT e.id, e.name, e.start_date, e.end_date, e.url
	FROM events e
	INNER JOIN events_parts ep ON e.id = ep.event_id
	WHERE extract(month FROM e.start_date) = $2 and extract(year FROM e.start_date = $3)`
	eventsSelectByCity = `
	SELECT e.id, e.name, e.start_date, e.end_date, e.url
	FROM events e`
)
