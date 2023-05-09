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

var UPDATE_MESSAGE_ERROR = &CustomError{Message: "Message is invalid!", Code: 500}
