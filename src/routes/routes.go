package routes

import (
	todosRoutes "github.com/dwi-wijonarko/go-fiber-todo/src/routes/todos"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	todosRoutes.SetupTodoRoutes(api)
}
