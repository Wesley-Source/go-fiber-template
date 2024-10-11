package main

import (
	"go-fiber-template/app/middleware"
	"go-fiber-template/app/routes"
	"go-fiber-template/config/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {

	// Initializing and connecting to the databases
	database.ConnectDatabase()
	middleware.ConnectSessionsDB()

	// Loading the global variables
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalln(err)
	}

	engine := html.New("./app/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Here is where you want to put your icons, scripts, thumbnails etc
	app.Static("/", "./app/public")

	// Routes:
	app.Get("/", routes.Index)

	app.Get("/login", middleware.AuthMiddleware, routes.LoginGet)
	app.Post("/login", middleware.AuthMiddleware, routes.LoginPost)

	app.Post("/register", middleware.AuthMiddleware, routes.RegisterPost)
	app.Get("/register", middleware.AuthMiddleware, routes.RegisterGet)

	log.Fatalln(app.Listen(os.Getenv("PORT")))
}
