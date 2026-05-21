package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

func database() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:pass@localhost:5432/postgres")
	if err != nil {
		fmt.Println("Connection database error: ", err)
		os.Exit(1)
	}

	fmt.Println("Database connected!")

	return conn
}

func createTable(conn *pgx.Conn) {
	_, err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		done BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT NOW()
		)
	`)

	if err != nil {
		fmt.Println("Create-table error: ", err)
		os.Exit(1)
	} else {
		fmt.Println("Table ready!")
	}
}
