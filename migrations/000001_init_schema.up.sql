CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    tg_id BIGINT UNIQUE NOT NULL,
    username TEXT,
    level INTEGER DEFAULT 1,
    experience INTEGER DEFAULT 0,
    current_quest TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);