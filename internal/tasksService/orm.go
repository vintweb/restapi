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
	IsDone bool  `json:"is_done"`
	User   *User `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"` // Связь с пользователем через указатель
	UserID uint  `json:"user_id"`
}

type User struct {
	UserID uint   `json:"id"`
	Email  string `json:"email"`
}

// Примерная структура для ответа на удаление
type DeleteTasksIdResponseObject struct {
	Message string `json:"message"`
}
