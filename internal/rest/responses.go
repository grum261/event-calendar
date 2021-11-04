package rest

import "github.com/gofiber/fiber/v2"

type jsonReponse struct {
	Error  string      `json:"error"`
	Result interface{} `json:"result"`
}

func respond(c *fiber.Ctx, statusCode int, result interface{}, err error) error {
	if err != nil {
		return c.Status(statusCode).JSON(jsonReponse{
			Error:  err.Error(),
			Result: nil,
		})
	}

	return c.Status(statusCode).JSON(jsonReponse{
		Error:  "",
		Result: result,
	})
}

func respondOK(c *fiber.Ctx, result interface{}) error {
	return respond(c, fiber.StatusOK, result, nil)
}

func respondCreated(c *fiber.Ctx, result interface{}) error {
	return respond(c, fiber.StatusCreated, result, nil)
}

func respondUnprocessableError(c *fiber.Ctx, err error) error {
	return respond(c, fiber.StatusUnprocessableEntity, nil, err)
}

func respondInternalError(c *fiber.Ctx, err error) error {
	return respond(c, fiber.StatusInternalServerError, nil, err)
}
