package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

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
		// Validation user document (XSLX/CSV)
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

		// Using encoding/csv
		readerCSV := csv.NewReader(bytes.NewReader(data))
		records, err := readerCSV.ReadAll()
		// Заголовки. Распаршу потом. Также надо будет err поменять на log.print че-то там и continue
		headers := []string{
			"Название", "Цена розничная с учётом согласованной скидки",
			"Размер кВВ, %", "Вознаграждение Вайлдбериз (ВВ), без НДС",
			"Услуги по доставке товара покупателю",
		}
		// Находим нужные колонки в первой строке
		columns := make(map[string]int)
		for i, colon := range records[0] {
			for _, header := range headers {
				if colon == header {
					columns[header] = i
				}

			}
		}
		// Проверяем наличие всех колонок
		if len(columns) < 5 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "В вашем файле количество нужных колонок меньше 5")
			bot.Send(msg)
			continue
		}
		// Создаём структуру для расхождений
		type Discperancy struct {
			Name     string
			Expected float64
			Actual   float64
			Diff     float64
		}
		errors := []Discperancy{}

		for i, record := range records[1:] {
			nameOfProduct := record[columns["Название"]]
			price := record[columns["Цена розничная с учётом согласованной скидки"]]
			kVV := record[columns["Размер кВВ, %"]]
			giftWB := record[columns["Вознаграждение Вайлдбериз (ВВ), без НДС"]]

			priceNum, err := strconv.ParseFloat(price, 64)
			if err != nil {
				continue
			}
			kVVNum, err := strconv.ParseFloat(kVV, 64)
			if err != nil {
				continue
			}
			giftWBNum, err := strconv.ParseFloat(giftWB, 64)
			if err != nil {
				continue
			}

			expected := priceNum * kVVNum / 100
			expected = math.Round(expected*100) / 100
			giftWBNum = math.Round(giftWBNum*100) / 100

			if expected != giftWBNum {
				discrep := Discperancy{
					Name: nameOfProduct, Expected: expected, Actual: giftWBNum, Diff: giftWBNum - expected,
				}
				errors = append(errors, discrep)
			}

		}
		// Проверка наличия расхождений
		if len(errors) == 0 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отлично! Расхождений не найдено")
			bot.Send(msg)
			continue
		}
		// Формирование листа расхождений
		sumDiff := 0.0
		var responseList []string
		for _, errWB := range errors {
			sumDiff += errWB.Diff
			response := fmt.Sprintf("- Товар %v: комиссия списана %v₽, должна быть %v₽. Переплата %v₽.", errWB.Name, errWB.Actual, errWB.Expected, errWB.Diff)
			responseList = append(responseList, response)
		}
		// Формирование готового ответа to user
		responseListStr := strings.Join(responseList, "\n")
		userResponse := fmt.Sprintf(`Результат проверки:
					%v

					Всего найдено расхождений: %v
					Общая сумма переплат: %v₽`, responseListStr, len(errors), sumDiff)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, userResponse)
		bot.Send(msg)

		// Создаём inline buttons для проверки всё ли правильно
		markup := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("✅ Всё верно", "ok"),
				tgbotapi.NewInlineKeyboardButtonData("❌ Есть ошибка", "error"),
			),
		)

		msg.ReplyMarkup = markup
		bot.Send(msg)

		//Тут хочу сделать проверку - всё правильно? Если нет - напишите пожалуйста в чат что не так. и возможно потом вызвать оцените продукт от 1 до 5

		// Okay, we sended our message, dont care about message, so well discard it
		if _, err := bot.Send(msg); err != nil {
			// Notion. Цитата: // Note that panics are a bad way to handle errors. Telegram can
			// have service outages or network errors, you should retry sending
			// messages or more gracefully handle failures.
			panic(err)
		}
	}

}
