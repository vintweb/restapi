package taskService

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	// наш сервер будет ожидать json с полем text
	Task string `json:"task"`
	// в Go используем CamelCase, в JSON - snake
	IsDone bool `json:"is_done"`
}
