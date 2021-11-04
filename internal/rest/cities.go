package rest

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/grum261/event-calendar/internal/models"
)

type CityService interface {
	Create(ctx context.Context, name string, timezone int) (int, error)
	Update(ctx context.Context, id, timezone int, name string) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]models.City, error)
}

type CityHandler struct {
	svc CityService
}

func newCityHandler(svc CityService) *CityHandler {
	return &CityHandler{
		svc: svc,
	}
}

type CityResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Timezone int    `json:"timezone"`
}

type CityCreateUpdateRequest struct {
	Name     string `json:"name"`
	Timezone int    `json:"timezone"`
}

func (ch *CityHandler) registerRoutes(r fiber.Router) {
	r.Get("/cities", ch.getAll)
	r.Post("/cities", ch.create)
	r.Put("/cities/:id", ch.update)
	r.Delete("/cities/:id", ch.delete)
}

func (ch *CityHandler) create(c *fiber.Ctx) error {
	req := CityCreateUpdateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableError(c, err)
	}

	id, err := ch.svc.Create(c.Context(), req.Name, req.Timezone)
	if err != nil {
		return respondInternalError(c, err)
	}

	return respondCreated(c, CityResponse{
		Id:       id,
		Name:     req.Name,
		Timezone: req.Timezone,
	})
}

func (ch *CityHandler) update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return respondUnprocessableError(c, err)
	}

	req := CityCreateUpdateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableError(c, err)
	}

	if err := ch.svc.Update(c.Context(), id, req.Timezone, req.Name); err != nil {
		return respondInternalError(c, err)
	}

	return respondOK(c, CityResponse{
		Id:       id,
		Name:     req.Name,
		Timezone: req.Timezone,
	})
}

func (ch *CityHandler) delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return respondUnprocessableError(c, err)
	}

	if err := ch.svc.Delete(c.Context(), id); err != nil {
		return respondInternalError(c, err)
	}

	return respondOK(c, id)
}

func (ch *CityHandler) getAll(c *fiber.Ctx) error {
	cities, err := ch.svc.GetAll(c.Context())
	if err != nil {
		return respondInternalError(c, err)
	}

	var res []CityResponse

	for _, c := range cities {
		res = append(res, CityResponse{
			Id:       c.Id,
			Name:     c.Name,
			Timezone: c.Timezone,
		})
	}

	return respondOK(c, res)
}
