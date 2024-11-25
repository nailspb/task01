package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task   string `json:"task"`    // Наш сервер будет ожидать json c полем text
	IsDone bool   `json:"is_done"` // В GO используем CamelCase, в Json - snake
}

type User struct {
	gorm.Model
	Email    string
	Password string
}