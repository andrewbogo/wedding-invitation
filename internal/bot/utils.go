package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// sendMessage отправляет текстовое сообщение в чат с необязательной клавиатурой
func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string, markup interface{}) {
	msg := tgbotapi.NewMessage(chatID, text)

	if markup != nil {
		msg.ReplyMarkup = markup
	}

	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
