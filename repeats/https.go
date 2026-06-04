package main

import (
	"fmt"
	"net/http"
)

func h() {
	http.HandleFunc("/solo", Handler)

	err := http.ListenAndServe(":7788", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Сам сделал"))
}
