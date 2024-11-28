package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserId uint   `json:"user_id"`
}

type User struct {
	gorm.Model
	Email    string
	Password string
	Tasks    []Task
}
