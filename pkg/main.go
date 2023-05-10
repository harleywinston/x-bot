package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/harleywinston/x-bot/pkg/consts"
	"github.com/harleywinston/x-bot/pkg/service"
)

type MessageHandler struct {
	commands   service.CommandHandlers
	buyService service.BuyService
}

func (h *MessageHandler) handleCommands(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	var res tgbotapi.MessageConfig
	var err error
	switch update.Message.Command() {
	case consts.HELP_COMMAND:
		res, err = h.commands.HandleHelp(update)
		if err != nil {
			return tgbotapi.MessageConfig{}, err
		}
	case consts.BUY_COMMAND:
		res, err = h.commands.HandleBuy(update)
		if err != nil {
			return tgbotapi.MessageConfig{}, err
		}
	case consts.STATUS_COMMAND:
		res, err = h.commands.HandleStatus(update)
		if err != nil {
			return tgbotapi.MessageConfig{}, err
		}
	default:
		return tgbotapi.MessageConfig{}, consts.UPDATE_MESSAGE_ERROR
	}
	return res, nil
}

func (h *MessageHandler) handleCallbackQuery(
	update tgbotapi.Update,
) (tgbotapi.MessageConfig, error) {
	var res tgbotapi.MessageConfig
	var err error

	switch update.CallbackQuery.Data {
	case consts.START_BUY_KEYBOARD:
		res, err = h.buyService.StartBuy(update)
		if err != nil {
			return tgbotapi.MessageConfig{}, err
		}
	case consts.CANCEL_BUY_KEYBOARD:
		res, err = h.buyService.CancelBuy(update)
		if err != nil {
			return tgbotapi.MessageConfig{}, err
		}
	default:
		return tgbotapi.MessageConfig{}, consts.UPDATE_MESSAGE_ERROR
	}
	return res, nil
}

func (h *MessageHandler) HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	var res tgbotapi.MessageConfig
	var err error

	if update.Message != nil {
		if update.Message.IsCommand() {
			res, err = h.handleCommands(update)
			if err != nil {
				return err
			}
		} else {
			res, err = h.buyService.HandleBuyConversation(update)
			if err != nil {
				return err
			}
		}
	}

	if update.CallbackQuery != nil {
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
		if _, err := bot.Request(callback); err != nil {
			return err
		}

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
