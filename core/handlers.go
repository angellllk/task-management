package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

func CreateTask(ctx *fiber.Ctx, db *gorm.DB) error {
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

	db.Create(t)

	return ctx.Status(fiber.StatusOK).JSON(baseResponse{
		Error:   false,
		Message: t.ID,
	})
}

func GetTasks(ctx *fiber.Ctx, db *gorm.DB) error {
	br := baseResponse{
		Error:   true,
		Message: "no tasks created",
	}

	var tasks []Task
	db.Find(&tasks)

	if len(tasks) == 0 {
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	ret := TasksJSON{
		Error: false,
		Tasks: tasks,
	}

	return ctx.Status(fiber.StatusOK).JSON(ret)
}

func GetTask(ctx *fiber.Ctx, db *gorm.DB) error {
	br := baseResponse{
		Error:   true,
		Message: "task not found",
	}

	var t Task
	id := ctx.Params("id")
	db.Where("id = ?", id).First(&t)

	if len(t.Title) == 0 {
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	return ctx.Status(fiber.StatusOK).JSON(t)
}

func UpdateTask(ctx *fiber.Ctx, db *gorm.DB) error {
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

	var task Task
	id := ctx.Params("id")
	tx := db.Where("id = ?", id).First(&task)
	if tx.Error != nil {
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	tx = db.Model(&task).Updates(updatedTask)
	if tx.Error != nil {
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	return ctx.Status(fiber.StatusOK).JSON(baseResponse{
		Error:   false,
		Message: "",
	})
}

func DeleteTask(ctx *fiber.Ctx, db *gorm.DB) error {
	br := baseResponse{
		Error:   true,
		Message: "can't find task",
	}

	id := ctx.Params("id")
	tx := db.Where("id = ?", id).Delete(&Task{})
	if tx.Error != nil {
		return ctx.Status(fiber.StatusOK).JSON(br)
	}

	return ctx.Status(fiber.StatusOK).JSON(baseResponse{
		Error:   false,
		Message: "",
	})
}
