package webhook

import (
	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-bot/pkg/service"
)

type PaymentWebhooks struct {
	buyService service.BuyService
}

func (wh *PaymentWebhooks) CryptoBotWebhook(ctx *gin.Context) {
}
