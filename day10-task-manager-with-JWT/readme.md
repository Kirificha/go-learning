# Task Manager API

Что делает проект — базовый CRUD, создаёт данные, выбирает их, изменяет, удаляет. Задачник где можно добавлять/удалять/изменять задачи/выбирать задачу

## Стек
- Go + Gin
- PostgreSQL
- Docker

## Роуты
POST /tasks — создать задачу
GET /tasks — выбрать все задачи
GET /tasks/:id - выбрать задачу
PUT /tasks/:id - обновить задачу
DELETE /tasks/:id - удалить

## Как запустить

Через Docker:
docker-compose up --build

Локально (нужен запущенный PostgreSQL):
go run .

## Примеры запросов
curl -X PUT http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d "{\"title\":\"Обновлённая задача\",\"done\":true}"
{"message":"Update task"}

curl -X DELETE http://localhost:8080/tasks/1
{"message":"Delete task"}