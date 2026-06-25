package main

import "github.com/gin-gonic/gin"

func AddSessionHandler(store *MemoryStore) gin.HandlerFunc {
	return func(c *gin.Context) {

		var sessions []Session
		err := c.ShouldBindJSON(&sessions)

		if err != nil {
			c.JSON(400, gin.H{"Error": err})
			return
		}

		var date string
		if len(sessions) != 0 {
			date = sessions[0].Date
		} else {
			c.JSON(400, gin.H{"Error of len sessions": len(sessions)})
			return
		}

		for _, v := range sessions {
			if v.Minutes < 0 || v.Minutes == 0 {
				c.JSON(400, gin.H{"Ошибка со временем сессии": "Не валидное количество времени"})
				return
			}
			if v.Date != date {
				c.JSON(400, gin.H{"Ошибка": "дата не совпадает"})
				return
			}
		}

		err = store.AddSession(sessions)
		if err != nil {
			c.JSON(500, gin.H{"Error": err}) //500 - ошибка внутреннего сервера, 400 - ошибка данных внешнего сервера
			return
		}
		sessions, err = store.ListSessionsByDate(date)

		if err != nil {
			c.JSON(500, gin.H{"ошибка": err})
			return
		}

		c.JSON(201, sessions)
	}
}

func ListSessionsByDateHandler(store *MemoryStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var date string
		date = c.Request.URL.Query().Get("date")
		if date == "" {
			c.JSON(400, gin.H{"Не корректные запрос": "Введите дату"})
			return
		}
		sessions, err := store.ListSessionsByDate(date)
		if err != nil {
			c.JSON(500, gin.H{"Error": err})
			return
		}

		c.JSON(200, sessions)
	}
}
