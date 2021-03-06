package db

const (
	eventPartInsert = `
	INSERT INTO events_parts (name, address, description, start_time, end_time, event_id, place, age) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	eventPartUpdate = `
	UPDATE events_parts SET name = $2, address = $3, description = $4, start_time = $5, 
	end_time = $6, place = $7, age = $8
	WHERE id = $1`
	eventPartDelete                = `DELETE FROM events_parts WHERE id = $1`
	eventPartsSelectByEventDayCity = `
	SELECT id, name, address, description, start_time, end_time, place, age 
	FROM events_parts
	WHERE event_id = $1 and extract(day FROM start_time) = $3`
	eventPartsSelectByEventCity = `
	SELECT id, name, address, description, start_time, end_time, place, age
	FROM events_parts
	WHERE event_id = $1`
)
