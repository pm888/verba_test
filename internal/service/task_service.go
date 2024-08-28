package service

import "todo_list_verba/internal/model"

type TaskService interface {
	CreateTask(task *model.Task) (*model.Task, error)
	GetTasks() ([]model.Task, error)
	GetTaskByID(id int) (*model.Task, error)
	UpdateTask(id int, task *model.Task) (*model.Task, error)
	DeleteTask(id int) error
}
