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

	auth := router.Group("/")

	router.POST("/register", Register(conn))
	router.POST("/login", Login(conn))

	auth.Use(JWTMiddleware())
	{
		auth.POST("/tasks", CreateTask(conn))
		auth.GET("/tasks", GetTasks(conn))
		auth.GET("/tasks/:id", GetTask(conn))
		auth.PUT("/tasks/:id", UpdateTask(conn))
		auth.DELETE("/tasks/:id", DeleteTask(conn))
	}
	router.Run(":8080")
}
