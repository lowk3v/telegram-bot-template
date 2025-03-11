package util

type HandlerReturnType struct {
	ReplyMsg    string
	ReplyMarkup interface{}
	Error       error
}
