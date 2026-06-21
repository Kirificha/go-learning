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
	if err := SaveURL(code, req.URL); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"short": code})
	fmt.Println("Получил URL:", req.URL)
}

func RedirectHandler(c *gin.Context) {
	code := c.Param("shortcode")
	originalURL, err := getURL(code)
	if err != nil {
		c.JSON(404, gin.H{"error": "404 Not found"})
		return
	}
	c.Redirect(302, originalURL)
}

var ctx = context.Background()

func SaveURL(shortcode string, originalURL string) error {
	_, err := db.Exec(ctx, "INSERT INTO urls (short_code, original_url) VALUES ($1, $2)", shortcode, originalURL)
	return err
}

func getURL(shortcode string) (string, error) {
	var originalURL string
	err := db.QueryRow(ctx, "SELECT original_url FROM urls WHERE short_code = $1", shortcode).Scan(&originalURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello!"})
}
