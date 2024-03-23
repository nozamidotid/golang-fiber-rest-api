package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozamidotid/golang-fiber-rest-api/database"
	"github.com/nozamidotid/golang-fiber-rest-api/models"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Ajitama!")
	})

	app.Get("/todo", getAllTodos)
	app.Post("/todo", createTodo)
	app.Get("/todo/:id", getTodoById)
	app.Put("/todo/:id", updateTodo)
	app.Delete("/todo/:id", deleteTodo)

	app.Listen(":3000")
}

func getAllTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	database.DB.Db.Find(&todos)
	return c.Status(fiber.StatusOK).JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&todo)
	return c.Status(fiber.StatusCreated).JSON(todo)
}

func getTodoById(c *fiber.Ctx) error {
	todo := &models.Todo{}
	id := c.Params("id")
	if err := database.DB.Db.First(todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	todo := &models.Todo{}
	id := c.Params("id")
	if err := database.DB.Db.First(todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Save(todo)
	return c.Status(fiber.StatusOK).JSON(todo)
}

func deleteTodo(c *fiber.Ctx) error {
	todo := &models.Todo{}
	id := c.Params("id")
	if err := database.DB.Db.First(todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	database.DB.Db.Delete(todo, id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}
