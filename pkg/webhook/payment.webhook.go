package webhook

import (
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/arthurshafikov/cryptobot-sdk-golang/cryptobot"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/harleywinston/x-bot/pkg/consts"
	"github.com/harleywinston/x-bot/pkg/models"
	"github.com/harleywinston/x-bot/pkg/service"
)

type PaymentWebhooks struct {
	Bot        *tgbotapi.BotAPI
	buyService service.BuyService
}

func (wh *PaymentWebhooks) CryptoBotWebhook(ctx *gin.Context) {
	var messages []tgbotapi.MessageConfig
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
			"message": consts.BIND_JSON_ERROR.Message,
			"detail":  err.Error(),
		})
		return
	}
	update, err := cryptobot.ParseWebhookUpdate(body)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
			"message": consts.BIND_JSON_ERROR.Message,
			"detail":  err.Error(),
		})
		return
	}

	switch update.UpdateType {
	case cryptobot.InvoicePaidWebhookUpdateType:
		invoice, err := cryptobot.ParseInvoice(body)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
				"message": consts.BIND_JSON_ERROR.Message,
				"detail":  err.Error(),
			})
			return
		}

		payloadData := strings.Split(strings.ReplaceAll(invoice.Payload, " ", ""), "|")
		if len(payloadData) < 4 {
			log.Println("payload len")
			ctx.JSON(consts.CRYPTO_BOT_PAYLOAD_ERROR.Code, gin.H{
				"message": consts.CRYPTO_BOT_PAYLOAD_ERROR.Message,
				"detail":  "payload data len is less than 4!",
			})
		}
		chatID, err := strconv.ParseInt(payloadData[0], 10, 64)
		if err != nil {
			log.Println("invalid chat id")
			ctx.JSON(consts.PARSE_STRING_ERROR.Code, gin.H{
				"message": consts.PARSE_STRING_ERROR.Message,
				"detail":  err.Error(),
			})
			return
		}
		email := payloadData[1]
		username := payloadData[2]
		fuckedUser, err := strconv.ParseBool(payloadData[3])
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(consts.PARSE_STRING_ERROR.Code, gin.H{
				"message": consts.PARSE_STRING_ERROR.Message,
				"detail":  err.Error(),
			})
			return
		}
		user := models.UserModel{
			ChatID:     chatID,
			Email:      email,
			Username:   username,
			FuckedUser: fuckedUser,
		}

		messages, err = wh.buyService.ProceedAfterPayment(user)
		if err != nil {
			log.Println(err.Error())
			msg1 := tgbotapi.NewMessage(user.ChatID, err.Error())
			msg2 := tgbotapi.NewMessage(user.ChatID, consts.INTERNAL_ERROR_CONTACT_SUPPORT_MESSGE)
			if _, err := wh.Bot.Send(msg1); err != nil {
				log.Println(err.Error())
			}
			if _, err := wh.Bot.Send(msg2); err != nil {
				log.Println(err.Error())
			}
			return
		}
	default:
		log.Println("Unsupported webhook type" + update.UpdateType)
	}

	ctx.JSON(200, nil)

	for _, msg := range messages {
		if _, err := wh.Bot.Send(msg); err != nil {
			log.Println(err.Error())
		}
	}
}
