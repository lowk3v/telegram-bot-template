package incoming

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Classify must be ordered by logics
func Classify(in tgbotapi.Update) string {
	if in.MyChatMember != nil {
		return "addOrRemoveGroup"
	}
	if in.CallbackQuery != nil {
		return "callback"
	}
	if in.Message == nil {
		return "ignore"
	}
	if in.Message.Text == "" {
		return "nothing"
	}
	if in.Message.IsCommand() {
		return "command"
	}
	if in.Message.Text != "" {
		return "new_message"
	}

	if in.EditedMessage != nil {
		return "edited_message"
	}
	if in.ChannelPost != nil {
		return "channel_post"
	}
	if in.EditedChannelPost != nil {
		return "edited_channel_post"
	}
	if in.InlineQuery != nil {
		return "inline_query"
	}
	if in.ChosenInlineResult != nil {
		return "chosen_inline_result"
	}
	if in.ShippingQuery != nil {
		return "shipping_query"
	}
	if in.PreCheckoutQuery != nil {
		return "pre_checkout_query"
	}
	if in.Poll != nil {
		return "poll"
	}
	if in.PollAnswer != nil {
		return "poll_answer"
	}
	if in.ChatMember != nil {
		return "chat_member"
	}
	return "unknown"
}

func MessageClassify(message string) string {
	return ""
}
