package main

import (
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"wedding-invitation/internal/bot"
	"wedding-invitation/internal/repository"
	"wedding-invitation/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env not found, using system env")
	}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is required")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./storage/wedding.db"
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	adminIDs := parseAdminIDs(os.Getenv("ADMIN_IDS"))

	b, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.Open(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewGuestRepository(db)
	if err := repository.InitDB(db); err != nil {
		log.Fatal(err)
	}
	s := service.NewGuestService(repo, logger)
	bot.SetAdmins(adminIDs)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := b.GetUpdatesChan(u)

	for upd := range updates {

		if upd.Message != nil {
			if upd.Message.IsCommand() {
				bot.HandleCommand(b, s, upd.Message)
			} else {
				bot.HandleText(b, s, upd.Message)
			}
		}

		if upd.CallbackQuery != nil {
			bot.HandleCallback(b, s, upd.CallbackQuery)
		}
	}
}

func parseAdminIDs(raw string) map[int64]bool {
	admins := make(map[int64]bool)

	for _, s := range strings.Split(raw, ",") {
		if s == "" {
			continue
		}
		id, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
		if err == nil {
			admins[id] = true
		}
	}
	return admins
}
