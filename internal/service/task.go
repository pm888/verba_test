package service

import (
	"todo_list_verba/internal/model"
	"todo_list_verba/internal/repository"
)

type taskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task *model.Task) (*model.Task, error) {
	return s.repo.CreateTask(task)
}

func (s *taskService) GetTasks() ([]model.Task, error) {
	return s.repo.GetTasks()
}

func (s *taskService) GetTaskByID(id int) (*model.Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(id int, task *model.Task) (*model.Task, error) {
	return s.repo.UpdateTask(id, task)
}

func (s *taskService) DeleteTask(id int) error {
	return s.repo.DeleteTask(id)
}
