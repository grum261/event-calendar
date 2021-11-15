package models

import "time"

type Event struct {
	Id   int
	Name string
	Date time.Time
	URL  string
}

type EventInsertParameters struct {
	EventCommons
	EventParts     []EventPartCreateParameters
	EventPositions []EventPosition
}

type EventUpdateParameters struct {
	Id             int
	EventParts     []EventPartUpdateParameters
	EventPositions []EventPosition
	EventCommons
}

type EventCommons struct {
	Name      string
	StartDate time.Time
	EndDate   time.Time
	URL       string
}
