package message

import (
	"github.com/author_name/project_name/internal/tgbot/handler/keyboard"
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

func (h *Handler) Handle(incomeStream *tgbotapi.Update) (string, error) {
	message := incomeStream.Message.Text
	if message == "" {
		return "", nil
	}
	return message, nil
}
