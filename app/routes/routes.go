package routes

import (
	"go-fiber-template/app/middleware"
	"go-fiber-template/config/database"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Render("layouts/main", fiber.Map{
		"Title": os.Getenv("TITLE"),
	}, "hello")

}

func LoginPost(c *fiber.Ctx) error {

	email := c.FormValue("email")
	if !database.UserExists(email, "email") {
		return c.SendString("Wrong email")
	}

	user := database.SearchUser(email, "email")

	if middleware.ValidatePassword(user.Password, c.FormValue("password")) {
		middleware.SetSessionCookie(c, user.ID)
		c.Set("HX-Redirect", "/")
		return c.SendStatus(fiber.StatusOK)
	} else {
		return c.SendString("Wrong password")
	}

}

func LoginGet(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": os.Getenv("TITLE"),
	}, "layouts/main")
}

func RegisterPost(c *fiber.Ctx) error {
	email := c.FormValue("email")
	if !database.UserExists(email, "email") {
		user := database.User{
			Username: c.FormValue("username"),
			Email:    c.FormValue("email"),
			Password: middleware.HashPassword(c.FormValue("password")),
		}
		database.Database.Create(&user)
		return c.SendString("Registred")

	}

	return c.SendString("Email already used")
}

func RegisterGet(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": os.Getenv("TITLE"),
	}, "layouts/main")
}
