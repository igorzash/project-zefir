package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func GetUserByEmail(email string) *User {
	sqlStmt := `SELECT id, created_at, updated_at, email, nickname, password_hash FROM users WHERE email = ?`
	row := DB.QueryRow(sqlStmt, email)

	var user User
	err := row.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Email, &user.Nickname, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not found
			return nil
		} else {
			log.Fatal(err)
		}
	}

	return &user
}

func HandleGetUsers(c *gin.Context) {
	// Logic for getting all users
}

func HandleCreateUser(c *gin.Context) {
	// Logic for creating a new user
}

func HandleGetUserByID(c *gin.Context) {
	// Logic for getting a user by ID
}

func HandleUpdateUser(c *gin.Context) {
	// Logic for updating a user
}
