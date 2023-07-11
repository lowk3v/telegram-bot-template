package command

import (
	"github.com/author_name/project_name/internal/tgbot/handler/keyboard"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
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

func (h *Handler) SetNewCommand() (*tgbotapi.APIResponse, error) {
	resp, err := h.bot.Request(tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{Command: HELP, Description: "Help info"},
	))
	return resp, err
}

func (h *Handler) Handle(cmd, arg string) (string, interface{}, error) {
	switch cmd {
	case HELP:
		return h.helpInfo(), nil, nil
	}
	return "", nil, nil
}

func (h *Handler) HandleReplyCommand(incomeStream *tgbotapi.Update) (string, error) {

	//chatId := incomeStream.CallbackQuery.Message.Chat.ID
	//messageId := incomeStream.CallbackQuery.Message.MessageID
	//cmd := strings.Split(incomeStream.CallbackData(), "__")[1]
	//reply := strings.Split(incomeStream.CallbackData(), "__")[2]

	return "", nil
}

func (h *Handler) helpInfo() string {
	return "Help info"
}
