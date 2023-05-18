package service

import (
	"fmt"
	"log"
	"os"

	"github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"

	"github.com/harleywinston/x-bot/pkg/consts"
	"github.com/harleywinston/x-bot/pkg/models"
)

type PaymentService struct{}

var paymentClient *cryptobot.Client

func (p *PaymentService) InitPaymentClient() error {
	testMode := os.Getenv("CRYPTO_BOT_TEST_MODE")
	apiToken := os.Getenv("CRYPTO_BOT_API_TOKEN")
	paymentClient = cryptobot.NewClient(cryptobot.Options{
		APIToken: apiToken,
		Testing:  testMode == "true",
	})

	appInfo, err := paymentClient.GetMe()
	if err != nil {
		return &consts.CustomError{
			Message: consts.CRYPTO_BOT_CRASH_ERROR.Message,
			Code:    consts.CRYPTO_BOT_CRASH_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	log.Printf(
		"Crypto App info: AppID - %v, Name - %s, PaymentProcessingBotUsername - %s \n",
		appInfo.AppID,
		appInfo.Name,
		appInfo.PaymentProcessingBotUsername,
	)

	return nil
}

func (p *PaymentService) CreateInvoice(user models.UserModel) (*cryptobot.Invoice, error) {
	envPrice := os.Getenv("PRICE")
	invoice, err := paymentClient.CreateInvoice(cryptobot.CreateInvoiceRequest{
		Asset: cryptobot.USDT,
		Amount: func() string {
			if envPrice == "" {
				return "2"
			}
			return envPrice
		}(),
		Description:   fmt.Sprintf(consts.PAYMENT_DESCRIPTION_MESSAGE, user.Email, user.Username),
		HiddenMessage: fmt.Sprintf(consts.PAYMENT_SUCCESS_MESSAGE, user.Email, user.Username),
		PaidBtnName:   "callback",
		PaidBtnUrl:    "https://t.me/fish_proxy_bot",
		Payload: fmt.Sprintf(
			"%d|%s|%s|%v",
			user.ChatID,
			user.Email,
			user.Username,
			user.FuckedUser,
		),
		AllowComments:  true,
		AllowAnonymous: false,
		ExpiresIn:      60 * 10,
	})
	if err != nil {
		return &cryptobot.Invoice{}, &consts.CustomError{
			Message: consts.CRYPTO_BOT_CREATE_INVOICE_ERROR.Message,
			Code:    consts.CRYPTO_BOT_CREATE_INVOICE_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	return invoice, nil
}
