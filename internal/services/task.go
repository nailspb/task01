package services

import (
	"fmt"
	"log/slog"
	"task01/internal/models"
)

type taskRepository interface {
	GetAllTasks() ([]models.Task, error)
	GetTasksById(id uint) (*models.Task, error)
	CreateTask(task models.Task) error
	UpdateTaskById(id uint, task *models.Task) error
	DeleteTaskByID(id uint) error
}

type taskService struct {
	repo taskRepository
	log  *slog.Logger
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewTaskService(repo taskRepository, log *slog.Logger) *taskService {
	log = log.With(slog.String("source", "TaskService"))
	return &taskService{
		repo: repo,
		log:  log,
	}
}

func (s *taskService) GetAll() ([]models.Task, error) {
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("[TaskService] error on get tasks: %w", err)
	}
	return tasks, nil
}
func (s *taskService) Create(task models.Task) error {
	err := s.repo.CreateTask(task)
	if err != nil {
		return fmt.Errorf("[TaskService] error on create task: %w", err)
	}
	return nil
}
func (s *taskService) UpdateByID(id uint, task *models.Task) error {
	err := s.repo.UpdateTaskById(id, task)
	if err != nil {
		return fmt.Errorf("[TaskService] error on update task by id: %w", err)
	}
	return nil
}
func (s *taskService) DeleteByID(id uint) error {
	err := s.repo.DeleteTaskByID(id)
	if err != nil {
		return fmt.Errorf("[TaskService] repository error on delete tasks by id: %w", err)
	}
	return nil
}

func (s *taskService) Get(id uint) (*models.Task, error) {
	t, err := s.repo.GetTasksById(id)
	if err != nil {
		return nil, fmt.Errorf("[TaskService] repository error on delete tasks by id: %w", err)
	}
	return t, nil
}
