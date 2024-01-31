package user

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/igorzash/project-zefir/db"
)

func GetByEmail(email string) *User {
	sqlStmt := `SELECT id, created_at, updated_at, email, nickname, password_hash FROM users WHERE email = ?`
	row := db.Conn.QueryRow(sqlStmt, email)

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

func HandleFind(c *gin.Context) {
	// Logic for getting all users
}

func HandleCreate(c *gin.Context) {
	// Logic for creating a new user
}

func HandleGetByID(c *gin.Context) {
	// Logic for getting a user by ID
}

func HandleUpdate(c *gin.Context) {
	// Logic for updating a user
}
