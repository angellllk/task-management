package handler

import (
	"github.com/angellllk/task-management/internal/models"
	"github.com/angellllk/task-management/internal/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

type TaskHandler struct {
	Service *service.TaskService
}

func (h *TaskHandler) CreateTask(ctx *fiber.Ctx) error {
	var dto models.TaskCreateDTO

	if errBP := ctx.BodyParser(&dto); errBP != nil {
		log.Printf("Error parsing body: %v\n", errBP)
		return h.handleValidationError(ctx, errBP.Error())
	}

	task, err := h.Service.Create(dto)
	if err != nil {
		log.Printf("Error creating task: %v\n", err)
		return h.handleValidationError(ctx, "can't create task")
	}

	resp := models.TaskResponseJSON{
		Error: false,
		Task:  *task,
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (h *TaskHandler) GetTasks(ctx *fiber.Ctx) error {
	tasks, err := h.Service.FetchAll()
	if err != nil {
		log.Printf("Error fetching all tasks: %v", err)
		return h.handleValidationError(ctx, "can't fetch tasks")
	}

	if len(tasks) == 0 {
		return h.handleValidationError(ctx, "no tasks found")
	}

	resp := models.TasksResponseJson{
		Error: false,
		Tasks: tasks,
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (h *TaskHandler) GetTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	task, err := h.Service.Fetch(id)
	if err != nil {
		log.Printf("Error fetching task for id %s: %v\n", id, err)
		return h.handleValidationError(ctx, "can't fetch task")
	}

	resp := models.TaskResponseJSON{
		Error: false,
		Task:  *task,
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (h *TaskHandler) UpdateTask(ctx *fiber.Ctx) error {
	var updatedTask models.TaskUpdateDTO

	errBP := ctx.BodyParser(&updatedTask)
	if errBP != nil {
		log.Printf("Error parsing body: %v\n", errBP)
		return h.handleValidationError(ctx, "can't parse body")
	}

	id := ctx.Params("id")
	task, err := h.Service.Update(id, updatedTask)
	if err != nil {
		log.Printf("Error updating task with id %s: %v\n", id, err)
		return h.handleValidationError(ctx, "can't update task")
	}

	resp := models.TaskResponseJSON{
		Error: false,
		Task:  *task,
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (h *TaskHandler) DeleteTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := h.Service.Delete(id); err != nil {
		log.Printf("Error deleting task with id %s: %v\n", id, err)
		return h.handleValidationError(ctx, "can't delete task")
	}

	resp := models.TaskResponseJSON{
		Error: false,
		Task:  models.TaskResponseDTO{},
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (h *TaskHandler) handleValidationError(ctx *fiber.Ctx, message string) error {
	br := models.TaskResponseJSON{
		Error:   true,
		Message: message,
	}
	return ctx.Status(fiber.StatusBadRequest).JSON(br)
}
