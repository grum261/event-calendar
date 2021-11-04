package service

import (
	"context"

	"github.com/grum261/event-calendar/internal/models"
)

type EventPartRepository interface {
	Update(ctx context.Context, p models.EventPartUpdateParameters) error
	Delete(ctx context.Context, id int) error
}

type EventPart struct {
	repo EventPartRepository
}

func newEventPart(repo EventPartRepository) *EventPart {
	return &EventPart{
		repo: repo,
	}
}

func (ep *EventPart) Update(ctx context.Context, p models.EventPartUpdateParameters) error {
	return ep.repo.Update(ctx, p)
}

func (ep *EventPart) Delete(ctx context.Context, id int) error {
	return ep.repo.Delete(ctx, id)
}
