package util

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func HandleCommand(command string, msg tgbotapi.MessageConfig) {
	switch command {
	case "start", "help":
		setDefaultMessage(msg)
	case "password":
		setMessageForGeneratePassword(msg)
	}
}
