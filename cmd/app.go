package main

import (
	"log"
	"strings"

	env "forbidden-swear-obscenities/pkg/env"
	words "forbidden-swear-obscenities/pkg/words"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	badWords, err := words.ReadWordsFromFile("./words.json")

	if err != nil {
		log.Panic(err)
	}

	botToken := env.Get("TELEGRAM_BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if containsBadWord(update.Message.Text, badWords) {
			deleteConfig := tgbotapi.DeleteMessageConfig{
				ChatID:    update.Message.Chat.ID,
				MessageID: update.Message.MessageID,
			}
			_, err := bot.DeleteMessage(deleteConfig)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}

func containsBadWord(message string, badWords []string) bool {
	for _, badWord := range badWords {
		if strings.Contains(strings.ToLower(message), badWord) {
			return true
		}
	}
	return false
}
