package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// var counter int
var task string

// Структура для обработки JSON из POST-запроса
type RequestBody struct {
	Message string `json:"message"`
}

// func HelloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello World")
// }

func GetHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method == http.MethodGet {
	// 	fmt.Fprintln(w, "Counter равен", strconv.Itoa(counter))
	// } else {
	// 	fmt.Fprintln(w, "Поддерживается только метод GET")
	// }

	// GET handler для возврата приветствия с задачей

	if task == "" {
		fmt.Fprintf(w, "Hello, no task set!")
	} else {
		fmt.Fprintf(w, "Hello, %s!", task)
	}
}

// func PostHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		counter++
// 		fmt.Fprintln(w, "Counter увеличен на 1")
// 	} else {
// 		fmt.Fprintln(w, "Поддерживается только метод POST")
// 	}
// }

// POST handler для записи значения в глобальную переменную
func TaskHundler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	var requestBody RequestBody

	// Декодируем JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Записываем значение из JSON в глобальную переменную
	task = requestBody.Message
	fmt.Fprintf(w, "Задача '%s' сохранена!", task)
}

func main() {
	// 	http.HandleFunc("/hello", HelloHandler)
	// 	http.ListenAndServe(":8080", nil)

	// router := mux.NewRouter()
	// Наше приложение будет слушать запросы на localhost:8080/api/hello
	// router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	// http.ListenAndServe(":8080", router)

	// Обработчик GET-запроса
	http.HandleFunc("/get", GetHandler)
	// http.HandleFunc("/post", PostHandler)

	// Обработчик POST-запроса Task
	http.HandleFunc("/task", TaskHundler)

	fmt.Println("The server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
