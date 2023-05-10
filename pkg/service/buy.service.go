package service

import (
	"fmt"

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

func (s *BuyService) confirmConversation(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	chatID := update.Message.Chat.ID
	if _, exists := buyConversations[chatID]; !exists {
		return tgbotapi.MessageConfig{}, consts.BUY_IS_NOT_STARTED_ERROR
	}
	conversation := buyConversations[chatID]

	res := tgbotapi.NewMessage(
		chatID,
		fmt.Sprintf(
			consts.CONFIRM_BUY_CONVERSATION_MESSAGE,
			conversation.state.answers[0],
			conversation.state.answers[1],
		),
	)

	res.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				consts.CONFIRM_BUY_CONVERSATION_KEYBOARD,
				consts.CONFIRM_BUY_CONVERSATION_KEYBOARD,
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				consts.EDIT_BUY_CONVERSATION_KEYBOARD,
				consts.EDIT_BUY_CONVERSATION_KEYBOARD,
			),
		),
	)

	return res, nil
}

func (s *BuyService) HandleBuyConversation(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	chatID := update.Message.Chat.ID
	var res tgbotapi.MessageConfig
	var err error
	if _, exists := buyConversations[chatID]; !exists {
		return tgbotapi.MessageConfig{}, consts.BUY_IS_NOT_STARTED_ERROR
	}

	conversation := buyConversations[chatID]
	conversation.state.answers = append(conversation.state.answers, update.Message.Text)
	conversation.state.currentQuestion += 1
	buyConversations[chatID] = conversation

	if len(conversation.questsions) < conversation.state.currentQuestion+1 {
		res, err = s.confirmConversation(update)
		if err != nil {
			return tgbotapi.MessageConfig{}, err
		}
	} else {
		res = tgbotapi.NewMessage(chatID, conversation.questsions[conversation.state.currentQuestion])
	}

	return res, nil
}

func (s *BuyService) StartBuy(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	chatID := update.CallbackQuery.Message.Chat.ID
	res := tgbotapi.NewMessage(chatID, consts.START_BUY_MESSAGE)
	buyConversations[chatID] = buyConversation{
		questsions: []string{
			consts.BUY_CONVERSATION_USERNAME_MESSAGE,
			consts.BUY_CONVERSATION_EMAIL_MESSAGE,
		},
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
