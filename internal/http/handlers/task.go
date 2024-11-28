package handlers

import (
	"context"
	"task01/internal/models"
	"task01/internal/web/tasks"
)

type taskService interface {
	GetAll() ([]models.Task, error)
	GetAllByUser(id uint) ([]models.Task, error)
	Get(id uint) (*models.Task, error)
	Create(task models.Task) (*models.Task, error)
	UpdateByID(id uint, task *models.Task) (bool, error)
	DeleteByID(id uint) (bool, error)
}

type tasksHandlers struct {
	service taskService
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewTasksHandler(service taskService) *tasksHandlers {
	return &tasksHandlers{
		service,
	}
}

func (t *tasksHandlers) GetUserTasksId(_ context.Context, request tasks.GetUserTasksIdRequestObject) (tasks.GetUserTasksIdResponseObject, error) {
	allTasks, err := t.service.GetAllByUser(request.Id)
	if err != nil {
		return nil, err
	}
	response := make(tasks.GetUserTasksId200JSONResponse, len(allTasks))
	for i, task := range allTasks {
		response[i] = tasks.Task{
			Created: &task.CreatedAt,
			Id:      &task.ID,
			IsDone:  &task.IsDone,
			Task:    &task.Task,
			Updated: &task.UpdatedAt,
			UserId:  &task.UserId,
		}
	}
	return response, nil
}

func (t *tasksHandlers) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := t.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := make(tasks.GetTasks200JSONResponse, len(allTasks))
	for i, task := range allTasks {
		response[i] = tasks.Task{
			Created: &task.CreatedAt,
			Id:      &task.ID,
			IsDone:  &task.IsDone,
			Task:    &task.Task,
			Updated: &task.UpdatedAt,
			UserId:  &task.UserId,
		}
	}
	return response, nil
}

func (t *tasksHandlers) PostTasks(_ context.Context, r tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	task := models.Task{}
	if r.Body.Task != nil {
		task.Task = *r.Body.Task
	}
	if r.Body.IsDone != nil {
		task.IsDone = *r.Body.IsDone
	}
	task.UserId = r.Body.UserId
	newTask, err := t.service.Create(task)
	if err != nil {
		return nil, err
	}
	return tasks.PostTasks201JSONResponse{
		Created: &newTask.CreatedAt,
		Id:      &newTask.ID,
		IsDone:  &newTask.IsDone,
		Task:    &newTask.Task,
		Updated: &newTask.UpdatedAt,
		UserId:  &newTask.UserId,
	}, nil

}

func (t *tasksHandlers) PatchTasksId(_ context.Context, r tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	task := models.Task{}
	if r.Body.Task != nil {
		task.Task = *r.Body.Task
	}
	if r.Body.IsDone != nil {
		task.IsDone = *r.Body.IsDone
	}
	task.UserId = r.Body.UserId
	done, err := t.service.UpdateByID(r.Id, &task)
	if err != nil {
		return nil, err
	}
	if done {
		return tasks.PatchTasksId200Response{}, nil
	}
	return tasks.PatchTasksId404Response{}, nil
}

func (t *tasksHandlers) DeleteTasksId(_ context.Context, r tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	done, err := t.service.DeleteByID(r.Id)
	if err != nil {
		return nil, err
	}
	if done {
		return tasks.DeleteTasksId200Response{}, nil
	}
	return tasks.DeleteTasksId404Response{}, nil
}

func (t *tasksHandlers) GetTasksId(_ context.Context, r tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	task, err := t.service.Get(r.Id)
	if err != nil {
		return nil, err
	}
	if task != nil {
		return tasks.GetTasksId200JSONResponse{
			Created: &task.CreatedAt,
			Id:      &task.ID,
			IsDone:  &task.IsDone,
			Task:    &task.Task,
			Updated: &task.UpdatedAt,
			UserId:  &task.UserId,
		}, nil
	}
	return tasks.GetTasksId404Response{}, nil
}
