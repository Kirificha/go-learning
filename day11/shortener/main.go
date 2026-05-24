package main

import (
	"crypto/rand"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.POST("/shorten")
	route.GET("/:code")

}

func Shorten(c *gin.Context) {
	err := c.ShouldBindJSON("url")
	if err != nil {
		c.JSON(400, gin.H{"error": "не получилось найти ссылку"})
		return
	}

	c.JSON(200, gin.H{"message": "Got url"})
	max := 20
	code, err := rand.Int(io.Reader, *max)
	return code, err
}

func Get(c *gin.Context) {

}
