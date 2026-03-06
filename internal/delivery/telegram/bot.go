package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ks1ze66/go-rpg-bot/internal/repository"
)

// Bot - основная структура нашего телеграм-интерфейса
type Bot struct {
	api  *tgbotapi.BotAPI
	repo *repository.UserRepository
}

// NewBot - конструктор, который принимает токен и репозиторий для работы с БД
func NewBot(token string, repo *repository.UserRepository) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		api:  api,
		repo: repo,
	}, nil
}

// Start запускает бесконечный цикл прослушивания сообщений
func (b *Bot) Start() error {
	log.Printf("Бот запущен под именем %s", b.api.Self.UserName)

	// Настраиваем конфиг получения обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // Игнорируем всё, кроме сообщений
			continue
		}

		// Логика обработки команды /start
		if update.Message.IsCommand() && update.Message.Command() == "start" {
			b.handleStart(update.Message)
		}
	}
	return nil
}

// handleStart отвечает за регистрацию пользователя
func (b *Bot) handleStart(msg *tgbotapi.Message) {
	// Мы пока просто отправим текстовый ответ
	reply := "Привет, искатель! Твой путь в Go Middle начинается здесь."
	newMsg := tgbotapi.NewMessage(msg.Chat.ID, reply)
	b.api.Send(newMsg)
}
