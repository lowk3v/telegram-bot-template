package tgbot

import (
	"fmt"
	config "github.com/author_name/project_name/configs"
	"github.com/author_name/project_name/internal/tgbot/handler/command"
	"github.com/author_name/project_name/internal/tgbot/handler/keyboard"
	"github.com/author_name/project_name/internal/tgbot/handler/message"
	"github.com/author_name/project_name/internal/util/incoming"
	"github.com/author_name/project_name/pkg/pubsub/consumer"
	"github.com/fatih/color"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

type TgBot struct {
	bot            *tgbotapi.BotAPI
	keyboardHandle *keyboard.Handler
	messageHandle  *message.Handler
	commandHandle  *command.Handler
	consumer       *consumer.Consumer
}

func New() *TgBot {
	bot, err := tgbotapi.NewBotAPI(config.Secret.Token)
	if err != nil {
		if config.Secret.Token == "" {
			config.Log.Errorf("Token is empty")
		} else {
			config.Log.Errorf("Error: %v", err)
		}
		os.Exit(1)
	}
	config.Log.WithField("account", color.BlueString(bot.Self.UserName)).
		Info("Authorized on account")

	keyboardHandle := keyboard.New(bot)
	messageHandle := message.New(bot, keyboardHandle)
	commandHandle := command.New(bot, keyboardHandle)
	resp, err := commandHandle.SetNewCommand()
	if err != nil {
		config.Log.Errorf("Error: %v", err)
		config.Log.Debugf("Set bot commands response: %v", resp)
	}

	return &TgBot{
		bot:            bot,
		keyboardHandle: keyboardHandle,
		messageHandle:  messageHandle,
		commandHandle:  commandHandle,
		consumer:       consumer.New(config.GlobalPublisher),
	}
}

func (tg *TgBot) RunBotController() {
	config.Log.Info("Start bot controller")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	for incomeStream := range tg.bot.GetUpdatesChan(u) {
		// Initialize message output or callback query output
		var err error
		var replyMarkup interface{}
		replyMsg := ""
		var outStream tgbotapi.MessageConfig

		if incoming.Classify(incomeStream) == "callback" {
			outStream = tgbotapi.NewMessage(incomeStream.CallbackQuery.Message.Chat.ID, "")
		} else {
			outStream = tgbotapi.NewMessage(incomeStream.Message.Chat.ID, incomeStream.Message.Text)
		}

		// Handle event
		switch incoming.Classify(incomeStream) {
		case "addOrRemoveGroup":
			handleEventAddOrRemoveFromGroup(incomeStream)
			break
		case "callback":
			handleCallback(tg.bot, tg.commandHandle, incomeStream)
			continue
		case "ignore":
			break
		case "nothing":
			tg.keyboardHandle.Close(&outStream)
			break
		case "command":
			replyMsg, replyMarkup, err = tg.commandHandle.Handle(incomeStream.Message.Command(), incomeStream.Message.CommandArguments())
			if err != nil {
				config.Log.Errorf("Error: %v", err)
				continue
			}
			break
		case "new_message":
			// Message handler
			replyMsg, err = tg.messageHandle.Handle(&incomeStream)
			if err != nil {
				config.Log.Errorf("Error: %v", err)
				// return err to user
			}
			break
		}

		if err != nil {
			replyMsg = err.Error()
			config.Log.Errorf("Error: %v", err)
		}

		reloadKeyboard(tg.keyboardHandle, &outStream, replyMsg, replyMarkup)
		sendToTelegram(tg.bot, outStream)

	}
}

func handleEventAddOrRemoveFromGroup(incomeStream tgbotapi.Update) {
	if incomeStream.MyChatMember.Chat.IsGroup() || incomeStream.MyChatMember.Chat.IsSuperGroup() {
		if incomeStream.MyChatMember.NewChatMember.Status == "member" {
			// add to group
			config.Log.WithField("group", incomeStream.MyChatMember.Chat.Title).
				WithField("by", incomeStream.MyChatMember.From.UserName).
				Info("Add to")
		} else if incomeStream.MyChatMember.NewChatMember.Status == "left" {
			{
				// remove from group
				config.Log.WithField("group", incomeStream.MyChatMember.Chat.Title).
					WithField("by", incomeStream.MyChatMember.From.UserName).
					Info("Remove from")
			}
		}
	}
}

func reloadKeyboard(keyboardHandle *keyboard.Handler, outStream *tgbotapi.MessageConfig, replyMsg string, replyMarkup interface{}) {
	if replyMarkup == nil {
		keyboardHandle.Close(outStream)
	} else {
		outStream.ReplyMarkup = replyMarkup
		outStream.Text = replyMsg
	}
}

func sendToTelegram(bot *tgbotapi.BotAPI, msgConfig tgbotapi.MessageConfig) {
	if msgConfig.Text == "" {
		return
	}
	msgConfig.ParseMode = tgbotapi.ModeMarkdownV2
	msgConfig.DisableWebPagePreview = true
	_, err := bot.Send(msgConfig)
	if err != nil {
		config.Log.Errorf("Error: %v", err)
	}
}

func editMessageTelegram(bot *tgbotapi.BotAPI, msgConfig tgbotapi.EditMessageTextConfig) {
	if msgConfig.Text == "" {
		return
	}
	msgConfig.ParseMode = tgbotapi.ModeMarkdownV2
	msgConfig.DisableWebPagePreview = true
	_, err := bot.Send(msgConfig)
	if err != nil {
		config.Log.Errorf("Error: %v", err)
	}
}

func handleCallback(bot *tgbotapi.BotAPI, commandHandle *command.Handler, incomeStream tgbotapi.Update) {
	chatId := incomeStream.CallbackQuery.Message.Chat.ID
	messageId := incomeStream.CallbackQuery.Message.MessageID

	// alert
	callback := tgbotapi.NewCallback(incomeStream.CallbackQuery.ID, "Waiting...")
	_, err := bot.Request(callback)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v", err)
		return
	}

	// When a user press button on an inline keyboard
	replyMsg, _ := commandHandle.HandleReplyCommand(&incomeStream)

	if replyMsg == "" {
		return
	}

	outStream := tgbotapi.NewEditMessageText(chatId, messageId, replyMsg)
	outStream.ParseMode = tgbotapi.ModeMarkdownV2
	outStream.DisableWebPagePreview = true
	_, err = bot.Send(outStream)
	if err != nil {
		config.Log.Errorf("Error: %v", err)
		return
	}
}

func (tg *TgBot) Stop() {
	tg.consumer.Unsubscribe()
}
