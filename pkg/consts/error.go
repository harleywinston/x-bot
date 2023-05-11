package consts

import "fmt"

type CustomError struct {
	Message string
	Code    int
	Detail  string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%d: %s\n%s", e.Code, e.Message, e.Detail)
}

var (
	UPDATE_MESSAGE_ERROR     = &CustomError{Message: "Message is invalid!", Code: 500}
	BUY_IS_NOT_STARTED_ERROR = &CustomError{Message: "Buy didn't start.", Code: 500}
	CRYPTO_BOT_CRASH_ERROR   = &CustomError{
		Message: "Crypto bot api didn't connect properly!",
		Code:    500,
	}
	CRYPTO_BOT_CREATE_INVOICE_ERROR = &CustomError{Message: "Couldn't create invoice!", Code: 500}
	STRING_PARSE_FLOAT_ERROR        = &CustomError{
		Message: "Couldn't parse string to float!",
		Code:    500,
	}
)
