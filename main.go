package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	bot,updates := botHandler()
	
	for update := range *updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		chatID := update.Message.Chat.ID
		userName := update.Message.From.UserName
		log.Printf("[%s] %s", userName, update.Message.Text)
		msg := tgbotapi.NewMessage(chatID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
