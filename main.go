package main

import (
	"encoding/json"
	"net/http"
	"task01/internal/models"
	"task01/internal/storage/postgres"
)

type Storage interface {
	GetTasks() ([]models.Message, error)
	AddTask(task models.Message) error
}

func GetTaskHandler(storage Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := storage.GetTasks()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		serialized, err := json.Marshal(tasks)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(serialized)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func PostTaskHandler(storage Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		var task models.Message
		err := dec.Decode(&task)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = storage.AddTask(task); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	str := postgres.New()
	err := str.Migrate(&models.Message{})
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("GET /task", GetTaskHandler(str))
	mux.Handle("POST /task", PostTaskHandler(str))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
