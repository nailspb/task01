package postgres

import (
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
	err = DB.AutoMigrate(&models.Message{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	return &storage{
		db: DB,
	}
}

func (s *storage) CreateTask(task models.Message) error {
	res := s.db.Create(&models.Message{
		Task:   task.Task,
		IsDone: task.IsDone,
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (s *storage) GetAllTasks() ([]models.Message, error) {
	messages := make([]models.Message, 0)
	res := s.db.Find(&messages)
	if res.Error != nil {
		return nil, res.Error
	}
	return messages, nil
}

func (s *storage) UpdateTaskById(id uint, task *models.Message) error {
	tx := s.db.Model(&models.Message{}).Where("id = ?", id).Updates(models.Message{Task: task.Task, IsDone: task.IsDone})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *storage) DeleteTaskByID(id uint) error {
	tx := s.db.Delete(&models.Message{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *storage) GetTasksById(id uint) (*models.Message, error) {
	var t models.Message
	if err := s.db.First(&t, id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}
