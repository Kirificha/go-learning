package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", HandleFunc)

	fmt.Println("Запустили сервер на :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error start:", err)
	}
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}
