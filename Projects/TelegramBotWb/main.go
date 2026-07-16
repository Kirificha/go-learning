package main

import (
	"io"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	// Tell telegram that we've handled previous responses
	updateConfig := tgbotapi.NewUpdate(0)
	// Waiting up for 30 seconds between each request
	updateConfig.Timeout = 30
	// Poiling Telegram for updates
	updates := bot.GetUpdatesChan(updateConfig)

	// Going trough each update
	for update := range updates {
		// Tell telegram that we need only type of messages now, discard others from history. But im dont understand that method
		// Cause why we used if message == nul = continue. Okay, later im gonna understand this
		if update.Message.Document == nil {
			continue
		}
		// User document (XSLX/CSV)
		switch update.Message.Document.MimeType {
		case "text/csv":
		case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, загрузите файл в формате CSV или Excel.")
			bot.Send(msg)
			continue
		}
		docID := update.Message.Document.FileID
		file, err := bot.GetFile(tgbotapi.FileConfig{FileID: docID})
		if err != nil {
			panic(err)
		}
		log.Printf("File size: %v\nFile path: %v", file.FileSize, file.FilePath)
		downloadURL, err := bot.GetFileDirectURL(file.FileID)
		if err != nil {
			panic(err)
		}

		downloadedFile, err := http.Get(downloadURL)
		if err != nil {
			panic(err)
		}
		defer downloadedFile.Body.Close()

		data, err := io.ReadAll(downloadedFile.Body)
		if err != nil {
			panic(err)
		}

		// Goten new message, using txt from message and id to create a second message to client
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// Reply to previous message of client
		msg.ReplyToMessageID = update.Message.MessageID

		// Okay, we sended our message, dont care about message, so well discard it
		if _, err := bot.Send(msg); err != nil {
			// Notion. Цитата: // Note that panics are a bad way to handle errors. Telegram can
			// have service outages or network errors, you should retry sending
			// messages or more gracefully handle failures.
			panic(err)
		}
	}

}
