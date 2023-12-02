package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

type Bot struct {
	telegramBot *tgbotapi.BotAPI
	chatId      int64
}

func NewBot(env Env) *Bot {
	telegramBot, err := tgbotapi.NewBotAPI(env.BotToken)
	if err != nil {
		log.Fatal("Error for creating bot")
	}

	return &Bot{telegramBot: telegramBot, chatId: env.ChatId}
}

func (bot Bot) Notify(event time.Time) {
	message := tgbotapi.NewMessage(bot.chatId, "З'явилось мiсце - "+event.Format("2006-01-02"))

	if _, err := bot.telegramBot.Send(message); err != nil {
		log.Fatal(err)
	}
}
