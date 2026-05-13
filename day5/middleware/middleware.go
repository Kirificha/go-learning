package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // 1. время до запроса

		c.Next() // 2. передаём управление дальше

		duration := time.Since(start) // 3. вычисляем длительность
		method := c.Request.Method    // 4. метод (GET, POST...)
		path := c.Request.URL.Path    // 5. путь (/hello, /ping)
		status := c.Writer.Status()   // 6. статус ответа (200, 404...)

		fmt.Printf("[%s] %s -> %d (%s)\n", method, path, status, duration.String())
	}
}

func main() {
	router := gin.Default()

	router.Use(LoggerMiddleware()) // регистрируем middleware

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello"})
	})

	fmt.Println("Сервер запущен на :8090")
	router.Run(":8090") // явно указываем порт
}
