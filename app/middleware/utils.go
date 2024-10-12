package middleware

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3/v2"
	"github.com/gofiber/utils"
	"golang.org/x/crypto/bcrypt"
)

var Session *session.Store

func Render(c *fiber.Ctx, view string) error {
	return c.Render(view, fiber.Map{
		"Title":  os.Getenv("TITLE"),
		"UserID": GetSessionCookie(c),
	}, "layouts/main")
}

func Redirect(c *fiber.Ctx, view, route string) error {
	if c.Get("HX-Request") == "true" {
		c.Set("HX-Redirect", route)
		return c.SendStatus(fiber.StatusOK)
	}

	return Render(c, view)
}

func ConnectSessionsDB() {
	storage := sqlite3.New(sqlite3.Config{
		Table:    "session_storage",
		Database: "./config/database/sessions.db",
	})
	Session = session.New(session.Config{
		Storage:        storage,
		Expiration:     30 * time.Minute,
		KeyLookup:      "cookie:session_id",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: "Strict",
		KeyGenerator:   utils.UUID,
	})
}

func HashPassword(password string) string {
	// Returns a hashed and salted password
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashBytes)
}

func ValidatePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func AuthMiddleware(c *fiber.Ctx) error {
	/*
		Prevents the logged user to access the login and register pages
		and prevents the no logged user to access protected pages
	*/
	sess, err := Session.Get(c)
	if err != nil {
		log.Println(err)
	}

	userID := sess.Get("user_id")
	switch c.Path() {

	// If the user is logged redirect him to the index page instead of /login or /register
	case "/login", "/register":

		// Redirects the user to the index page if they're already logged in
		if userID != nil {
			return Redirect(c, "index", "/")
		} else {
			return c.Next()
		}
	}

	// if the user is not
	if userID == nil {
		return Redirect(c, "index", "/")
	}

	sess.Save()
	c.Locals("user_id", userID.(uint))

	return c.Next()
}

func SetSessionCookie(c *fiber.Ctx, id uint) {
	session, err := Session.Get(c)
	if err != nil {
		log.Println("Failed to get session.")
	}

	// Saves the user_id as a cookie in the user's browser
	session.Set("user_id", id)
	session.Save()
}

func GetSessionCookie(c *fiber.Ctx) interface{} {
	session, err := Session.Get(c)
	if err != nil {
		log.Println("Failed to get session.")
	}

	return session.Get("user_id")
}

func ClearSessionCookie(c *fiber.Ctx) {
	c.ClearCookie("user_id")
}
