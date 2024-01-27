CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at TEXT,
    updated_at TEXT,
    nickname VARCHAR(255),
    email VARCHAR(255),
    password_hash VARCHAR(255)
);
