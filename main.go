package main

import (
	"log"

	"github.com/abdule-yo/eCommerce-api/database"
	"github.com/abdule-yo/eCommerce-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to an API built on Go")
}

func SetupRoutes(app *fiber.App) {

	app.Get("/api", welcome)
	//USER ENDPOINTS
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)

}

func main() {

	database.ConnectDb()

	app := fiber.New()

	SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
