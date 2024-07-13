package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"todo_list/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/qinains/fastergoding"
)

func main() {
	fastergoding.Run() // hot reload
	app := fiber.New()

	todos := []models.Todo{}

	app.Post("/todos", func(c *fiber.Ctx) error {
		// string data type id auto generated
		id := strings.ReplaceAll(uuid.New().String(), "-", "")
		body := c.Body()
		fmt.Println(string(body))
		title := c.FormValue("title")
		fmt.Println(title)
		todo := models.Todo{
			ID:        id,
			Title:     c.FormValue("title"),
			Status:    "active",
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		todos = append(todos, todo)
		// return JSON response with status code
		c.Status(fiber.StatusCreated)
		return c.JSON(todo)
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":3000"))
}
