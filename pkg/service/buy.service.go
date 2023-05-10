package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/harleywinston/x-bot/pkg/consts"
)

type BuyService struct{}

func (s *BuyService) HandleBuyConversation(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	res := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "tmp")
	return res, nil
}

func (s *BuyService) CancelBuy(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	res := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, consts.CANCEL_BUY_MESSAGE)
	return res, nil
}
