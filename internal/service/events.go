package service

import (
	"context"

	"github.com/grum261/event-calendar/internal/models"
)

type EventRepository interface {
	Create(ctx context.Context, p models.EventInsertParameters) (int, []int, error)
	Update(ctx context.Context, eventId int, p models.EventUpdateParameters) error
	Delete(ctx context.Context, id int) error
}

type Event struct {
	repo EventRepository
}

func newEvent(repo EventRepository) *Event {
	return &Event{
		repo: repo,
	}
}

func (e *Event) Create(ctx context.Context, p models.EventInsertParameters) (int, []int, error) {
	return e.repo.Create(ctx, p)
}

func (e *Event) Update(ctx context.Context, eventId int, p models.EventUpdateParameters) error {
	return e.repo.Update(ctx, eventId, p)
}

func (e *Event) Delete(ctx context.Context, id int) error {
	return e.repo.Delete(ctx, id)
}
