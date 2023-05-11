package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/harleywinston/x-bot/pkg/consts"
	"github.com/harleywinston/x-bot/pkg/models"
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

func (s *BuyService) ProceedPayment(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	var res tgbotapi.MessageConfig
	chatID := update.CallbackQuery.Message.Chat.ID

	if _, exists := buyConversations[chatID]; !exists {
		return tgbotapi.MessageConfig{}, consts.BUY_IS_NOT_STARTED_ERROR
	}
	user := models.UserModel{
		ChatID:   chatID,
		Username: buyConversations[chatID].state.answers[0],
		Email:    buyConversations[chatID].state.answers[1],
	}

	paymentService := PaymentService{}
	invoice, err := paymentService.CreateInvoice(user)
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}

	res = tgbotapi.NewMessage(chatID, consts.PROCEED_PAYMENT_MESSAGE)
	res.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.InlineKeyboardButton{
				Text:         consts.PROCEED_PAYMENT_KEYBOARD,
				Pay:          true,
				URL:          &invoice.PayUrl,
				CallbackData: &consts.PROCEED_PAYMENT_KEYBOARD,
			},
		),
	)

	return res, nil
}

func (s *BuyService) ProceedAfterPayment(user models.UserModel) ([]tgbotapi.MessageConfig, error) {
	var res []tgbotapi.MessageConfig
	res = append(res, tgbotapi.NewMessage(user.ChatID, consts.PROCEED_AFTER_PAYMENT_MESSAGE))

	HTTPClient := &http.Client{}
	baseURL := os.Getenv("MASTER_URL")
	jsonBody, err := json.Marshal(user)
	if err != nil {
		return []tgbotapi.MessageConfig{}, &consts.CustomError{
			Message: consts.BIND_JSON_ERROR.Message,
			Code:    consts.BIND_JSON_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	req, err := http.NewRequest(http.MethodPost, baseURL+"/user", bytes.NewBuffer(jsonBody))
	if err != nil {
		return []tgbotapi.MessageConfig{}, &consts.CustomError{
			Message: consts.CREATE_HTTP_REQ_ERROR.Message,
			Code:    consts.CREATE_HTTP_REQ_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := HTTPClient.Do(req)
	if err != nil {
		return []tgbotapi.MessageConfig{}, &consts.CustomError{
			Message: consts.CREATE_HTTP_REQ_ERROR.Message,
			Code:    consts.CREATE_HTTP_REQ_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	if resp.StatusCode != 200 {
		return []tgbotapi.MessageConfig{}, &consts.CustomError{
			Message: consts.MASTER_CREATE_USER_ERROR.Message,
			Code:    consts.MASTER_CREATE_USER_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	return res, nil
}

func (s *BuyService) EditBuyConversation(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	var res tgbotapi.MessageConfig
	var err error
	chatID := update.CallbackQuery.Message.Chat.ID

	_, err = s.CancelBuy(update)
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}
	_, err = s.StartBuy(update)
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}

	res = tgbotapi.NewMessage(chatID, consts.EDIT_BUY_CONVERSATIN_MESSAGE)

	return res, nil
}
