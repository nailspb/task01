package services

import (
	"task01/internal/models"
)

type taskRepository interface {
	GetAllTasks() ([]models.Message, error)
	GetTasksById(id uint) (*models.Message, error)
	CreateTask(task models.Message) error
	UpdateTaskById(id uint, task *models.Message) error
	DeleteTaskByID(id uint) error
}

type taskService struct {
	repo taskRepository
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewTaskService(repo taskRepository) *taskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) GetAll() ([]models.Message, error) {
	return s.repo.GetAllTasks()
}
func (s *taskService) Create(task models.Message) error {
	return s.repo.CreateTask(task)
}
func (s *taskService) UpdateByID(id uint, task *models.Message) error {
	return s.repo.UpdateTaskById(id, task)
}
func (s *taskService) DeleteByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
