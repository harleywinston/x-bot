package service_test

import (
	"testing"

	"github.com/harleywinston/x-bot/pkg/models"
	"github.com/harleywinston/x-bot/pkg/service"
)

type buyServiceTest struct {
	buyService service.BuyService
}

func (tb *buyServiceTest) TestAfterPayment(t *testing.T) {
	tests := []models.UserModel{
		{
			ChatID:     0,
			Username:   "harley",
			Email:      "harleywinston@proton.me",
			FuckedUser: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Email, func(t *testing.T) {
			_, err := tb.buyService.ProceedAfterPayment(test)
			if err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func TestBuyService(t *testing.T) {
	tests := buyServiceTest{}

	t.Run("Test after payment", tests.TestAfterPayment)
}
