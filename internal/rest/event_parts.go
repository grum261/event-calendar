package rest

import (
	"context"
	"time"

	"github.com/grum261/event-calendar/internal/models"
)

type EventPartService interface {
	Update(ctx context.Context, id int, p models.EventPartCreateParameters) error
	Delete(ctx context.Context, id int) error
}

type EventPartHandler struct {
	svc EventPartService
}

func newEventPartHandler(svc EventPartService) *EventPartHandler {
	return &EventPartHandler{
		svc: svc,
	}
}

type EventPartResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Age         int       `json:"age"`
	Description string    `json:"description,omitempty"`
	Address     string    `json:"address,omitempty"`
	Place       string    `json:"place"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
}

type EventPartCreateRequest struct {
	Name        string    `json:"name"`
	Age         int       `json:"age"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	Place       string    `json:"place"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	EventId     int       `json:"eventID"`
}

type EventPartUpdateRequest struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Age         int       `json:"age"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	Place       string    `json:"place"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
}
