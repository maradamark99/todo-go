package main

import (
	"fmt"
	"time"
	"todo/m/v2/scheduler"
	"todo/m/v2/todo"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

func main() {
	app := fiber.New()
	todoStorage := &todo.InMemoryTodoStorage{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(todoStorage.GetAll())
	})

	app.Post("/", func(c *fiber.Ctx) error {
		body := new(todo.Todo)
		err := c.BodyParser(body)
		if err != nil {
			c.Status(fiber.StatusBadRequest).SendString(err.Error())
			return err
		}

		// maybe create a seperate validator func or smth
		now := time.Now()
		if !body.Deadline.IsZero() && body.Deadline.Before(now) {
			c.Status(fiber.StatusBadRequest).SendString("Invalid deadline given.")
			return err
		}

		if body.Deadline.After(now) {
			go scheduler.Schedule(body.Deadline.Sub(now), func() {
				log.Info(body)
			})
		}
		todo := todo.CreateTodo(body.Name, todo.WithDescription(body.Description))
		todoStorage.Store(todo)
		return c.Status(fiber.StatusCreated).JSON(todo.Id)
	})

	app.Delete("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			c.Status(fiber.StatusBadRequest).SendString(err.Error())
			return err
		}
		todos := todoStorage.Todos
		// move this into a separate func that takes another func as arg
		for i, e := range todos {
			if e.Id == uuid {
				todoStorage.Todos = append(todos[:i], todos[i+1:]...)
				c.Status(fiber.StatusOK)
				return nil
			}
		}
		return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("Todo with id: %s not found", id))
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			c.Status(fiber.StatusBadRequest).SendString(err.Error())
			return err
		}
		todos := todoStorage.Todos
		for _, e := range todos {
			if e.Id == uuid {
				c.Status(fiber.StatusOK).JSON(e)
				return nil
			}
		}
		return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("Todo with id: %s not found", id))
	})

	app.Listen(":3000")
}
