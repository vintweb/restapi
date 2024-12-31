package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var messages []Message // Слайс для хранения записей

	// Получаем все сообщения из БД
	if result := DB.Find(&messages); result.Error != nil {
		http.Error(w, "Ошибка получения сообщений", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок ответа и отправляем JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message Message

	// Декодируем JSON из тела запроса в структуру Messega
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	// Сохраняем сообщение в БД
	if result := DB.Create(&message); result.Error != nil {
		http.Error(w, "Ошибка создания сообщения", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок ответа и отправляем JSON созданной сущности
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // HTTP 201 Created
	json.NewEncoder(w).Encode(message)
}

func main() {
	// вызываем метод InitDB() из файла db.go
	InitDB()

	//Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	// Настройка маршрутов
	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessage).Methods("GET")

	// Запуск сервера
	fmt.Println("The server is running on port 8080...")
	http.ListenAndServe(":8080", router)
}
