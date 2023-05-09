package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/harleywinston/x-bot/pkg/consts"
)

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	if update.Message == nil {
		return consts.UPDATE_MESSAGE_ERROR
	}

	res := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "help":
			res.Text = "good."
		default:
			res.Text = "fuck you!"
		}
	}

	if _, err := bot.Send(res); err != nil {
		return err
	}

	return nil
}
