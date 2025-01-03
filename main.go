package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	// Декодируем JSON из тела запроса в структуру Message
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

// PATHC: Обновление задачи
func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	// Получение ID из параметров
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Не корректный ID", http.StatusBadRequest)
		return
	}

	//Поиск задачи по ID
	var message Message
	if result := DB.First(&message, id); result.Error != nil {
		http.Error(w, "Задача не найдена", http.StatusBadRequest)
		return
	}

	// Декодируем JSON из запроса
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Неверый формат JSON", http.StatusBadRequest)
		return
	}

	// Обновляем запись
	if result := DB.Model(&message).Updates(updates); result.Error != nil {
		http.Error(w, "Ошибка обновления задачи", http.StatusInternalServerError)
		return
	}

	// Возвращаем обновленную задачу
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// DELETE: Удаление задачи
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	// Получение ID из каталога
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	// DELETE: Удаление задачи по ID
	if result := DB.Delete(&Message{}, id); result.Error != nil {
		http.Error(w, "Ошибка удаления задачи", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный статус
	w.WriteHeader(http.StatusNoContent) // HTTP 204 No Content
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
	router.HandleFunc("/api/messages/{id}", UpdateMessage).Methods("PATCH")
	router.HandleFunc("/api/messages/{id}", DeleteMessage).Methods("DELETE")

	// Запуск сервера
	fmt.Println("The server is running on port 8080...")
	http.ListenAndServe(":8080", router)
}
