package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"task01/internal/models"
)

type taskService interface {
	GetAll() ([]models.Message, error)
	Create(task models.Message) error
	UpdateByID(id uint, task *models.Message) error
	DeleteByID(id uint) error
}

func GetTaskHandler(service taskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := service.GetAll()
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

func PostTaskHandler(service taskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		var task models.Message
		err := dec.Decode(&task)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = service.Create(task); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func PatchTaskHandler(service taskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		dec := json.NewDecoder(r.Body)
		var task models.Message
		err = dec.Decode(&task)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = service.UpdateByID(uint(id), &task); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteTaskHandler(service taskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = service.DeleteByID(uint(id)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
