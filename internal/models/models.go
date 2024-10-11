package models

import (
	"errors"
	"time"
)

type HTTPResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type TaskDB struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

type TasksJSON struct {
	Error bool     `json:"error"`
	Tasks []TaskDB `json:"tasks"`
}

type TaskAPI struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t *TaskAPI) Validate() error {
	if t.Title == "" {
		return errors.New("title cannot be empty")
	}
	if t.Description == "" {
		return errors.New("description cannot be empty")
	}
	return nil
}
