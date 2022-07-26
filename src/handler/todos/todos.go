package handlerTodos

import (
	"github.com/dwi-wijonarko/go-fiber-todo/database"
	"github.com/dwi-wijonarko/go-fiber-todo/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetTodos(c *fiber.Ctx) error {
	db := database.DB
	var todos []models.Todo

	//find all todos
	db.Find(&todos)

	//if no todos found
	if len(todos) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No todos found",
		})
	}

	//return all todos
	return c.JSON(fiber.Map{"status": "success", "data": todos})
}

func CreateTodo(c *fiber.Ctx) error {
	db := database.DB
	var todo models.Todo

	//get the body of the request
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	//create UUID
	todo.ID = uuid.New()

	//create the todo
	err := db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	//return the todo
	return c.JSON(fiber.Map{"status": "success", "data": todo})
}

func GetTodo(c *fiber.Ctx) error {
	db := database.DB
	var todo models.Todo

	//get the id from the url
	id := c.Params("id")

	//find the todo
	db.First(&todo, "id=?", id)

	//if no todo found
	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "No todo found",
		})
	}

	//return the todo
	return c.JSON(fiber.Map{"status": "success", "data": todo})
}

func UpdateTodo(c *fiber.Ctx) error {
	type UpdateTodo struct {
		Title       string
		Description string
		Done        bool
	}

	db := database.DB
	var todo models.Todo

	//get the id from the url
	id := c.Params("id")

	//find the todo
	db.First(&todo, "id=?", id)

	//if no todo found
	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "No todo found",
		})
	}

	// Store the body containing the updated data and return error if encountered
	var updateTodoData UpdateTodo
	err := c.BodyParser(&updateTodoData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	todo.Title = updateTodoData.Title
	todo.Description = updateTodoData.Description
	todo.Done = updateTodoData.Done

	//update the todo
	db.Save(&todo)

	//return the todo
	return c.JSON(fiber.Map{"status": "success", "data": todo})
}

func DeleteTodo(c *fiber.Ctx) error {
	db := database.DB
	var todo models.Todo

	//get the id from the url
	id := c.Params("id")

	//find the todo
	db.First(&todo, "id=?", id)

	//if no todo found
	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "No todo found",
		})
	}

	//delete the todo
	err := db.Delete(&todo).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	//return the todo
	return c.JSON(fiber.Map{"status": "success", "data": todo})
}
