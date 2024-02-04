package userpkg

import (
	"database/sql"
)

type UserRepository interface {
	Insert(user *User) (sql.Result, error)
	GetByID(ID int) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) (sql.Result, error)
}
