package core

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"testing"
)

func TestCreateTask(t *testing.T) {
	app := StartServer()
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

	var ret baseResponse
	errUnmarshal := json.Unmarshal(respBody, &ret)
	if errUnmarshal != nil {
		t.Fatalf("got error: %v", errUnmarshal)
	}

	if ret.Error {
		t.Fatalf("got error: %s", ret.Message)
	}

	defer testCleanup(t, app, ret.Message)
}

func TestGetTasks(t *testing.T) {
	app := StartServer()
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

	var retCreate baseResponse
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

	var ret TasksJSON
	errUnmarshal = json.Unmarshal(respBody, &ret)
	if errUnmarshal != nil {
		t.Fatalf("got error: %v", errUnmarshal)
	}

	if retCreate.Error || len(ret.Tasks) == 0 {
		t.Fatalf("unexpected test result")
	}
}

func TestGetTask(t *testing.T) {
	app := StartServer()
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

	var ret baseResponse
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
	app := StartServer()
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

	var ret baseResponse
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
	app := StartServer()
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

	var ret baseResponse
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
