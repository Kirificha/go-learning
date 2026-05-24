package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()

	route.GET("/ping", handler)

	route.Run(":8080")
}

func handler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
