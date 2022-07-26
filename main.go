package main

import (
	"github.com/dwi-wijonarko/go-fiber-todo/database"
	"github.com/dwi-wijonarko/go-fiber-todo/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//connect to database
	database.ConnectDB()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
