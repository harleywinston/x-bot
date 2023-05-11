package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/harleywinston/x-bot/pkg"
	"github.com/harleywinston/x-bot/pkg/service"
)

func main() {
	paymentService := service.PaymentService{}
	if err := paymentService.InitPaymentClient(); err != nil {
		log.Fatal(err)
	}

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
		messageHandler := pkg.MessageHandler{}
		err := messageHandler.HandleMessage(bot, update)
		if err == nil {
			continue
		}

		var chatID int64
		if update.Message != nil {
			chatID = update.Message.Chat.ID
		}
		if update.CallbackQuery != nil {
			chatID = update.CallbackQuery.Message.Chat.ID
		}
		msg := tgbotapi.NewMessage(chatID, err.Error())
		if _, err := bot.Send(msg); err != nil {
			log.Println(err.Error())
		}
	}
}
