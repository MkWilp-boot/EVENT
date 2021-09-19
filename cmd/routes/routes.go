package routes

import "github.com/gofiber/fiber/v2"

func EventGetById(c *fiber.Ctx) error {
	sample := struct {
		Route string
	}{
		Route: "EventGetById_" + string(c.Params("id")),
	}
	return c.JSON(sample)
}
