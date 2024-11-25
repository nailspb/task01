package services

import (
	"log/slog"
	"task01/internal/models"
)

type taskRepository interface {
	GetAllTasks() ([]models.Task, error)
	GetTasksById(id uint) (*models.Task, error)
	CreateTask(task models.Task) (*models.Task, error)
	UpdateTaskById(id uint, task *models.Task) (bool, error)
	DeleteTaskByID(id uint) (bool, error)
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
	return s.repo.GetAllTasks()
}
func (s *taskService) Create(task models.Task) (*models.Task, error) {
	return s.repo.CreateTask(task)
}
func (s *taskService) UpdateByID(id uint, task *models.Task) (bool, error) {
	return s.repo.UpdateTaskById(id, task)
}
func (s *taskService) DeleteByID(id uint) (bool, error) {
	return s.repo.DeleteTaskByID(id)
}

func (s *taskService) Get(id uint) (*models.Task, error) {
	return s.repo.GetTasksById(id)
}
