package message

import (
	"github.com/author_name/project_name/internal/tgbot/handler/keyboard"
	"github.com/author_name/project_name/internal/util"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	sync sync.Mutex
	wg   *sync.WaitGroup

	bot             *tgbotapi.BotAPI
	keyboardHandler *keyboard.Handler
}

func New(bot *tgbotapi.BotAPI, k *keyboard.Handler) *Handler {
	return &Handler{
		sync:            sync.Mutex{},
		wg:              &sync.WaitGroup{},
		bot:             bot,
		keyboardHandler: k,
	}
}

func (h *Handler) Handle(incomeStream *tgbotapi.Update) util.HandlerReturnType {
	message := incomeStream.Message.Text
	return util.HandlerReturnType{
		ReplyMsg: message + " does not supported",
	}
}
