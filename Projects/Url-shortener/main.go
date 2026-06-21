package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var UserURL = map[string]string{}

type Url struct {
	URL string `json:"url"`
}

func main() {

	router := gin.Default()

	pool, err := initDB()
	db = pool
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	router.POST("/shorten", ShortenUrl)
	router.GET("/hello", Hello)
	router.GET("/:shortcode", RedirectHandler)
	router.Run(":8080")
	fmt.Println("Запустили сервер на :8080")
}
