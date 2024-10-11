package internal

import (
	"github.com/angellllk/task-management/config"
	"github.com/angellllk/task-management/internal/handlers"
	"github.com/angellllk/task-management/internal/repository"
	"github.com/angellllk/task-management/internal/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func StartServer() {
	cfg, errNew := config.New("./config.json")
	if errNew != nil {
		log.Printf("Error creating config: %v", errNew)
	}

	taskRepo, err := repository.New(cfg.Dsn)
	if err != nil {
		log.Printf("got error: %v", err)
	}

	taskService := &service.TaskService{Repo: taskRepo}
	taskHandler := &handlers.TaskHandler{Service: taskService}

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

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err = app.Listen(cfg.Host); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<-stop

	log.Println("Shutting down server...")
	if err = app.Shutdown(); err != nil {
		log.Fatalf("Error during shutdown: %v", err)
	}
}
