package routes

import (
	"todo-app/app/middleware"
	"todo-app/config/database"

	"github.com/gofiber/fiber/v2"
)

// func convertUser(user database.User) map[string]interface{} {
// 	return map[string]interface{}{
// 		"ID":       user.ID,
// 		"Username": user.Username,
// 		"Email":    user.Email,
// 	}
// }

func Index(c *fiber.Ctx) error {
	return middleware.Redirect(c, "index", "/")
}

func LoginPost(c *fiber.Ctx) error {
	email := c.FormValue("email")
	if !database.UserExists(email, "email") {
		return c.SendString("Wrong email")
	}

	user := database.SearchUserByString(email, "email")

	// Check if the password matches the password hash
	if middleware.ValidatePassword(user.Password, c.FormValue("password")) {
		middleware.SetSessionCookie(c, user.ID)
		return middleware.Redirect(c, "index", "/")
	}

	return c.SendString("Wrong password")
}

func LoginGet(c *fiber.Ctx) error {
	return middleware.Redirect(c, "login", "/login")
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
		return middleware.Redirect(c, "login", "/login")
	}

	return c.SendString("Email already used")
}

func RegisterGet(c *fiber.Ctx) error {
	return middleware.Redirect(c, "register", "/register")
}

func LogoutGet(c *fiber.Ctx) error {
	middleware.ClearSessionCookie(c)
	return middleware.Redirect(c, "index", "/")
}

func TodoGet(c *fiber.Ctx) error {
	return middleware.Redirect(c, "todo", "/todo")
}

func AddListPost(c *fiber.Ctx) error {
	list := database.List{
		Name:   c.FormValue("list-name"), // Mudando para "list-name" para corresponder ao frontend
		UserID: c.Locals("user_id").(uint),
	}
	database.Database.Create(&list)
	return c.SendString("List created")
}

func RemoveListPost(c *fiber.Ctx) error {
	listID := c.FormValue("list_id")
	database.Database.Where("id = ?", listID).Delete(&database.List{})
	return c.SendString("List removed")
}
