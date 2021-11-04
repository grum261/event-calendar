package models

import "time"

type EventPart struct {
	Id          int
	Name        string
	Address     string
	Description string
	CityId      int
	StartTime   time.Time
	EndTime     time.Time
	Place       string
}

type EventPartCreateParameters struct {
	EventId     int
	Name        string
	Address     string
	Description string
	CityId      int
	StartTime   time.Time
	EndTime     time.Time
	Place       string
	Age         int
}

type EventPartUpdateParameters struct {
	Id          int
	Name        string
	Address     string
	Description string
	CityId      int
	StartTime   time.Time
	EndTime     time.Time
	Place       string
	Age         int
}
