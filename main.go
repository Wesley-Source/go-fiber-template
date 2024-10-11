package main

import (
	"go-fiber-template/app/routes"
	"go-fiber-template/config/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalln(err)
	}

	engine := html.New("./app/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./app/plubic")

	database.ConnectDatabase()

	app.Get("/", routes.Index)

	log.Fatalln(app.Listen(os.Getenv("PORT")))
}
