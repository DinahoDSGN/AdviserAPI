package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func Init() {
	bot, err := tgbotapi.NewBotAPI("2017691701:AAHj7yS3wdTNMfAn2HE_Cbj5CC4uuifffKE")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	tgbot := NewBot(bot)
	tgbot.Start()

	log.Printf("Authorized on account %s", bot.Self.UserName)

}
