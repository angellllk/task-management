package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"time"
)

func CreateTask(ctx *fiber.Ctx, tasks map[string]Task) error {
	br := baseResponse{
		Error:   true,
		Message: "can't parse data",
	}

	var t Task

	errBP := ctx.BodyParser(&t)
	if errBP != nil {
		log.Printf("got error parsing body: %v\n", errBP)
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	if len(t.Title) == 0 {
		br.Message = "can't have task title empty"
		log.Printf("got error: %s", br.Message)
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	if len(t.Description) == 0 {
		br.Message = "can't have task description empty"
		log.Printf("got error: %s", br.Message)
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	t.ID = uuid.NewString()
	t.CreatedAt = time.Now()
	tasks[t.ID] = t

	return ctx.Status(fiber.StatusOK).JSON(baseResponse{
		Error:   false,
		Message: t.ID,
	})
}

func GetTasks(ctx *fiber.Ctx, tasks map[string]Task) error {
	br := baseResponse{
		Error:   true,
		Message: "no tasks created",
	}

	if len(tasks) == 0 {
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	var tasksList []Task
	for _, t := range tasks {
		tasksList = append(tasksList, t)
	}

	type TasksJSON struct {
		Error bool   `json:"error"`
		Tasks []Task `json:"tasks"`
	}

	ret := TasksJSON{
		Error: false,
		Tasks: tasksList,
	}

	return ctx.Status(fiber.StatusOK).JSON(ret)
}

func GetTask(ctx *fiber.Ctx, tasks map[string]Task) error {
	br := baseResponse{
		Error:   true,
		Message: "task not found",
	}

	task, found := tasks[ctx.Params("id")]
	if !found {
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	return ctx.Status(fiber.StatusOK).JSON(task)
}

func UpdateTask(ctx *fiber.Ctx, tasks map[string]Task) error {
	br := baseResponse{
		Error:   true,
		Message: "can't find task",
	}

	var updatedTask Task

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

	task, found := tasks[ctx.Params("id")]
	if !found {
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Completed = updatedTask.Completed
	tasks[task.ID] = task

	return ctx.Status(fiber.StatusOK).JSON(baseResponse{
		Error:   false,
		Message: "",
	})
}

func DeleteTask(ctx *fiber.Ctx, tasks map[string]Task) error {
	br := baseResponse{
		Error:   true,
		Message: "can't find task",
	}

	_, found := tasks[ctx.Params("id")]
	if !found {
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	delete(tasks, ctx.Params("id"))

	return ctx.Status(fiber.StatusOK).JSON(baseResponse{
		Error:   false,
		Message: "",
	})
}
