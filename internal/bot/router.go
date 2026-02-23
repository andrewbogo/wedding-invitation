package bot

import (
	"wedding-invitation/internal/models"
	"wedding-invitation/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCallback(bot *tgbotapi.BotAPI, service *service.GuestService, q *tgbotapi.CallbackQuery) {
	switch q.Data {
	case models.AttendYes:
		err := service.SetAttending(q.From.ID, true)
		if err != nil {

		}

		msg := tgbotapi.NewMessage(q.Message.Chat.ID,
			"Отлично! 🙌 Ты придёшь один или с парой?")
		msg.ReplyMarkup = plusOneKeyboard()
		_, _ = bot.Send(msg)

	case models.AttendNo:
		_ = service.SetAttending(q.From.ID, false)

		msg := tgbotapi.NewMessage(q.Message.Chat.ID,
			"Очень жаль... Будем думать о тебе в этот день.")
		_, _ = bot.Send(msg)
	case models.PlusYes:
		_ = service.SetPlusOne(q.From.ID, true)
		_ = service.UpdateState(q.From.ID, models.StateWaitingName)

		sendMessage(bot, q.Message.Chat.ID,
			"Как тебя зовут? (Фамилия и Имя через пробел)", nil)

	case models.PlusNo:
		_ = service.SetPlusOne(q.From.ID, false)
		_ = service.UpdateState(q.From.ID, models.StateWaitingName)

		sendMessage(bot, q.Message.Chat.ID,
			"Как тебя зовут? (Фамилия и Имя через пробел)", nil)

	case models.MealMeat:
		_ = service.SetMeal(q.From.ID, "meat")

		msg := tgbotapi.NewMessage(q.Message.Chat.ID,
			"Планируешь ли ты алкоголь? 🍷")
		msg.ReplyMarkup = drinksKeyboard()
		_, _ = bot.Send(msg)

	case models.MealFish:
		_ = service.SetMeal(q.From.ID, "fish")

		msg := tgbotapi.NewMessage(q.Message.Chat.ID,
			"Планируешь ли ты алкоголь? 🍷")
		msg.ReplyMarkup = drinksKeyboard()
		_, _ = bot.Send(msg)

	case models.DrinksYes:
		_ = service.SetDrinks(q.From.ID, true)

		msg := tgbotapi.NewMessage(q.Message.Chat.ID,
			"Что предпочитаешь?")
		msg.ReplyMarkup = drinkTypeKeyboard()
		_, _ = bot.Send(msg)

	case models.DrinksNo:
		_ = service.SetDrinks(q.From.ID, false)

		_, _ = bot.Send(tgbotapi.NewMessage(q.Message.Chat.ID,
			"Отлично, всё записали! 🎉"))

	case models.DrinkWine:
		_ = service.SetDrinkType(q.From.ID, "wine")

	case models.DrinkChampagne:
		_ = service.SetDrinkType(q.From.ID, "champagne")

	case models.DrinkStrong:
		_ = service.SetDrinkType(q.From.ID, "strong")
	}

	_, _ = bot.Request(tgbotapi.NewCallback(q.ID, ""))
}
