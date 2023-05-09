package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/harleywinston/x-bot/pkg"
)

func main() {
	proxyURL, err := url.Parse(os.Getenv("PROXY_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}

	proxyClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}

	bot, err := tgbotapi.NewBotAPIWithClient(
		os.Getenv("BOT_TOKEN"),
		"https://api.telegram.org/bot%s/%s",
		proxyClient,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 1

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		err := pkg.HandleMessage(bot, update)
		if err == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
		if _, err := bot.Send(msg); err != nil {
			log.Println(err.Error())
		}
	}
}
