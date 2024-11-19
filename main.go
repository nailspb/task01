package main

import (
	"net/http"
	"task01/internal/http/handlers"
	"task01/internal/models"
	"task01/internal/storage/postgres"
)

func main() {
	str := postgres.New()
	err := str.Migrate(&models.Message{})
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("GET /task", handlers.GetTaskHandler(str))
	mux.Handle("POST /task", handlers.PostTaskHandler(str))
	mux.Handle("PATCH /task", handlers.PatchTaskHandler(str))
	mux.Handle("DELETE /task", handlers.DeleteTaskHandler(str))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
