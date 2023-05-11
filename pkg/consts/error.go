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
	UPDATE_MESSAGE_ERROR           = &CustomError{Message: "Message is invalid!", Code: 500}
	BOT_SEND_ERROR                 = &CustomError{Message: "Bot send error!", Code: 500}
	BOT_HANDLE_CALLBACKQUERY_ERROR = &CustomError{
		Message: "Bot handle callback query error!",
		Code:    500,
	}
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
	URL_PARSE_ERROR    = &CustomError{Message: "URL parse error!", Code: 500}
	PARSE_STRING_ERROR = &CustomError{Message: "Parse sting error!", Code: 500}
	BIND_JSON_ERROR    = &CustomError{Message: "Bind json error!", Code: 500}

	CREATE_HTTP_REQ_ERROR    = &CustomError{Message: "Create http request error!", Code: 500}
	MASTER_CREATE_USER_ERROR = &CustomError{Message: "Create user on master error!", Code: 500}
)
