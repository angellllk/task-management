package handler

import (
	"github.com/angellllk/task-management/internal/models"
	"github.com/angellllk/task-management/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"time"
)

type TaskHandler struct {
	Service *service.TaskService
}

func (h *TaskHandler) CreateTask(ctx *fiber.Ctx) error {
	br := models.HTTPResponse{
		Error:   true,
		Message: "can't parse data",
	}

	var task models.Task

	if errBP := ctx.BodyParser(&task); errBP != nil {
		log.Printf("got error parsing body: %v\n", errBP)
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	if len(task.Title) == 0 {
		br.Message = "can't have task title empty"
		log.Printf("got error: %s", br.Message)
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	if len(task.Description) == 0 {
		br.Message = "can't have task description empty"
		log.Printf("got error: %s", br.Message)
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	task.ID = uuid.NewString()
	task.CreatedAt = time.Now()

	if err := h.Service.Create(task); err != nil {
		log.Printf("got error: %v\n", err)
		br.Message = "can't create task"
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	return ctx.Status(fiber.StatusOK).JSON(models.HTTPResponse{
		Error:   false,
		Message: task.ID,
	})
}

func (h *TaskHandler) GetTasks(ctx *fiber.Ctx) error {
	br := models.HTTPResponse{
		Error:   true,
		Message: "no tasks created",
	}

	tasks, err := h.Service.FetchAll()
	if err != nil {
		log.Printf("got error: %v\n", err)
		br.Message = "can't fetch tasks"
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	if len(tasks) == 0 {
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	ret := models.TasksJSON{
		Error: false,
		Tasks: tasks,
	}

	return ctx.Status(fiber.StatusOK).JSON(ret)
}

func (h *TaskHandler) GetTask(ctx *fiber.Ctx) error {
	br := models.HTTPResponse{
		Error:   true,
		Message: "task not found",
	}

	id := ctx.Params("id")
	task, err := h.Service.Fetch(id)
	if err != nil {
		log.Printf("got error: %v\n", err)
		br.Message = "can't fetch task"
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	return ctx.Status(fiber.StatusOK).JSON(task)
}

func (h *TaskHandler) UpdateTask(ctx *fiber.Ctx) error {
	br := models.HTTPResponse{
		Error:   true,
		Message: "can't find task",
	}

	var updatedTask models.Task

	errBP := ctx.BodyParser(&updatedTask)
	if errBP != nil {
		log.Printf("got error parsing body: %v\n", errBP)
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	if len(updatedTask.Title) == 0 {
		br.Message = "can't have task title empty"
		log.Printf("got error: %s", br.Message)
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	if len(updatedTask.Description) == 0 {
		br.Message = "can't have task description empty"
		log.Printf("got error: %s", br.Message)
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	id := ctx.Params("id")
	if err := h.Service.Update(id, updatedTask); err != nil {
		log.Printf("got error: %v\n", err)
		br.Message = "can't update task"
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	return ctx.Status(fiber.StatusOK).JSON(models.HTTPResponse{
		Error:   false,
		Message: "",
	})
}

func (h *TaskHandler) DeleteTask(ctx *fiber.Ctx) error {
	br := models.HTTPResponse{
		Error:   true,
		Message: "can't find task",
	}

	id := ctx.Params("id")
	if err := h.Service.Delete(id); err != nil {
		log.Printf("got error: %v\n", err)
		br.Message = "can't delete task"
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	return ctx.Status(fiber.StatusOK).JSON(models.HTTPResponse{
		Error:   false,
		Message: "",
	})
}
