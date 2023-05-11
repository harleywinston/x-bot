package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/harleywinston/x-bot/pkg"
	"github.com/harleywinston/x-bot/pkg/consts"
	"github.com/harleywinston/x-bot/pkg/service"
	"github.com/harleywinston/x-bot/pkg/webhook"
)

func setupPaymentService() {
	paymentService := service.PaymentService{}
	if err := paymentService.InitPaymentClient(); err != nil {
		log.Fatal(err.Error())
	}
}

func setupWebhook() {
	r := gin.Default()

	paymentWebhooks := webhook.PaymentWebhooks{}
	r.POST(os.Getenv("WEBHOOK_URL_PATH"), paymentWebhooks.CryptoBotWebhook)

	if err := r.Run(":3000"); err != nil {
		log.Fatal(err.Error())
	}
}

func setupMainBot() {
	proxyURL, err := url.Parse(os.Getenv("PROXY_URL"))
	if err != nil {
		err = &consts.CustomError{
			Message: consts.URL_PARSE_ERROR.Message,
			Code:    consts.URL_PARSE_ERROR.Code,
			Detail:  err.Error(),
		}
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

	_, err = bot.Request(tgbotapi.DeleteWebhookConfig{})
	if err != nil {
		log.Fatal(err.Error())
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 1

	botUpdates := bot.GetUpdatesChan(updateConfig)

	for update := range botUpdates {
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

func main() {
	setupPaymentService()
	go setupWebhook()
	go setupMainBot()
	select {}
}
