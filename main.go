package main

import (
	"log"

	"github.com/abdule-yo/eCommerce-api/database"
	"github.com/gofiber/fiber/v2"
)

// Todo: sending welcoming message (the *fiber.Ctx is anything come with the fiber framework like request,body and any thing else)
func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to an API built on Go")
}

func main() {
	//Todo: connect to the DB
	database.ConnectDb()
	//Todo: Initialize the app
	app := fiber.New()

	//Todo: setting up the main router
	app.Get("/api", welcome)

	//Todo: loging the stages and the listeing the port
	log.Fatal(app.Listen(":3000"))
}
