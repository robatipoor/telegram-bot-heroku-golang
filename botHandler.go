package main

import (
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var port string
var token string
var appURL string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	appURL = os.Getenv("APP_URL")
	if appURL == "" {
		log.Fatal("Application URL not set")
	}
	token = os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatalln("TELEGRAM_TOKEN not set !")
	}
}

func botHandler() (*tgbotapi.BotAPI, *tgbotapi.UpdatesChannel) {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	botURL := appURL + ":443/" + token
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(botURL))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	log.Println(info)
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe(":"+port, nil)

	return bot, &updates
}
