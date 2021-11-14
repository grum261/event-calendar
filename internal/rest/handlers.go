package rest

import "github.com/gofiber/fiber/v2"

type Handlers struct {
	*TagHandler
	// *CityHandler
	*EventHandler
}

func NewHandlers(t TagService, e EventService) *Handlers {
	return &Handlers{
		TagHandler: newTagHandler(t),
		// CityHandler:  newCityHandler(c),
		EventHandler: newEventHandler(e),
	}
}

func (h *Handlers) RegisterRoutes(r fiber.Router) {
	h.TagHandler.registerRoutes(r)

	// h.CityHandler.registerRoutes(r)

	h.EventHandler.registerRoutes(r)
}
