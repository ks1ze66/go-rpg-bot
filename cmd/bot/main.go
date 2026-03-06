package main

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib" // Импорт драйвера
	"github.com/jmoiron/sqlx"
	"github.com/ks1ze66/go-rpg-bot/internal/config"
	"github.com/ks1ze66/go-rpg-bot/internal/delivery/telegram"
	"github.com/ks1ze66/go-rpg-bot/internal/domain"
	"github.com/ks1ze66/go-rpg-bot/internal/repository"
)

func main() {
	// 1. Загружаем конфиг
	cfg := config.LoadConfig()

	// 2. Подключаемся к базе
	db, err := sqlx.Connect("pgx", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	// Проверяем связь
	err = db.Ping()
	if err != nil {
		log.Fatalf("База недоступна: %v", err)
	}

	log.Println("Успех! Мы подключились к базе данных.")

	repo := repository.NewUserRepository(db)
	testUser := &domain.User{
		TelegramID: 1234345,
		Username:   "ks1ze",
	}

	err = repo.CreateUser(testUser)
	if err != nil {
		log.Printf("Не удалось создать пользователя: %v", err)
	} else {
		log.Println("Игрок успешно создан в базе данных!")
	}

	bot, err := telegram.NewBot("8526180490:AAFtWk99ro7nNeoAfAcspxkb6VJoEzOXrSQ", repo)
	err = bot.Start()
}
