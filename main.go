package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type baseResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	app := fiber.New()
	tasks := make(map[string]Task)

	app.Post("/tasks", func(ctx *fiber.Ctx) error {
		return CreateTask(ctx, tasks)
	})

	app.Get("/tasks", func(ctx *fiber.Ctx) error {
		return GetTasks(ctx, tasks)
	})

	app.Get("/tasks/:id", func(ctx *fiber.Ctx) error {
		return GetTask(ctx, tasks)
	})

	app.Put("/tasks/:id", func(ctx *fiber.Ctx) error {
		return UpdateTask(ctx, tasks)
	})

	app.Delete("/tasks/:id", func(ctx *fiber.Ctx) error {
		return DeleteTask(ctx, tasks)
	})

	app.Listen(":8080")
}
