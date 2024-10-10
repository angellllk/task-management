package internal

import (
	"github.com/angellllk/task-management/internal/handler"
	"github.com/angellllk/task-management/internal/repository"
	"github.com/angellllk/task-management/internal/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

func StartServer() {
	taskRepo, err := repository.New()
	if err != nil {
		log.Printf("got error: %v", err)
	}

	taskService := &service.TaskService{Repo: taskRepo}
	taskHandler := &handler.TaskHandler{Service: taskService}

	app := fiber.New()
	app.Post("/tasks", func(ctx *fiber.Ctx) error {
		return taskHandler.CreateTask(ctx)
	})
	app.Get("/tasks", func(ctx *fiber.Ctx) error {
		return taskHandler.GetTasks(ctx)
	})
	app.Get("/tasks/:id", func(ctx *fiber.Ctx) error {
		return taskHandler.GetTask(ctx)
	})
	app.Put("/tasks/:id", func(ctx *fiber.Ctx) error {
		return taskHandler.UpdateTask(ctx)
	})
	app.Delete("/tasks/:id", func(ctx *fiber.Ctx) error {
		return taskHandler.DeleteTask(ctx)
	})

	app.Listen(":8080")
}
