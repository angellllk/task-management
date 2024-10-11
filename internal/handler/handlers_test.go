package handler

import (
	"bytes"
	"encoding/json"
	"github.com/angellllk/task-management/internal/models"
	"github.com/angellllk/task-management/internal/repository"
	"github.com/angellllk/task-management/internal/service"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"testing"
)

func TestCreateTask(t *testing.T) {
	app := testServer(t)
	url := "http://localhost:8080/tasks"

	tests := []struct {
		name    string
		body    models.TaskCreateDTO
		wantErr bool
	}{
		{
			"create-task",
			models.TaskCreateDTO{
				Title:       "test",
				Description: "test",
			},
			false,
		},
		{
			"create-task-empty-title",
			models.TaskCreateDTO{
				Title:       "",
				Description: "test",
			},
			true,
		},
		{
			"create-task-empty-description",
			models.TaskCreateDTO{
				Title:       "test",
				Description: "",
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bodyBytes, errMarshal := json.Marshal(tt.body)
			if errMarshal != nil {
				t.Fatalf("got error: %v", errMarshal)
			}

			createTaskReq, errReq := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyBytes))
			if errReq != nil {
				t.Fatalf("got error: %v", errReq)
			}

			createTaskReq.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(createTaskReq, -1)
			if err != nil {
				t.Fatalf("got error: %v", err)
			}

			respBody, errRead := io.ReadAll(resp.Body)
			if errRead != nil {
				t.Fatalf("got error: %v", errRead)
			}

			var ret models.TaskResponseJSON
			errUnmarshal := json.Unmarshal(respBody, &ret)
			if errUnmarshal != nil {
				t.Fatalf("got error: %v", errUnmarshal)
			}

			hasError := ret.Error != false
			if hasError != tt.wantErr {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, ret.Error)
			}

			if !hasError {
				defer testCleanup(t, app, ret.Task.ID)
			}
		})
	}
}

func TestGetTasks(t *testing.T) {
	app := testServer(t)
	url := "http://localhost:8080/tasks"

	tests := []struct {
		name      string
		wantTasks bool
		wantErr   bool
	}{
		{
			"get-tasks",
			true,
			false,
		},
		{
			"get-tasks-with-body",
			true,
			false,
		},
		{
			"get-no-tasks-created",
			false,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantTasks {
				id := createTestTask(t, app)
				defer testCleanup(t, app, id)
			}

			getTasksReq, errReq := http.NewRequest(http.MethodGet, url, nil)
			if errReq != nil {
				t.Fatalf("got error: %v", errReq)
			}

			resp, err := app.Test(getTasksReq, -1)
			if err != nil {
				t.Fatalf("got error: %v", err)
			}

			respBody, errRead := io.ReadAll(resp.Body)
			if errRead != nil {
				t.Fatalf("got error: %v", errRead)
			}

			var ret models.TasksResponseJson
			errUnmarshal := json.Unmarshal(respBody, &ret)
			if errUnmarshal != nil {
				t.Fatalf("got error: %v", errUnmarshal)
			}

			hasError := ret.Error != false
			if hasError != tt.wantErr {
				t.Fatal("got wrong error value")
			}

			hasTasks := ret.Tasks != nil
			if hasTasks != tt.wantTasks {
				t.Fatal("got wrong tasks value")
			}
		})
	}
}

