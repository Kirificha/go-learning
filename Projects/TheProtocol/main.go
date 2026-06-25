package main

import "github.com/gin-gonic/gin"

func main() {
	store := &MemoryStore{}
	router := gin.Default()
	router.POST("/sessions", AddSessionHandler(store))
	router.GET("/sessions", ListSessionsByDateHandler(store))

	router.Run(":8080")

}
