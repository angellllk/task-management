package models

import "time"

type HTTPResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

type TasksJSON struct {
	Error bool   `json:"error"`
	Tasks []Task `json:"tasks"`
}

type CreateTaskAPI struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
