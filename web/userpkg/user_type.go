package userpkg

import (
	"time"

	"github.com/igorzash/project-zefir/web/helpers"
)

type User struct {
	ID           int    `db:"id" json:"id"`
	Email        string `db:"email" json:"-"`
	CreatedAt    string `db:"created_at" json:"createdAt"`
	UpdatedAt    string `db:"updated_at" json:"updatedAt"`
	Nickname     string `db:"nickname" json:"nickname"`
	PasswordHash string `db:"password_hash" json:"-"`
}

func NewUser(email string, nickname string, password string) (*User, error) {
	currentTime := time.Now().Format(time.RFC3339)
	hashedPassword, err := helpers.HashPassword(password)

	if err != nil {
		return nil, err
	}

	return &User{
		Email:        email,
		Nickname:     nickname,
		PasswordHash: hashedPassword,
		CreatedAt:    currentTime,
		UpdatedAt:    currentTime,
	}, nil
}
