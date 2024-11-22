package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"task01/internal/models"
)

type taskService interface {
	GetAll() ([]models.Task, error)
	Create(task models.Task) error
	UpdateByID(id uint, task *models.Task) error
	DeleteByID(id uint) error
}

func GetTaskHandler(service taskService, log *slog.Logger) http.HandlerFunc {
	log = log.With(slog.String("source", "GetTaskHandler"))
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := service.GetAll()
		if err != nil {
			log.Error("service error", slog.String("error", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		serialized, err := json.Marshal(tasks)
		if err != nil {
			log.Error("error on stringify data to json", slog.String("error", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(serialized)
		if err != nil {
			log.Error("error on write response", slog.String("error", err.Error()))
			return
		}

	}
}

func PostTaskHandler(service taskService, log *slog.Logger) http.HandlerFunc {
	log = log.With(slog.String("source", "PostTaskHandler"))
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error("error on read body", slog.String("error", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		var task models.Task
		err = json.Unmarshal(body, &task)
		if err != nil {
			log.Warn("bad request from user",
				slog.String("body", string(body)),
				slog.String("error", err.Error()),
			)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = service.Create(task); err != nil {
			log.Warn("the service was unable to create task record", slog.String("error", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func PatchTaskHandler(service taskService, log *slog.Logger) http.HandlerFunc {
	log = log.With(slog.String("source", "PatchTaskHandler"))
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			log.Warn("bad id parameter", slog.String("id", r.PathValue("id")))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error("error on read body", slog.String("error", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		var task models.Task
		err = json.Unmarshal(body, &task)
		if err != nil {
			log.Warn("bad request from user",
				slog.String("body", strings.Replace(string(body), "\r\n", "", -1)),
				slog.String("error", err.Error()),
			)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = service.UpdateByID(uint(id), &task); err != nil {
			log.Warn("the service was unable to update the data", slog.String("error", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteTaskHandler(service taskService, log *slog.Logger) http.HandlerFunc {
	log = log.With(slog.String("source", "PatchTaskHandler"))
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Warn("bad id parameter", slog.String("id", r.PathValue("id")))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = service.DeleteByID(uint(id)); err != nil {
			log.Warn("the service was unable to delete task", slog.String("error", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
