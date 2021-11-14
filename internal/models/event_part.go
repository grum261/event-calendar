package models

import "time"

type EventPart struct {
	Id          int
	Name        string
	Address     string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	Place       string
}

type EventPartCreateParameters struct {
	EventId int
	EventPartInsertUpdateParams
}

type EventPartUpdateParameters struct {
	Id int
	EventPartInsertUpdateParams
}

type EventPartInsertUpdateParams struct {
	Name        string
	Address     string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	Place       string
	Age         int
}
