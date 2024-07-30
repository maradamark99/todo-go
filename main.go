package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type InMemoryTodoStorage struct {
	Todos []*Todo
}

func (i *InMemoryTodoStorage) GetAll() []*Todo {
	return i.Todos
}

func (i *InMemoryTodoStorage) Store(t *Todo) {
	i.Todos = append(i.Todos, t)
}

func main() {
	app := fiber.New()
	todoStorage := &InMemoryTodoStorage{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(todoStorage.GetAll())
	})

	app.Post("/", func(c *fiber.Ctx) error {
		body := new(Todo)
		err := c.BodyParser(body)
		if err != nil {
			c.Status(fiber.StatusBadRequest).SendString(err.Error())
			return err
		}

		todo := CreateTodo(body.Name, WithDescription(body.Description))
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
