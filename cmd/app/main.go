package main

import (
	"net/http"
	"task01/internal/http/handlers"
	"task01/internal/services"
	"task01/internal/storage/postgres"
	"task01/pkg/prettylogger"
)

func main() {
	mux := http.NewServeMux()

	log := prettylogger.New("local")
	log.Info("Start service")
	storage := postgres.New()
	taskService := services.NewTaskService(storage, log)

	mux.Handle("GET /task", handlers.GetTaskHandler(taskService, log))
	mux.Handle("POST /task", handlers.PostTaskHandler(taskService, log))
	mux.Handle("PATCH /task/{id}", handlers.PatchTaskHandler(taskService, log))
	mux.Handle("DELETE /task/{id}", handlers.DeleteTaskHandler(taskService, log))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
