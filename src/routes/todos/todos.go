package routesTodos

import (
	handlerTodos "github.com/dwi-wijonarko/go-fiber-todo/src/handler/todos"
	"github.com/gofiber/fiber/v2"
)

func SetupTodoRoutes(router fiber.Router) {
	todo := router.Group("/todos")

	//create
	todo.Post("/", handlerTodos.CreateTodo)

	//read
	todo.Get("/", handlerTodos.GetTodos)

	//read one
	todo.Get("/:id", handlerTodos.GetTodo)

	//update
	todo.Put("/:id", handlerTodos.UpdateTodo)

	//delete
	todo.Delete("/:id", handlerTodos.DeleteTodo)
}
