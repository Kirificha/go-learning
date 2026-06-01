package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

// handler.go - просто функции без роутера
func Register(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "не удалось зарегистрировать"})
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		_, err = conn.Exec(context.Background(),
			"INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, hash)

		if err != nil {
			fmt.Println("Register: ", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(201, gin.H{"message": "Регистрация успешна"})
		}
	}
}

func Login(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "не удалось зарегистрировать"})
		}

		inputPassword := user.Password

		err := conn.QueryRow(context.Background(),
			"SELECT id, email, password FROM users WHERE email=$1", user.Email).Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			c.JSON(500, gin.H{"error": "пользователь не найден"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputPassword))
		if err != nil {
			c.JSON(301, gin.H{"error": "неверный пароль"})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"exp":     time.Now().Add(24 * time.Hour).Unix(),
		})

		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			c.JSON(500, gin.H{"error": "ошибка токена"})
			return
		}
		c.JSON(200, gin.H{"Token": tokenString})
	}
}

func CreateTask(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(400, gin.H{"error": "неверный формат"})
			return
		}
		_, err := conn.Exec(context.Background(),
			"INSERT INTO tasks (title) VALUES ($1)", task.Title)

		if err != nil {
			fmt.Println("Create task error: ", err)
		} else {
			c.JSON(201, gin.H{"message": "Задача создана"})
		}

	}
}

func GetTasks(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := conn.Query(context.Background(),
			"SELECT id, title, done, created_at FROM tasks")
		if err != nil {
			c.JSON(500, gin.H{"error": "Ошибка из БД"})
			return
		}
		defer rows.Close()
		var tasks []Task
		for rows.Next() {
			var t Task
			err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt)
			if err != nil {
				c.JSON(500, gin.H{"error": "ошибка чтения"})
				return
			}
			tasks = append(tasks, t)
		}

		c.JSON(200, tasks)
	}

}

func GetTask(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var task Task
		row := conn.QueryRow(context.Background(),
			"SELECT * FROM tasks WHERE id=$1", id)
		err := row.Scan(&task.ID, &task.Title, &task.Done, &task.CreatedAt)

		if err != nil {
			c.JSON(500, gin.H{"error": "Ошибка выбора задачи"})
			return
		}
		c.JSON(200, gin.H{"message": "Get task"})
	}

}

func UpdateTask(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var task Task
		c.ShouldBindJSON(&task)
		_, err := conn.Exec(context.Background(),
			"UPDATE tasks SET title=$1, done=$2 WHERE id=$3", task.Title, task.Done, id)

		if err != nil {
			c.JSON(200, gin.H{"Ошибка обновления задачи": err})
			return
		}

		c.JSON(200, gin.H{"message": "Update task"})
	}

}

func DeleteTask(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		_, err := conn.Exec(context.Background(),
			"DELETE FROM tasks WHERE id=$1", id)

		if err != nil {
			c.JSON(500, gin.H{"Ошибка удаления задачи": err})
			return
		}
		c.JSON(200, gin.H{"message": "Delete task"})
	}

}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()

		fmt.Printf("[%s] %s -> %d (%s)\n", method, path, status, duration.String())
	}
}
