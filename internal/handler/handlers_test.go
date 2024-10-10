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
		body    models.CreateTaskAPI
		wantErr bool
	}{
		{
			"create-task",
			models.CreateTaskAPI{
				Title:       "test",
				Description: "test",
			},
			false,
		},
		{
			"create-task-empty-title",
			models.CreateTaskAPI{
				Title:       "",
				Description: "test",
			},
			true,
		},
		{
			"create-task-empty-description",
			models.CreateTaskAPI{
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

			var ret models.HTTPResponse
			errUnmarshal := json.Unmarshal(respBody, &ret)
			if errUnmarshal != nil {
				t.Fatalf("got error: %v", errUnmarshal)
			}

			hasError := ret.Error != false
			if hasError != tt.wantErr {
				t.Fatal()
			}

			if !hasError {
				defer testCleanup(t, app, ret.Message)
			}
		})
	}
}

func TestGetTasks(t *testing.T) {
	app := testServer(t)
	url := "http://localhost:8080/tasks"
	body := `{"title": "test", "description":"test"}`

	bodyBytes := []byte(body)

	createTaskReq, errReq := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewReader(bodyBytes))
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

	var retCreate models.HTTPResponse
	errUnmarshal := json.Unmarshal(respBody, &retCreate)
	if errUnmarshal != nil {
		t.Fatalf("got error: %v", errUnmarshal)
	}

	defer testCleanup(t, app, retCreate.Message)

	getTasksReq, errReq := http.NewRequest(
		http.MethodGet,
		url,
		nil)
	if errReq != nil {
		t.Fatalf("got error: %v", errReq)
	}

	resp, err = app.Test(getTasksReq, -1)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	respBody, errRead = io.ReadAll(resp.Body)
	if errRead != nil {
		t.Fatalf("got error: %v", errRead)
	}

	var ret models.TasksJSON
	errUnmarshal = json.Unmarshal(respBody, &ret)
	if errUnmarshal != nil {
		t.Fatalf("got error: %v", errUnmarshal)
	}

	if retCreate.Error || len(ret.Tasks) == 0 {
		t.Fatalf("unexpected test result")
	}
}

func TestGetTask(t *testing.T) {
	app := testServer(t)
	url := "http://localhost:8080/tasks"
	body := `{"title": "test", "description":"test"}`

	bodyBytes := []byte(body)

	createTaskReq, errReq := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewReader(bodyBytes))
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

	var ret models.HTTPResponse
	errUnmarshal := json.Unmarshal(respBody, &ret)
	if errUnmarshal != nil {
		t.Fatalf("got error: %v", errUnmarshal)
	}

	defer testCleanup(t, app, ret.Message)

	id := ret.Message
	getTasksUrl := "http://localhost:8080/tasks/" + id
	getTaskReq, errReq := http.NewRequest(
		http.MethodGet,
		getTasksUrl,
		nil)

	resp, err = app.Test(getTaskReq, -1)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	if resp.Body == nil {
		t.Fatalf("unexpected test result")
	}
}

func TestUpdateTask(t *testing.T) {
	app := testServer(t)
	url := "http://localhost:8080/tasks"
	body := `{"title": "test", "description":"test"}`

	bodyBytes := []byte(body)

	createTaskReq, errReq := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewReader(bodyBytes))
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

	var ret models.HTTPResponse
	errUnmarshal := json.Unmarshal(respBody, &ret)
	if errUnmarshal != nil {
		t.Fatalf("got error: %v", errUnmarshal)
	}

	defer testCleanup(t, app, ret.Message)

	body = `{"title":"new-title", "description":"new-description"}`
	id := ret.Message
	updateTaskUrl := "http://localhost:8080/tasks/" + id
	updateTaskReq, errReq := http.NewRequest(
		http.MethodPut,
		updateTaskUrl,
		bytes.NewReader([]byte(body)))

	updateTaskReq.Header.Set("Content-Type", "application/json")
	resp, err = app.Test(updateTaskReq, -1)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	errUnmarshal = json.Unmarshal(respBody, &ret)
	if errUnmarshal != nil {
		t.Fatalf("got error: %v", errUnmarshal)
	}

	if ret.Error {
		t.Fatalf("unexpected test result")
	}
}

func TestDeleteTask(t *testing.T) {
	app := testServer(t)
	url := "http://localhost:8080/tasks"
	body := `{"title": "test", "description":"test"}`

	bodyBytes := []byte(body)

	createTaskReq, errReq := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewReader(bodyBytes))
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

	var ret models.HTTPResponse
	errUnmarshal := json.Unmarshal(respBody, &ret)
	if errUnmarshal != nil {
		t.Fatalf("got error: %v", errUnmarshal)
	}

	id := ret.Message
	deleteTaskUrl := "http://localhost:8080/tasks/" + id
	deleteTaskReq, errReq := http.NewRequest(
		http.MethodDelete,
		deleteTaskUrl,
		nil)
	if errReq != nil {
		t.Fatalf("got error: %v", errReq)
	}

	resp, err = app.Test(deleteTaskReq, -1)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	respBody, errRead = io.ReadAll(resp.Body)
	if errRead != nil {
		t.Fatalf("got error: %v", errRead)
	}

	errUnmarshal = json.Unmarshal(respBody, &ret)
	if errUnmarshal != nil {
		t.Fatalf("got error: %v", errUnmarshal)
	}

	if ret.Error {
		t.Fatalf("unexpected test result")
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

	var ret models.HTTPResponse
	if err = json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		t.Fatalf("got error: %v", err)
	}

	if ret.Error {
		t.Fatalf("got error: %s", ret.Message)
	}

	return ret.Message
}
