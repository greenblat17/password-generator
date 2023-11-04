package util

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func setDefaultMessage(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	msg.Text = "I know command '/password'"
	return msg
}

func setMessageForGeneratePassword(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	msg.Text = "Введите сервис, в котором будете регистроваться: "
	return msg
}

func NewMessage(chatId int64) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatId, "")
}
