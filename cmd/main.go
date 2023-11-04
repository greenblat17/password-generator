package main

import (
	util "bot/pkg"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	state := map[int64]util.UserInput{}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatId := int64(update.Message.Chat.ID)
		msg := util.NewMessage(chatId)

		if _, ok := state[chatId]; !ok {
			state[chatId] = util.UserInput{
				Service:       "",
				SecretKey:     "",
				SecretNumbers: "",
				Count:         0,
			}
		}

		ui := state[chatId]

		if ui.Count > 0 && update.Message.Text == "" {
			if ui.Count == 1 {
				msg.Text = "Введите сервис, в котором будете регистроваться: "
			} else if ui.Count == 2 {
				msg.Text = "Введите ключевое слово: "
			} else if ui.Count == 3 {
				msg.Text = "Введите 4 секректные цифры"
			}
		} else if ui.Count > 0 && !update.Message.IsCommand() {

			if ui.Count == 1 {
				text := update.Message.Text
				if containNums := regexp.MustCompile(`\d`).MatchString(text); containNums {
					msg.Text = "Введите сервис, в котором будете регистроваться: "
				} else {
					ui.Count = 2
					ui.Service = update.Message.Text

					msg.Text = "Введите ключевое слово: "
				}

			} else if ui.Count == 2 {
				ui.Count = 3
				ui.SecretKey = update.Message.Text

				msg.Text = "Введите 4 секректные цифры"
			} else if ui.Count == 3 {
				text := update.Message.Text
				if len(text) != 4 {
					msg.Text = "Введите 4 секректные цифры"
				} else if _, err := strconv.Atoi(text); err != nil {
					msg.Text = "Введите 4 секректные цифры"

				} else {
					ui.SecretNumbers = text

					state[chatId] = ui

					msg.Text = fmt.Sprintf("Your password: %s", util.GeneratePassword(state[chatId]))

					state[chatId] = util.UserInput{
						Service:       "",
						SecretKey:     "",
						SecretNumbers: "",
						Count:         0,
					}
				}
			}
		}

		switch update.Message.Command() {
		case "start":
			msg.Text = "I know command '/password'"
		case "password":
			ui.Count = 1
			msg.Text = "Введите сервис, в котором будете регистроваться: "
		}

		state[chatId] = ui
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

	}
}
