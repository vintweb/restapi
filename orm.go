package main

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	// наш сервер будет ошидать json с полем text
	Task string `json:"task"`
	// в Go используем CamelCase, в JSON - snake
	IsDone bool `json:"is_done"`
}
