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

	router.POST("/shorten", ShortenUrl)
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello!"})
	})
	router.GET("/:shortcode", func(c *gin.Context) {
		code := c.Param("shortcode")
		originalURL := UserURL[code]
		c.Redirect(302, originalURL)
	})
	router.Run(":8080")
	fmt.Println("Запустили сервер на :8080")
}

func ShortenUrl(c *gin.Context) {
	var req Url
	c.ShouldBindJSON(&req)
	code := strconv.Itoa(rand.Int())
	UserURL[code] = req.URL
	c.JSON(200, gin.H{"short": code})
	fmt.Println("Получил URL:", req.URL)
}
