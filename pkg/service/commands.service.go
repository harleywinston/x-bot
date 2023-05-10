package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/harleywinston/x-bot/pkg/consts"
)

type CommandHandlers struct{}

func (h *CommandHandlers) HandleHelp(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	res := tgbotapi.NewMessage(update.Message.Chat.ID, consts.HELP_COMMAND)
	return res, nil
}

func (h *CommandHandlers) HandleBuy(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	res := tgbotapi.NewMessage(update.Message.Chat.ID, consts.BUY_COMMAND)
	return res, nil
}

func (h *CommandHandlers) HandleStatus(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	res := tgbotapi.NewMessage(update.Message.Chat.ID, consts.STATUS_COMMAND)
	return res, nil
}
