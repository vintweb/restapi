package usersService

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint       `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	DeletedAt *time.Time `json:"delete_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	Tasks     []Task     `json:"tasks"` // Связь с задачами через указатель
}

type Task struct {
	ID     uint   `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id"`
}
