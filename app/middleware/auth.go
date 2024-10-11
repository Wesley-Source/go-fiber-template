package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3/v2"
	"github.com/gofiber/utils"
	"golang.org/x/crypto/bcrypt"
)

var Session *session.Store

func ConnectSessionsDB() {
	storage := sqlite3.New(sqlite3.Config{
		Table: "session_storage",
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
	sess, err := Session.Get(c)
	if err != nil {
		if c.Get("HX-Request") == "true" {
			c.Set("HX-Redirect", "/login")
			return c.SendStatus(fiber.StatusOK)
		} else {

			return c.Redirect("/")
		}
	}

	user_id := sess.Get("user_id")
	if user_id == nil {
		if c.Get("HX-Request") == "true" {
			c.Set("HX-Redirect", "/login")
			return c.SendStatus(fiber.StatusOK)
		} else {

			return c.Redirect("/")
		}
	}

	if err := sess.Save(); err != nil {
		log.Printf("Erro ao salvar sessão: %v", err)
		// Decidir se deve continuar ou redirecionar para login
	}

	c.Locals("user_id", user_id.(uint))

	return c.Next()
}
