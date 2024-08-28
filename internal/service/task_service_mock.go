package service

import (
	"todo_list_verba/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) CreateTask(task *model.Task) (*model.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*model.Task), args.Error(1)
}

func (m *MockTaskService) GetTasks() ([]model.Task, error) {
	args := m.Called()
	return args.Get(0).([]model.Task), args.Error(1)
}

func (m *MockTaskService) GetTaskByID(id int) (*model.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Task), args.Error(1)
}

func (m *MockTaskService) UpdateTask(id int, task *model.Task) (*model.Task, error) {
	args := m.Called(id, task)
	return args.Get(0).(*model.Task), args.Error(1)
}

func (m *MockTaskService) DeleteTask(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
