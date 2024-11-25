package services

import (
	"log/slog"
	"task01/internal/models"
)

type userRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (*models.User, error)
	CreateUser(task *models.User) error
	UpdateUserById(id uint, task *models.User) (bool, error)
	DeleteUserByID(id uint) (bool, error)
}

type userService struct {
	repo userRepository
	log  *slog.Logger
}

func (u userService) GetAll() ([]models.User, error) {
	return u.repo.GetAllUsers()
}

func (u userService) Get(id uint) (*models.User, error) {
	return u.repo.GetUserById(id)
}

func (u userService) Create(user *models.User) error {
	return u.repo.CreateUser(user)
}

func (u userService) UpdateByID(id uint, task *models.User) (bool, error) {
	return u.repo.UpdateUserById(id, task)
}

func (u userService) DeleteByID(id uint) (bool, error) {
	return u.repo.DeleteUserByID(id)
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewUserService(repo userRepository, log *slog.Logger) *userService {
	log = log.With(slog.String("source", "TaskService"))
	return &userService{
		repo: repo,
		log:  log,
	}
}
