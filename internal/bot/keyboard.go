package bot

import (
	"wedding-invitation/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func attendingKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✅ Приду", models.AttendYes),
			tgbotapi.NewInlineKeyboardButtonData("❌ Не смогу", models.AttendNo),
		),
	)
}

func plusOneKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🙋 Один", models.PlusNo),
			tgbotapi.NewInlineKeyboardButtonData("👫 С +1", models.PlusYes),
		),
	)
}

func mealKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🥩 Мясо", models.MealMeat),
			tgbotapi.NewInlineKeyboardButtonData("🐟 Рыба", models.MealFish),
		),
	)
}

func drinksKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🍷 Да", models.DrinksYes),
			tgbotapi.NewInlineKeyboardButtonData("🚫 Нет", models.DrinksNo),
		),
	)
}

func drinkTypeKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🍷 Вино", models.DrinkWine),
			tgbotapi.NewInlineKeyboardButtonData("🥂 Шампанское", models.DrinkChampagne),
			tgbotapi.NewInlineKeyboardButtonData("🥃 Крепкое", models.DrinkStrong),
		),
	)
}
