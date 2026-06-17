package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

var UserURL = map[string]string{}

type Url struct {
	URL string `json:"url"`
}

func main() {

	router := gin.Default()

	database, _ := initDB()
	defer db.Close()


	router.POST("/shorten", ShortenUrl)
	router.GET("/hello", /////)
	router.GET("/:shortcode", //// )
	router.Run(":8080")
	fmt.Println("Запустили сервер на :8080")
}
