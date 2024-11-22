package handlers

import (
	"context"
	"log/slog"
	"task01/internal/models"
	"task01/internal/web/tasks"
)

type taskService interface {
	GetAll() ([]models.Task, error)
	Get(id uint) (*models.Task, error)
	Create(task models.Task) error
	UpdateByID(id uint, task *models.Task) error
	DeleteByID(id uint) error
}

type tasksHandlers struct {
	service taskService
	log     *slog.Logger
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewTasksHandler(service taskService, log *slog.Logger) *tasksHandlers {
	log = log.With(slog.String("handler", "task"))
	return &tasksHandlers{
		service, log,
	}
}

func (t *tasksHandlers) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	logger := t.log.With(slog.String("method", "GET"))
	allTasks, err := t.service.GetAll()
	if err != nil {
		logger.Error("service error", slog.String("error", err.Error()))
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}
	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (t *tasksHandlers) PostTasks(_ context.Context, r tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	logger := t.log.With(slog.String("method", "POST"))
	if err := t.service.Create(models.Task{
		Task:   *r.Body.Task,
		IsDone: *r.Body.IsDone,
	}); err != nil {
		logger.Warn("the service was unable to create task record", slog.String("error", err.Error()))
		return nil, err
	}
	return tasks.PostTasks201Response{}, nil

}

func (t *tasksHandlers) PatchTasksId(_ context.Context, r tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	logger := t.log.With(slog.String("method", "PATCH"))
	if err := t.service.UpdateByID(r.Id, &models.Task{
		Task:   *r.Body.Task,
		IsDone: *r.Body.IsDone,
	}); err != nil {
		logger.Warn("the service was unable to update the data", slog.String("error", err.Error()))
		return nil, err
	}
	return tasks.PatchTasksId200Response{}, nil
}

func (t *tasksHandlers) DeleteTasksId(_ context.Context, r tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	logger := t.log.With(slog.String("method", "DELETE"))
	if err := t.service.DeleteByID(r.Id); err != nil {
		logger.Warn("the service was unable to delete task", slog.String("error", err.Error()))
		return nil, err
	}
	return tasks.DeleteTasksId200Response{}, nil
}

func (t *tasksHandlers) GetTasksId(_ context.Context, r tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	logger := t.log.With(slog.String("method", "GET"))
	task, err := t.service.Get(r.Id)
	if err != nil {
		logger.Warn("the service was unable to get task by id", slog.String("error", err.Error()))
		return nil, err
	}
	if task != nil {
		return tasks.GetTasksId200JSONResponse{
			Id:     &task.ID,
			IsDone: &task.IsDone,
			Task:   &task.Task,
		}, nil
	}
	return tasks.GetTasksId404Response{}, nil
}
