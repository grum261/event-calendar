package models

import "time"

type Event struct {
	Id        int
	Name      string
	StartDate time.Time
	EndDate   time.Time
	URL       string
}

type EventInsertParameters struct {
	Name       string
	StartDate  time.Time
	EndDate    time.Time
	URL        string
	EventParts []EventPartCreateParameters
}

type EventUpdateParameters struct {
	Name       string
	StartDate  time.Time
	EndDate    time.Time
	URL        string
	EventParts []EventPartUpdateParameters
}
