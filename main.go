package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

func GetHandler(w http.ResponseWriter, r *http.Request) {
	// GET handler для возврата приветствия с задачей

	if task == "" {
		fmt.Fprintf(w, "Hello, no task set!")
	} else {
		fmt.Fprintf(w, "Hello, %s!", task)
	}
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var messages []Message // Слайс для хранения записей

	result := DB.Find(&messages)
	if result.Error != nil {
		fmt.Println(w, "Error fetching messages:", http.StatusInternalServerError)
		return
	}

	// Печатаем все меседжи
	for _, message := range messages {
		DB.Find(&message)
		fmt.Fprintf(w, "Task: %s, IsDone: %t\n", message.Task, message.IsDone)
	}
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	messages := []Message{
		{Task: "Новая задача #1", IsDone: true},
		{Task: "Новая задача #2", IsDone: true},
		{Task: "Новая задача #3", IsDone: false},
	}

	for _, message := range messages {
		DB.Create(&message)
	}

	fmt.Fprintln(w, "Сообщения успешно добавлены!")
}

func main() {
	// вызываем метод InitDB() из файла db.go
	InitDB()

	//Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessage).Methods("GET")

	// Запуск сервера
	fmt.Println("The server is running on port 8080...")
	http.ListenAndServe(":8080", router)
}
