package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/ping", ping)
	fmt.Println("Сервер запущен на :8090")
	http.ListenAndServe(":8090", nil)
}
