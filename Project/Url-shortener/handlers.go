package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShortenUrl(c *gin.Context) {
	var req Url
	c.ShouldBindJSON(&req)
	code := strconv.Itoa(rand.Int())
	UserURL[code] = req.URL
	c.JSON(200, gin.H{"short": code})
	fmt.Println("Получил URL:", req.URL)
}

func RedirectHandler(c *gin.Context) {
	code := c.Param("shortcode")
	originalURL, ok := UserURL[code]
	if ok != true {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.Redirect(302, originalURL)
}

func SaveURL(c *context.Background)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello!"})
}
