package bot

import (
	"log"

	"wedding-invitation/internal/models"
	"wedding-invitation/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleStart(bot *tgbotapi.BotAPI, service *service.GuestService, msg *tgbotapi.Message) {
	err := service.InitGuest(
		msg.From.ID,
		msg.From.FirstName,
		msg.From.UserName,
	)
	if err != nil {
		return
	}
	m := tgbotapi.NewMessage(msg.Chat.ID, "Мы будем очень рады видеть тебя на нашей свадьбе!\\n\\nПодтвердишь участие?")
	m.ReplyMarkup = attendingKeyboard()
	_, err = bot.Send(m)
	if err != nil {
		log.Println(err)
	}
}

func HandleCommand(bot *tgbotapi.BotAPI, service *service.GuestService, msg *tgbotapi.Message) {
	switch msg.Command() {
	case "stats":
		if !IsAdmin(msg.From.ID) {
			return
		}
		// HandleStats(bot, db, msg)

	case "export":
		if !IsAdmin(msg.From.ID) {
			return
		}
		// HandleExport(bot, db, msg)

	case "start":
		HandleStart(bot, service, msg)
	}
}

func HandleText(bot *tgbotapi.BotAPI, service *service.GuestService, msg *tgbotapi.Message) {
	guestID := msg.From.ID
	guest, err := service.GetGuest(guestID)
	if err != nil {
		return
	}
	switch guest.State {
	case models.StateWaitingName:

		err = service.SetName(msg.From.ID, msg.Text)
		if err != nil {
			return
		}

		if guest.PlusOne.Valid && guest.PlusOne.Bool {
			err = service.UpdateState(msg.From.ID, models.StateWaitingPlusOneName)
			if err != nil {
				return
			}
			sendMessage(bot, msg.Chat.ID, "Как зовут твоего +1? (Фамилия и Имя через пробел)", nil)
		} else {
			_ = service.UpdateState(msg.From.ID, models.StateWaitingMeal)
			sendMessage(bot, msg.Chat.ID, "Какое горячее ты предпочитаешь? 🍽", mealKeyboard())
		}

	case models.StateWaitingPlusOneName:

		_ = service.SetPlusOneName(msg.From.ID, msg.Text)

		_ = service.UpdateState(msg.From.ID, models.StateWaitingMeal)
		sendMessage(bot, msg.Chat.ID, "Какое горячее ты предпочитаешь? 🍽", mealKeyboard())
	}
}
