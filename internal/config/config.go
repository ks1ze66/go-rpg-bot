package config

import "os"

type Config struct {
	DatabaseURL string
	BotToken    string
}

// LoadConfig достает данные из переменных окружения
func LoadConfig() *Config {
	return &Config{
		// Если переменной нет, используем дефолтную строку подключения
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/rpg_bot?sslmode=disable"),
		BotToken:    os.Getenv("BOT_TOKEN"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
