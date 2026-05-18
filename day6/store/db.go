package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:pass@localhost:5432/postgres") // подключаемся к таблице
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	fmt.Println("Connected!")

	_, err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE
		)
	`) // создаем таблицу

	if err != nil {
		fmt.Fprintf(os.Stderr, "Create table failed: v%\n", err)
		os.Exit(1)
	}

	fmt.Println("Table created")

	err = CreateUser(conn, "Kirill", "kirill@test.com") // создаём пользователя
	if err != nil {
		fmt.Println("Create user error: ", err)
	} else {
		fmt.Println("User created!")
	}

	name, email, err := GetUser(conn, 1) // выбираем пользователя
	if err != nil {
		fmt.Println("GetUser error: ", err)
	} else {
		fmt.Printf("Got user: %s, %s\n", name, email)
	}

	// обновляем
	err = UpdateUser(conn, 1, "Kirill updated", "kirill_new@test.com")
	if err != nil {
		fmt.Println("Update error: ", err)
	} else {
		fmt.Println("User updated!")
	}

	// удаление
	err = DeleteUser(conn, 1)
	if err != nil {
		fmt.Println("DeleteUser err: ", err)
	} else {
		fmt.Println("User deleted!")
	}
}

func CreateUser(conn *pgx.Conn, name, email string) error {
	_, err := conn.Exec(context.Background(),
		"INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
	return err
}

func GetUser(conn *pgx.Conn, id int) (string, string, error) {
	var name, email string
	err := conn.QueryRow(context.Background(),
		"SELECT name, email FROM users WHERE id=$1", id).Scan(&name, &email)
	return name, email, err
}

func UpdateUser(conn *pgx.Conn, id int, name, email string) error {
	_, err := conn.Exec(context.Background(),
		"UPDATE users SET name=$1, email=$2 WHERE id=$3", name, email, id)
	return err
}

func DeleteUser(conn *pgx.Conn, id int) error {
	_, err := conn.Exec(context.Background(),
		"DELETE FROM users WHERE id=$1", id)
	return err
}
