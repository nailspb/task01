package main

import (
	"net/http"
	"task01/internal/http/handlers"
	"task01/internal/services"
	"task01/internal/storage/postgres"
)

func main() {
	mux := http.NewServeMux()

	storage := postgres.New()
	taskService := services.NewTaskService(storage)

	mux.Handle("GET /task", handlers.GetTaskHandler(taskService))
	mux.Handle("POST /task", handlers.PostTaskHandler(taskService))
	mux.Handle("PATCH /task/{id}", handlers.PatchTaskHandler(taskService))
	mux.Handle("DELETE /task/{id}", handlers.DeleteTaskHandler(taskService))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
