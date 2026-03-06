package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ks1ze66/go-rpg-bot/internal/domain"
)

// UserRepository - это структура, которая "владеет" соединением с базой
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository - конструктор. Мы даем ему базу, он возвращает репозиторий.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser записывает нового игрока в таблицу
func (r *UserRepository) CreateUser(user *domain.User) error {
	// ON CONFLICT - это магия Postgres.
	// Если tg_id совпал, мы просто обновляем username и время захода.
	query := `
        INSERT INTO users (tg_id, username) 
        VALUES ($1, $2)
        ON CONFLICT (tg_id) DO UPDATE SET 
            username = EXCLUDED.username,
            created_at = NOW()`

	_, err := r.db.Exec(query, user.TelegramID, user.Username)
	return err
}
