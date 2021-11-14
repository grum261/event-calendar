package rest

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/grum261/event-calendar/internal/models"
)

type EventService interface {
	Create(ctx context.Context, p models.EventInsertParameters) (int, []int, error)
	Update(ctx context.Context, eventId int, p models.EventUpdateParameters) error
	Delete(ctx context.Context, id int) error
}

type EventHandler struct {
	svc EventService
}

func newEventHandler(svc EventService) *EventHandler {
	return &EventHandler{
		svc: svc,
	}
}

type EventCreateRequest struct {
	Name       string                   `json:"name"`
	StartDate  time.Time                `json:"startDate"`
	EndDate    time.Time                `json:"endDate"`
	URL        string                   `json:"url"`
	EventParts []EventPartCreateRequest `json:"eventParts"`
}

type EventResponse struct {
	Id         int                  `json:"id"`
	Name       string               `json:"name"`
	StartDate  time.Time            `json:"startDate"`
	EndDate    time.Time            `json:"endDate"`
	URL        string               `json:"url,omitempty"`
	EventParts *[]EventPartResponse `json:"eventParts,omitempty"`
}

type EventUpdateRequest struct {
	Name       string                   `json:"name"`
	StartDate  time.Time                `json:"startDate"`
	EndDate    time.Time                `json:"endDate"`
	URL        string                   `json:"url"`
	EventParts []EventPartUpdateRequest `json:"eventParts"`
}

func (e *EventHandler) registerRoutes(r fiber.Router) {
	r.Post("/events", e.create)
	r.Put("/events/:id", e.update)
	r.Delete("/events/:id", e.delete)
}

func (e *EventHandler) create(c *fiber.Ctx) error {
	req := EventCreateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableError(c, err)
	}

	p := models.EventInsertParameters{
		Name:      req.Name,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		URL:       req.URL,
	}

	// TODO: прописать логику для ивента, проходящего 1 день и у которого нет частей
	for _, ep := range req.EventParts {
		p.EventParts = append(p.EventParts, models.EventPartCreateParameters{
			EventId: ep.EventId,
			EventPartInsertUpdateParams: models.EventPartInsertUpdateParams{
				Name:        ep.Name,
				Address:     ep.Address,
				Description: ep.Description,
				StartTime:   ep.StartTime,
				EndTime:     ep.EndTime,
				Place:       ep.Place,
				Age:         ep.Age,
			},
		})
	}

	eventId, partsIds, err := e.svc.Create(c.Context(), p)
	if err != nil {
		return respondInternalError(c, err)
	}

	res := EventResponse{
		Id:         eventId,
		Name:       req.Name,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		URL:        req.URL,
		EventParts: &[]EventPartResponse{},
	}

	for i, partId := range partsIds {
		*res.EventParts = append(*res.EventParts, EventPartResponse{
			Id:          partId,
			Name:        req.EventParts[i].Name,
			Age:         req.EventParts[i].Age,
			Description: req.EventParts[i].Description,
			Address:     req.EventParts[i].Address,
			Place:       req.EventParts[i].Place,
			StartTime:   req.EventParts[i].StartTime,
			EndTime:     req.EventParts[i].EndTime,
		})
	}

	return respondCreated(c, res)
}

func (e *EventHandler) update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return respondUnprocessableError(c, err)
	}

	req := EventUpdateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableError(c, err)
	}

	p := models.EventUpdateParameters{
		Name:      req.Name,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		URL:       req.URL,
	}

	res := EventResponse{
		Id:        id,
		Name:      req.Name,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		URL:       req.URL,
	}

	for _, ep := range req.EventParts {
		p.EventParts = append(p.EventParts, models.EventPartUpdateParameters{
			Id: ep.Id,
			EventPartInsertUpdateParams: models.EventPartInsertUpdateParams{
				Name:        ep.Name,
				Address:     ep.Address,
				Description: ep.Description,
				StartTime:   ep.StartTime,
				EndTime:     ep.EndTime,
				Place:       ep.Place,
				Age:         ep.Age,
			},
		})

		*res.EventParts = append(*res.EventParts, EventPartResponse{
			Id:          ep.Id,
			Name:        ep.Name,
			Age:         ep.Age,
			Description: ep.Description,
			Address:     ep.Address,
			Place:       ep.Place,
			StartTime:   ep.StartTime,
			EndTime:     ep.EndTime,
		})
	}

	if err := e.svc.Update(c.Context(), id, p); err != nil {
		return respondInternalError(c, err)
	}

	return respondOK(c, res)
}

func (e *EventHandler) delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return respondUnprocessableError(c, err)
	}

	if err := e.svc.Delete(c.Context(), id); err != nil {
		return respondInternalError(c, err)
	}

	return respondOK(c, id)
}
