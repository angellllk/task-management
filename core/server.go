package core

import "github.com/gofiber/fiber/v2"

func StartServer() *fiber.App {
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

	return app
}
