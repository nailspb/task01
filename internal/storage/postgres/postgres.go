package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"task01/internal/models"
)

type Storage struct {
	db *gorm.DB
}

func New() *Storage {
	connection := "host=srv2.spartatn.ru user=user password=super_puper_user_password dbname=tests port=5430 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	return &Storage{
		db: DB,
	}
}

func (s *Storage) Migrate(dst ...any) error {
	for _, d := range dst {
		err := s.db.AutoMigrate(&d)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) AddTask(task models.Message) error {
	res := s.db.Create(&task)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (s *Storage) GetTasks() ([]models.Message, error) {
	messages := make([]models.Message, 0)
	res := s.db.Find(&messages)
	if res.Error != nil {
		return nil, res.Error
	}
	return messages, nil
}

func (s *Storage) UpdateTask(task models.Message) error {
	var t models.Message
	if err := s.db.First(&t, task.ID).Error; err != nil {
		return err
	}
	if task.Task != "" {
		t.Task = task.Task
	}
	t.IsDone = task.IsDone
	if err := s.db.Save(&t).Error; err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteTask(id uint) error {
	if err := s.db.Delete(&models.Message{}, id).Error; err != nil {
		return err
	}
	return nil
}
