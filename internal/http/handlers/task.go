package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"task01/internal/models"
)

type Storage interface {
	GetTasks() ([]models.Message, error)
	AddTask(task models.Message) error
	UpdateTask(task models.Message) error
	DeleteTask(id uint) error
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
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(serialized)
		if err != nil {
			fmt.Printf("Error on write response %v\n", err)
			return
		}
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

func PatchTaskHandler(storage Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		var task models.Message
		err := dec.Decode(&task)
		if err != nil || task.ID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = storage.UpdateTask(task); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteTaskHandler(storage Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = storage.DeleteTask(uint(id)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
