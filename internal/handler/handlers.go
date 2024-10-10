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
	var task models.CreateTaskAPI

	if errBP := ctx.BodyParser(&task); errBP != nil {
		log.Printf("Error parsing body: %v\n", errBP)
		return h.handleValidationError(ctx, errBP.Error())
	}

	id, err := h.Service.Create(task)
	if err != nil {
		log.Printf("Error creating task: %v\n", err)
		return h.handleValidationError(ctx, "can't create task")
	}

	return ctx.Status(fiber.StatusOK).JSON(models.HTTPResponse{
		Error:   false,
		Message: id,
	})
}

func (h *TaskHandler) GetTasks(ctx *fiber.Ctx) error {
	tasks, err := h.Service.FetchAll()
	if err != nil {
		log.Printf("got error: %v", err)
		return h.handleValidationError(ctx, "can't fetch tasks")
	}

	if len(tasks) == 0 {
		return h.handleValidationError(ctx, "no tasks found")
	}

	ret := models.TasksJSON{
		Error: false,
		Tasks: tasks,
	}

	return ctx.Status(fiber.StatusOK).JSON(ret)
}

func (h *TaskHandler) GetTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	task, err := h.Service.Fetch(id)
	if err != nil {
		log.Printf("got error: %v\n", err)
		return h.handleValidationError(ctx, "can't fetch task")
	}

	return ctx.Status(fiber.StatusOK).JSON(task)
}

func (h *TaskHandler) UpdateTask(ctx *fiber.Ctx) error {
	var updatedTask models.TaskDB

	errBP := ctx.BodyParser(&updatedTask)
	if errBP != nil {
		log.Printf("got error parsing body: %v\n", errBP)
		return h.handleValidationError(ctx, "can't parse body")
	}

	if len(updatedTask.Title) == 0 {
		return h.handleValidationError(ctx, "can't have task title empty")
	}

	if len(updatedTask.Description) == 0 {
		return h.handleValidationError(ctx, "can't have task description empty")
	}

	id := ctx.Params("id")
	if err := h.Service.Update(id, updatedTask); err != nil {
		log.Printf("got error: %v\n", err)
		return h.handleValidationError(ctx, "can't update task")
	}

	return ctx.Status(fiber.StatusOK).JSON(models.HTTPResponse{
		Error:   false,
		Message: "",
	})
}

func (h *TaskHandler) DeleteTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := h.Service.Delete(id); err != nil {
		log.Printf("got error: %v\n", err)
		return h.handleValidationError(ctx, "can't delete task")
	}

	return ctx.Status(fiber.StatusOK).JSON(models.HTTPResponse{
		Error:   false,
		Message: "",
	})
}

func (h *TaskHandler) handleValidationError(ctx *fiber.Ctx, message string) error {
	br := models.HTTPResponse{
		Error:   true,
		Message: message,
	}
	return ctx.Status(fiber.StatusBadRequest).JSON(br)
}
