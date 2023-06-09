package consts

var (
	HELP_COMMAND    = "help"
	BUY_COMMAND     = "buy"
	STATUS_COMMAND  = "status"
	START_COMMAND   = "start"
	DEFAULT_COMMAND = "unkown"

	HELP_COMMAND_MESSAGE    = "help command"
	BUY_COMMAND_MESSAGE     = "buy command"
	STATUS_COMMAND_MESSAGE  = "status command"
	DEFAULT_COMMAND_MESSAGE = "default command"

	START_BUY_KEYBOARD  = "start"
	START_BUY_MESSAGE   = "start buy"
	CANCEL_BUY_KEYBOARD = "cancle"
	CANCEL_BUY_MESSAGE  = "cancel succeed"

	DEFAULT_CALLBACK_MESSAGE = "default callback message"

	BUY_CONVERSATION_USERNAME_MESSAGE = "username:"
	BUY_CONVERSATION_EMAIL_MESSAGE    = "email:"
	CONFIRM_BUY_CONVERSATION_MESSAGE  = `
	 username: %s
	email: %s
	`
	CONFIRM_BUY_CONVERSATION_KEYBOARD = "confirm"
	EDIT_BUY_CONVERSATION_KEYBOARD    = "edit"
	EDIT_BUY_CONVERSATIN_MESSAGE      = "edit conversation"

	PROCEED_PAYMENT_MESSAGE       = "pay!"
	PROCEED_PAYMENT_KEYBOARD      = "go to pay."
	PAYMENT_DESCRIPTION_MESSAGE   = "paying for email %s with username %s"
	PAYMENT_SUCCESS_MESSAGE       = "you payed for email %s with username %s"
	PROCEED_AFTER_PAYMENT_MESSAGE = "you bought!"
	LAST_MESSAGE_AFTER_BUY        = "here is instruction:"

	INTERNAL_ERROR_CONTACT_SUPPORT_MESSGE = "something bad happened please contact @support"
)
