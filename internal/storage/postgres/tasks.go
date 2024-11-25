package postgres

import (
	"errors"
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

func (s *storage) CreateTask(task models.Task) (*models.Task, error) {
	res := s.db.Create(task)
	if res.Error != nil {
		return nil, res.Error
	}
	return &task, nil
}

func (s *storage) GetAllTasks() ([]models.Task, error) {
	messages := make([]models.Task, 0)
	res := s.db.Find(&messages)
	if res.Error != nil {
		return nil, res.Error
	}
	return messages, nil
}

func (s *storage) UpdateTaskById(id uint, task *models.Task) (bool, error) {
	tx := s.db.Model(&models.Task{}).Where("id = ?", id).Updates(task)
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (s *storage) DeleteTaskByID(id uint) (bool, error) {
	tx := s.db.Delete(&models.Task{}, id)
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (s *storage) GetTasksById(id uint) (*models.Task, error) {
	var t models.Task
	if err := s.db.First(&t, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}
