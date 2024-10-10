package core

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func StartServer() *fiber.App {
	app := fiber.New()
	db, err := GetDBConnection()
	if err != nil {
		log.Printf("got error: %v", err)
		return nil
	}

	app.Post("/tasks", func(ctx *fiber.Ctx) error {
		return CreateTask(ctx, db)
	})

	app.Get("/tasks", func(ctx *fiber.Ctx) error {
		return GetTasks(ctx, db)
	})

	app.Get("/tasks/:id", func(ctx *fiber.Ctx) error {
		return GetTask(ctx, db)
	})

	app.Put("/tasks/:id", func(ctx *fiber.Ctx) error {
		return UpdateTask(ctx, db)
	})

	app.Delete("/tasks/:id", func(ctx *fiber.Ctx) error {
		return DeleteTask(ctx, db)
	})

	return app
}
