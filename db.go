package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// переменая для работы с БД
var DB *gorm.DB

func InitDB() {
	// в dns вводим данные указанные при создании контейнера
	dns := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed  to connect to database ", err)
	}
}
