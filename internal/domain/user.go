package domain

import "time"

// User представляет игрока в нашей системе.
// Мы используем теги `db`, чтобы библиотека sqlx понимала, куда мапить поля.
type User struct {
	ID           int64     `db:"id"`
	TelegramID   int64     `db:"tg_id"`
	Username     string    `db:"username"`
	Level        int       `db:"level"`
	Experience   int       `db:"experience"`
	CurrentQuest string    `db:"current_quest"`
	CreatedAt    time.Time `db:"created_at"`
}

// Помнишь, мы говорили про прогресс?
// Давай создадим метод, который рассчитывает опыт до следующего уровня.
func (u *User) GetRequiredExperience() int {
	return u.Level * 100 // Упрощенная формула для начала
}
