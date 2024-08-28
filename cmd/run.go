package main

import (
	"fmt"
	"log"
	"net/http"

	pg "todo_list_verba/bd"
	"todo_list_verba/config"
	"todo_list_verba/internal/handler"
	"todo_list_verba/internal/repository"
	"todo_list_verba/internal/service"
	"todo_list_verba/migrations"

	"github.com/gorilla/mux"
)

func Run() error {

	c, err := config.New()
	if err != nil {
		fmt.Println(err)
	}
	db, err := pg.ConnectPostgresql(*c)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	if err = migrations.Run(db); err != nil {
		fmt.Println(err)
	}

	repo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(repo)
	taskHandler := handler.NewTaskHandler(taskService)

	r := mux.NewRouter()
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods(http.MethodPost)
	r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.GetTaskByID).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.UpdateTask).Methods(http.MethodPut)
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.DeleteTask).Methods(http.MethodDelete)

	http.Handle("/", r)
	log.Println("Server starting on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
	return err
}
