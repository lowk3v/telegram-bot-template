package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot *tgbotapi.BotAPI
}

func New(bot *tgbotapi.BotAPI) *Handler {
	return &Handler{
		bot: bot,
	}
}

func (h *Handler) Open(outStream *tgbotapi.MessageConfig, text string, btns []string) {
	if text == "" {
		text = h.HandleHelp()
	}
	outStream.Text = text
	outStream.ReplyMarkup = h.loadKeyboard(btns)
}

func (h *Handler) Close(outStream *tgbotapi.MessageConfig) {
	outStream.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
}

func (h *Handler) LoadInlineKeyboard(btnTitle ...[]string) tgbotapi.InlineKeyboardMarkup {
	var btns []tgbotapi.InlineKeyboardButton

	for _, btn := range btnTitle {
		aRow := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(btn[0], btn[1]),
		)
		btns = append(btns, aRow...)
	}

	return tgbotapi.NewInlineKeyboardMarkup(btns)
}

func (h *Handler) loadKeyboard(btnTitles []string) tgbotapi.ReplyKeyboardMarkup {
	defaultBtn := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Cancel"),
	)
	var btns []tgbotapi.KeyboardButton
	for _, title := range btnTitles {
		aRow := tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(title),
		)
		btns = append(btns, aRow...)
	}

	if len(btns) > 0 {
		return tgbotapi.NewOneTimeReplyKeyboard(btns, defaultBtn)
	}
	return tgbotapi.NewOneTimeReplyKeyboard()
}

func (h *Handler) HandleHelp() string {
	return `
		Welcome to the bot. You can use the following commands:
		- /help: show this message
`
}