func TestGetTask(t *testing.T) {
	app := testServer(t)
	id := createTestTask(t, app)
	defer testCleanup(t, app, id)

	url := "http://localhost:8080/tasks/"

	tests := []struct {
		name     string
		url      string
		wantTask bool
		wantErr  bool
	}{
		{
			"get-task",
			url + id,
			true,
			false,
		},
		{
			"get-task-invalid-id",
			url + "test",
			false,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getTaskReq, errReq := http.NewRequest(http.MethodGet, tt.url, nil)
			if errReq != nil {
				t.Fatalf("got error: %v", errReq)
			}

			resp, err := app.Test(getTaskReq, -1)
			if err != nil {
				t.Fatalf("got error: %v", err)
			}

			respBody, errRead := io.ReadAll(resp.Body)
			if errRead != nil {
				t.Fatalf("got error: %v", errRead)
			}

			var ret models.TaskResponseJSON
			errUnmarshal := json.Unmarshal(respBody, &ret)
			if errUnmarshal != nil {
				t.Fatalf("got error: %v", errUnmarshal)
			}

			hasError := ret.Error != false
			if hasError != tt.wantErr {
				t.Fatal("got wrong error value")
			}

			hasTask := ret.Task.ID != ""
			if hasTask != tt.wantTask {
				t.Fatal("got wrong tasks value")
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	app := testServer(t)
	id := createTestTask(t, app)
	defer testCleanup(t, app, id)

	url := "http://localhost:8080/tasks/"

	tests := []struct {
		name     string
		url      string
		body     models.TaskUpdateDTO
		wantTask bool
		wantErr  bool
	}{
		{
			"update-task",
			url + id,
			models.TaskUpdateDTO{
				Title:       "updated",
				Description: "updated",
			},
			true,
			false,
		},
		{
			"invalid-id",
			url + "test",
			models.TaskUpdateDTO{
				Title:       "update",
				Description: "update",
			},
			false,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bodyBytes, errMarshal := json.Marshal(tt.body)
			if errMarshal != nil {
				t.Fatalf("got error: %v", errMarshal)
			}

			updateTaskReq, errReq := http.NewRequest(http.MethodPut, tt.url, bytes.NewReader(bodyBytes))
			if errReq != nil {
				t.Fatalf("got error: %v", errReq)
			}

			updateTaskReq.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(updateTaskReq, -1)
			if err != nil {
				t.Fatalf("got error: %v", err)
			}

			respBody, errRead := io.ReadAll(resp.Body)
			if errRead != nil {
				t.Fatalf("got error: %v", errRead)
			}

			var ret models.TaskResponseJSON
			errUnmarshal := json.Unmarshal(respBody, &ret)
			if errUnmarshal != nil {
				t.Fatalf("got error: %v", errUnmarshal)
			}

			hasError := ret.Error != false
			if hasError != tt.wantErr {
				t.Fatal("got wrong error value")
			}

			if tt.wantTask {
				getResp := getTestTask(t, app, id)
				if getResp.Task.Title != tt.body.Title {
					t.Fatal("got wrong task result")
				}
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	app := testServer(t)
	id := createTestTask(t, app)

	url := "http://localhost:8080/tasks/"

	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			"delete-task",
			url + id,
			false,
		},
		{
			"invalid-id",
			url + "test",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteTaskReq, errReq := http.NewRequest(http.MethodDelete, tt.url, nil)
			if errReq != nil {
				t.Fatalf("got error: %v", errReq)
			}

			resp, err := app.Test(deleteTaskReq, -1)
			if err != nil {
				t.Fatalf("got error: %v", err)
			}

			respBody, errRead := io.ReadAll(resp.Body)
			if errRead != nil {
				t.Fatalf("got error: %v", errRead)
			}

			var ret models.TasksResponseJson
			errUnmarshal := json.Unmarshal(respBody, &ret)
			if errUnmarshal != nil {
				t.Fatalf("got error: %v", errUnmarshal)
			}

			hasError := ret.Error != false
			if hasError != tt.wantErr {
				t.Fatal("got wrong error value")
			}
		})
	}
}

func testServer(t *testing.T) *fiber.App {
	t.Helper()

	taskRepo, err := repository.New()
	if err != nil {
		t.Logf("got error: %v", err)
	}

	taskService := &service.TaskService{Repo: taskRepo}
	taskHandler := &TaskHandler{Service: taskService}

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

	return app
}

func testCleanup(t *testing.T, app *fiber.App, id string) {
	t.Helper()

	deleteTaskUrl := "http://localhost:8080/tasks/" + id
	deleteTaskReq, errReq := http.NewRequest(
		http.MethodDelete,
		deleteTaskUrl,
		nil)
	if errReq != nil {
		t.Fatalf("got error: %v", errReq)
	}

	_, err := app.Test(deleteTaskReq, -1)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}
}

func createTestTask(t *testing.T, app *fiber.App) string {
	url := "http://localhost:8080/tasks"
	body := `{"title": "test", "description":"test"}`

	createTaskReq, errReq := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	if errReq != nil {
		t.Fatalf("got error: %v", errReq)
	}

	createTaskReq.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(createTaskReq, -1)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	var ret models.TaskResponseJSON
	if err = json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		t.Fatalf("got error: %v", err)
	}

	return ret.Task.ID
}

func getTestTask(t *testing.T, app *fiber.App, id string) models.TaskResponseJSON {
	url := "http://localhost:8080/tasks/" + id

	getTaskReq, errReq := http.NewRequest(http.MethodGet, url, nil)
	if errReq != nil {
		t.Fatalf("got error: %v", errReq)
	}
	resp, err := app.Test(getTaskReq, -1)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	var ret models.TaskResponseJSON
	if err = json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		t.Fatalf("got error: %v", err)
	}

	return ret
}
