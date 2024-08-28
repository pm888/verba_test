package repository

import (
	"database/sql"
	"time"

	"todo_list_verba/internal/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task *model.Task) (*model.Task, error) {
	var id int
	err := r.db.QueryRow(`
        INSERT INTO tasks (title, description, due_date, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `, task.Title, task.Description, task.DueDate, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return nil, err
	}
	task.ID = id
	return task, nil
}

func (r *TaskRepository) GetTasks() ([]model.Task, error) {
	rows, err := r.db.Query(`
        SELECT id, title, description, due_date, created_at, updated_at
        FROM tasks
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) GetTaskByID(id int) (*model.Task, error) {
	var task model.Task
	err := r.db.QueryRow(`
        SELECT id, title, description, due_date, created_at, updated_at
        FROM tasks
        WHERE id = $1
    `, id).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) UpdateTask(id int, task *model.Task) (*model.Task, error) {
	_, err := r.db.Exec(`
        UPDATE tasks
        SET title = $1, description = $2, due_date = $3, updated_at = $4
        WHERE id = $5
    `, task.Title, task.Description, task.DueDate, time.Now(), id)
	if err != nil {
		return nil, err
	}
	task.ID = id
	return task, nil
}

func (r *TaskRepository) DeleteTask(id int) error {
	_, err := r.db.Exec(`
        DELETE FROM tasks
        WHERE id = $1
    `, id)
	return err
}
