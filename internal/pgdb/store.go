package pgdb

import "github.com/jackc/pgx/v4/pgxpool"

type Store struct {
	*Tag
	*City
	*Event
	*EventPart
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		Tag:       newTag(db),
		City:      newCity(db),
		Event:     newEvent(db),
		EventPart: newEventPart(db),
	}
}
