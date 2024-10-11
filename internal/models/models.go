package models

import (
	"errors"
	"time"
)

type TaskDB struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

func (tDB *TaskDB) ToDTO() *TaskResponseDTO {
	dto := TaskResponseDTO{
		ID:          tDB.ID,
		Title:       tDB.Title,
		Description: tDB.Description,
		Completed:   tDB.Completed,
		CreatedAt:   tDB.CreatedAt,
	}
	return &dto
}

type TaskResponseJSON struct {
	Error   bool            `json:"error"`
	Message string          `json:"message"`
	Task    TaskResponseDTO `json:"task"`
}

type TasksResponseJson struct {
	Error bool              `json:"error"`
	Tasks []TaskResponseDTO `json:"tasks"`
}

// TaskCreateDTO defines the fields needed to create a new task
type TaskCreateDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (t *TaskCreateDTO) Validate() error {
	if t.Title == "" {
		return errors.New("title cannot be empty")
	}
	if t.Description == "" {
		return errors.New("description cannot be empty")
	}
	return nil
}

// TaskUpdateDTO defines the fields allowed when updating an existing task
type TaskUpdateDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Completed   bool   `json:"completed" validate:"required"`
}

func (t *TaskUpdateDTO) Validate() error {
	if t.Title == "" {
		return errors.New("title cannot be empty")
	}
	if t.Description == "" {
		return errors.New("description cannot be empty")
	}
	return nil
}

// TaskResponseDTO defines the structure of data sent back to the client
type TaskResponseDTO struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}
