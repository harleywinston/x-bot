package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/harleywinston/x-bot/pkg/consts"
)

type BuyService struct{}

type userState struct {
	currentQuestion int
	answers         []string
}

type buyConversation struct {
	questsions []string
	state      userState
}

var buyConversations = make(map[int64]buyConversation)

func (s *BuyService) HandleBuyConversation(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	chatID := update.Message.Chat.ID
	res := tgbotapi.NewMessage(chatID, "tmp")
	if _, exists := buyConversations[chatID]; !exists {
		return tgbotapi.MessageConfig{}, consts.BUY_IS_NOT_STARTED_ERROR
	}
	return res, nil
}

func (s *BuyService) StartBuy(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	chatID := update.CallbackQuery.Message.Chat.ID
	res := tgbotapi.NewMessage(chatID, consts.START_BUY_MESSAGE)
	buyConversations[chatID] = buyConversation{
		questsions: []string{consts.BUY_CONVERSATION_USERNAME_MESSAGE},
		state: userState{
			currentQuestion: 0,
			answers:         []string{},
		},
	}
	return res, nil
}

func (s *BuyService) CancelBuy(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	chatID := update.CallbackQuery.Message.Chat.ID
	res := tgbotapi.NewMessage(chatID, consts.CANCEL_BUY_MESSAGE)
	delete(buyConversations, chatID)
	return res, nil
}
