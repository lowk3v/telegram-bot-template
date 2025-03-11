package command

import (
	"github.com/author_name/project_name/internal/tgbot/handler/keyboard"
	"github.com/author_name/project_name/internal/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"reflect"
	"strings"
	"sync"
)

const HelpCommand = "help"
const StartCommand = "start"
const StopCommand = "stop"
const StatusCommand = "status"
const RestartCommand = "restart"

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
		tgbotapi.BotCommand{Command: HelpCommand, Description: "Help info"},
		tgbotapi.BotCommand{Command: StartCommand, Description: "Start the bot"},
		tgbotapi.BotCommand{Command: StopCommand, Description: "Stop the bot"},
		tgbotapi.BotCommand{Command: StatusCommand, Description: "Status of the bot"},
		tgbotapi.BotCommand{Command: RestartCommand, Description: "Restart the bot"},
	))

	return resp, err
}

func (h *Handler) Handle(cmd, arg string) util.HandlerReturnType {
	handleName := "Command"

	parts := strings.Split(cmd, "_")
	for _, part := range parts {
		handleName += strings.ToTitle(part[:1]) + part[1:]
	}

	cmd = strings.Replace(cmd, "_", "", -1)
	result := reflect.ValueOf(h).MethodByName(handleName).Call([]reflect.Value{reflect.ValueOf(arg)})
	return result[0].Interface().(util.HandlerReturnType)
}

func (h *Handler) HandleReplyCommand(incomeStream *tgbotapi.Update) (string, error) {

	//chatId := incomeStream.CallbackQuery.Message.Chat.ID
	//messageId := incomeStream.CallbackQuery.Message.MessageID
	//cmd := strings.Split(incomeStream.CallbackData(), "__")[1]
	//reply := strings.Split(incomeStream.CallbackData(), "__")[2]

	return "", nil
}
