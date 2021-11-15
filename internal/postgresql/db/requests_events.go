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
	eventsSelectByYearMonth = `
	WITH d as (
		SELECT id, name, url, generate_series(start_date, end_date, '1 day'::interval) as date
		FROM events
		WHERE extract(year FROM start_date) = $1 and extract(month FROM start_date) = $2
	)
	SELECT id, name, coalesce(url, ''), date FROM (
		SELECT d.id, d.name, d.url, d.date, row_number() OVER (PARTITION BY d.date ORDER BY d.date, ep.position) as rn
		FROM d
		LEFT JOIN events_positions ep ON d.id = ep.event_id and d.date = ep.event_date
	) as c
	WHERE rn <= 5`
)
