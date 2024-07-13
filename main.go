package main

import (
	"fmt"
	"log"
	"strings"
	"todo_list/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func main() {
	app := fiber.New()

	todos := []models.Todo{}

	app.Post("/todos", func(c *fiber.Ctx) error {
		// string data type id auto generated
		id := strings.ReplaceAll(uuid.New().String(), "-", "")
		todo := models.Todo{}
		if err := c.BodyParser(&todo); err != nil {
			return err
		}
		fmt.Printf("todo %+v\n", todo)
		todo.ID = &id
		todo.Status = &[]string{"active"}[0]
		todo.CreatedAt = &[]string{"2021-09-01T00:00:00Z"}[0]
		todos = append(todos, todo)
		return c.JSON(todo)
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":3000"))
}
