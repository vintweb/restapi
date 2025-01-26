package tasksService

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID uint `json:"id"`
	// наш сервер будет ожидать json с полем text
	Task string `json:"task"`
	// в Go используем CamelCase, в JSON - snake
	IsDone bool `json:"is_done"`
}

// Примерная структура для ответа на удаление
type DeleteTasksIdResponseObject struct {
	Message string `json:"message"`
}
