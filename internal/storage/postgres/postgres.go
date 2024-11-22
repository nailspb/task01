package postgres

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"task01/internal/models"
)

type storage struct {
	db *gorm.DB
}

//goland:noinspection GoExportedFuncWithUnexportedType
func New() *storage {
	connection := "host=srv2.spartatn.ru user=user password=super_puper_user_password dbname=tests port=5430 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	//migrations
	/*err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}*/
	return &storage{
		db: DB,
	}
}

func (s *storage) CreateTask(task models.Task) error {
	res := s.db.Create(&models.Task{
		Task:   task.Task,
		IsDone: task.IsDone,
	})
	if res.Error != nil {
		return errors.Join(errors.New("error executing a database request to add a task"), res.Error)
	}
	return nil
}

func (s *storage) GetAllTasks() ([]models.Task, error) {
	messages := make([]models.Task, 0)
	res := s.db.Find(&messages)
	if res.Error != nil {
		return nil, fmt.Errorf("[Storage/postgres] error when getting all the tasks: %w", res.Error)
	}
	return messages, nil
}

func (s *storage) UpdateTaskById(id uint, task *models.Task) error {
	tx := s.db.Model(&models.Task{}).Where("id = ?", id).Updates(models.Task{Task: task.Task, IsDone: task.IsDone})
	if tx.Error != nil {
		return fmt.Errorf("[Storage/postgres] error when updating task: %w", tx.Error)
	}
	if tx.RowsAffected == 0 {
		return errors.New("[Storage/postgres] the task being updated was not found")
	}
	return nil
}

func (s *storage) DeleteTaskByID(id uint) error {
	tx := s.db.Delete(&models.Task{}, id)
	if tx.Error != nil {
		return fmt.Errorf("[Storage/postgres] error when delete task: %w", tx.Error)
	}
	if tx.RowsAffected == 0 {
		return errors.New("[Storage/postgres] task not found")
	}
	return nil
}

func (s *storage) GetTasksById(id uint) (*models.Task, error) {
	var t models.Task
	if err := s.db.First(&t, id).Error; err != nil {
		return nil, fmt.Errorf("[Storage/postgres] error when get task: %w", err)
	}
	return &t, nil
}
