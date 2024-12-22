package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var counter int

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Counter равен", strconv.Itoa(counter))
	} else {
		fmt.Fprintln(w, "Поддерживается только метод GET")
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		counter++
		fmt.Fprintln(w, "Counter увеличен на 1")
	} else {
		fmt.Fprintln(w, "Поддерживается только метод POST")
	}
}

func main() {
	// 	http.HandleFunc("/hello", HelloHandler)
	// 	http.ListenAndServe(":8080", nil)

	router := mux.NewRouter()
	// Наше приложение будет слушать запросы на localhost:8080/api/hello
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	http.ListenAndServe(":8080", router)

	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/post", PostHandler)

	http.ListenAndServe(":8080", nil)
}
