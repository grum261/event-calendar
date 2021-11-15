package rest

import "github.com/gofiber/fiber/v2"

type Handlers struct {
	*TagHandler
	*EventHandler
	*EventPartHandler
}

func NewHandlers(t TagService, e EventService, ep EventPartService) *Handlers {
	return &Handlers{
		TagHandler:       newTagHandler(t),
		EventHandler:     newEventHandler(e),
		EventPartHandler: newEventPartHandler(ep),
	}
}

func (h *Handlers) RegisterRoutes(r fiber.Router) {
	h.TagHandler.registerRoutes(r)

	h.EventHandler.registerRoutes(r)
}
