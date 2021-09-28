package telegram

import (
	"AdviserAPI/pkg/telegram/router"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (b *Bot) CommandHandler(message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		msg := "Hi " + message.Chat.UserName + ", type /help command to get available commands"
		newMessage := tgbotapi.NewMessage(message.Chat.ID, msg)
		b.bot.Send(newMessage)
		return
	case "help":
		msg := "/go 1-5 - _get an advice for number of person_"
		newMessage := tgbotapi.NewMessage(message.Chat.ID, msg)
		newMessage.ParseMode = "markdown"
		b.bot.Send(newMessage)
		return
	case "go":
		param := message.CommandArguments()
		intParam, _ := strconv.Atoi(param)
		request, _ := router.GetAdvice(intParam)

		newMessage := tgbotapi.NewMessage(message.Chat.ID, request)
		b.bot.Send(newMessage)
		return
	default:
		newMessage := tgbotapi.NewMessage(message.Chat.ID, "Such command does not exist!")
		b.bot.Send(newMessage)
	}
}
