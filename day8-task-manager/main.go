package main

import (
	"context"

	"github.com/gin-gonic/gin"
)

func main() {
	conn := database()
	defer conn.Close(context.Background())
	createTable(conn)

	router := gin.Default()
	router.Use(LoggerMiddleware())

	router.POST("/tasks", CreateTask(conn))
	router.GET("/tasks", GetTasks(conn))
	router.GET("/tasks/:id", GetTask(conn))
	router.PUT("/tasks/:id", UpdateTask(conn))
	router.DELETE("/tasks/:id", DeleteTask(conn))

	router.Run(":8080")
}
