package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Render("layouts/main", fiber.Map{
		"Title": os.Getenv("TITLE"),
	}, "hello")

}
