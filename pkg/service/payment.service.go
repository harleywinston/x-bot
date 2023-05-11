package service

import (
	"log"
	"os"

	"github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"

	"github.com/harleywinston/x-bot/pkg/consts"
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
