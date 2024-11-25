package postgres

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"task01/internal/models"
)

func (s *storage) CreateUser(d *models.User) error {
	res := s.db.Create(d)
	if res.Error != nil {
		return errors.Join(errors.New("error executing a database request to add a user"), res.Error)
	}
	return nil
}

func (s *storage) GetAllUsers() ([]models.User, error) {
	users := make([]models.User, 0)
	res := s.db.Find(&users)
	if res.Error != nil {
		return nil, fmt.Errorf("[Storage/postgres] error when getting all the users: %w", res.Error)
	}
	return users, nil
}

func (s *storage) UpdateUserById(id uint, d *models.User) (bool, error) {
	tx := s.db.Model(&models.User{}).Where("id = ?", id).Updates(d)
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (s *storage) DeleteUserByID(id uint) (bool, error) {
	tx := s.db.Delete(&models.User{}, id)
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (s *storage) GetUserById(id uint) (*models.User, error) {
	var t models.User
	if err := s.db.First(&t, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("[Storage/postgres] error when get user: %w", err)
	}
	return &t, nil
}
