package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		autHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(autHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "невалидный токен"})
			c.Abort()
			return
		}

		//достать user_id из токена
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"]
		c.Set("user_id", userID) // положить в контекст для хендлеров (обработчиков)
		c.Next()
	}
}
