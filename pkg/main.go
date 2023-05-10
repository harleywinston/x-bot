package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/harleywinston/x-bot/pkg/consts"
	"github.com/harleywinston/x-bot/pkg/service"
)

type MessageHandler struct {
	commands service.CommandHandlers
}

func (h *MessageHandler) handleCommands(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	var res tgbotapi.MessageConfig
	var err error
	switch update.Message.Command() {
	case "help":
		res, err = h.commands.HandleHelp(update)
		if err != nil {
			return tgbotapi.MessageConfig{}, err
		}
	case "buy":
		res, err = h.commands.HandleBuy(update)
		if err != nil {
			return tgbotapi.MessageConfig{}, err
		}
	case "status":
		res, err = h.commands.HandleStatus(update)
		if err != nil {
			return tgbotapi.MessageConfig{}, err
		}
	default:
		res = tgbotapi.NewMessage(update.Message.Chat.ID, consts.DEFAULT_COMMAND)
	}
	return res, nil
}

func (h *MessageHandler) handleCallbackQuery(
	update tgbotapi.Update,
) (tgbotapi.MessageConfig, error) {
	var res tgbotapi.MessageConfig
	var err error
	return res, err
}

func (h *MessageHandler) HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	var res tgbotapi.MessageConfig
	var err error
	if update.Message == nil {
		return consts.UPDATE_MESSAGE_ERROR
	}

	if update.Message.IsCommand() {
		res, err = h.handleCommands(update)
		if err != nil {
			return err
		}
	}

	if update.CallbackQuery != nil {
		res, err = h.handleCallbackQuery(update)
		if err != nil {
			return err
		}
	}

	if _, err := bot.Send(res); err != nil {
		return err
	}

	return nil
}
