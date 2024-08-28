package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"todo_list_verba/internal/model"
	"todo_list_verba/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	mockService := &service.MockTaskService{}
	handler := NewTaskHandler(mockService)

	task := &model.Task{Title: "New Task", Description: "Task Description"}
	createdTask := &model.Task{ID: 1, Title: "New Task", Description: "Task Description"}
	mockService.On("CreateTask", task).Return(createdTask, nil)

	body, _ := json.Marshal(task)
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.CreateTask(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var responseTask model.Task
	err = json.NewDecoder(rr.Body).Decode(&responseTask)
	assert.NoError(t, err)
	assert.Equal(t, createdTask.ID, responseTask.ID)
	assert.Equal(t, createdTask.Title, responseTask.Title)
	assert.Equal(t, createdTask.Description, responseTask.Description)
}

func TestGetTasks(t *testing.T) {
	mockService := &service.MockTaskService{}
	handler := NewTaskHandler(mockService)

	tasks := []model.Task{
		{ID: 1, Title: "Task 1", Description: "Description 1"},
		{ID: 2, Title: "Task 2", Description: "Description 2"},
	}
	mockService.On("GetTasks").Return(tasks, nil)

	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.GetTasks(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var returnedTasks []model.Task
	err = json.NewDecoder(rr.Body).Decode(&returnedTasks)
	assert.NoError(t, err)
	assert.Equal(t, len(tasks), len(returnedTasks))
	assert.Equal(t, tasks[0].Title, returnedTasks[0].Title)
	assert.Equal(t, tasks[1].Description, returnedTasks[1].Description)
}
