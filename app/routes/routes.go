package routes

import (
<<<<<<< HEAD
	"strconv"
	"todo-app/app/middleware"
	"todo-app/config/database"
=======
	"go-fiber-template/app/middleware"
	"go-fiber-template/config/database"
>>>>>>> parent of 696f3ff (todo-app)

	"github.com/gofiber/fiber/v2"
)

func convertUser(user database.User) map[string]interface{} {
	return map[string]interface{}{
		"ID":       user.ID,
		"Username": user.Username,
		"Email":    user.Email,
	}
}

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
		return c.SendString("Registred")

	}

	return c.SendString("Email already used")
}

func RegisterGet(c *fiber.Ctx) error {
	return middleware.Redirect(c, "register", "/register")
}

func LogoutPost(c *fiber.Ctx) error {
	middleware.ClearSessionCookie(c)
<<<<<<< HEAD
	return middleware.Redirect(c, "index", "/")
}

func TodoGet(c *fiber.Ctx) error {
	return middleware.Redirect(c, "todo", "/todo")
}

func ListAddPost(c *fiber.Ctx) error {
	list := database.List{
		Title:  c.FormValue("list_title"),
		UserID: c.Locals("user_id").(uint),
	}

	database.Database.Create(&list)
	// Quando for uma requisição HTMX, use RenderPartial
	return middleware.Render(c, "partials/menus-list", true)
}

func TaskAddPost(c *fiber.Ctx) error {
	listID, err := strconv.ParseUint(c.FormValue("list_id"), 10, 32)
	if err != nil {
		return c.Status(400).SendString("Invalid list ID")
	}

	task := database.Task{
		Title:       c.FormValue("task_title"),
		Description: c.FormValue("task_description"),
		DueDate:     c.FormValue("task_date"),
		Completed:   false,
		ListID:      uint(listID),
	}

	// todo, task e lista direito, tem coisas erradas

	database.Database.Create(&task)
	return c.SendString("Task created")
=======
	return c.SendStatus(fiber.StatusOK)
>>>>>>> parent of 696f3ff (todo-app)
}
