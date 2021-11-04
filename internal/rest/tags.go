package rest

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/grum261/event-calendar/internal/models"
)

type TagService interface {
	Create(ctx context.Context, names []string) ([]int, error)
	Update(ctx context.Context, id int, name string) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]models.Tag, error)
}

type TagHandler struct {
	svc TagService
}

func newTagHandler(svc TagService) *TagHandler {
	return &TagHandler{
		svc: svc,
	}
}

type TagReponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TagsCreateRequest struct {
	Tags []string `json:"tags"`
}

type TagNameUpdateRequest struct {
	Name string `json:"name"`
}

func (t *TagHandler) registerRoutes(r fiber.Router) {
	r.Get("/tags", t.getAll)
	r.Post("/tags", t.create)
	r.Put("/tags/:id", t.update)
	r.Delete("/tags/:id", t.delete)
}

func (t *TagHandler) create(c *fiber.Ctx) error {
	req := TagsCreateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableError(c, err)
	}

	ids, err := t.svc.Create(c.Context(), req.Tags)
	if err != nil {
		return respondInternalError(c, err)
	}

	res := []TagReponse{}

	for i, id := range ids {
		res = append(res, TagReponse{
			Id:   id,
			Name: req.Tags[i],
		})
	}

	return respondCreated(c, res)
}

func (t *TagHandler) update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return respondUnprocessableError(c, err)
	}

	req := TagNameUpdateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableError(c, err)
	}

	if err := t.svc.Update(c.Context(), id, req.Name); err != nil {
		return respondInternalError(c, err)
	}

	return respondOK(c, TagReponse{
		Id:   id,
		Name: req.Name,
	})
}

func (t *TagHandler) delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return respondUnprocessableError(c, err)
	}

	if err := t.svc.Delete(c.Context(), id); err != nil {
		return respondInternalError(c, err)
	}

	return respondOK(c, id)
}

func (t *TagHandler) getAll(c *fiber.Ctx) error {
	tags, err := t.svc.GetAll(c.Context())
	if err != nil {
		return respondInternalError(c, err)
	}

	res := []TagReponse{}

	for _, t := range tags {
		res = append(res, TagReponse{
			Id:   t.Id,
			Name: t.Name,
		})
	}

	return respondOK(c, res)
}
