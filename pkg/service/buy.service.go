package service

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BuyService struct {
	Bot    *tgbotapi.BotAPI
	Update tgbotapi.Update
}

func (s *BuyService) Buy() {
}
