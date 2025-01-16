package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"restapi/internal/database" // Импортируем пакет нашей БД
	"restapi/internal/handlers"
	"restapi/internal/taskService"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	// Функция может вернуть ошибку, если сервер не может быть запущен.
	// Поэтому важно использовать log.Fatal, чтобы вывести ошибку и завершить программу,
	// если запуск не удался.
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Финальное задание - переписать main.go файл,
// убедиться что все работает, залить изменения
// в репозиторий и прислать мне скриншоты / видео с результатом проделанной работы
