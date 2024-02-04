package userpkg

import (
	"database/sql"
	"fmt"
)

type SQLUserRepository struct {
	dbConn             *sql.DB
	createUserStmt     *sql.Stmt
	getUserStmtByID    *sql.Stmt
	getUserStmtByEmail *sql.Stmt
	updateUserStmt     *sql.Stmt
}

func prepareGetUsersStmt(dbConn *sql.DB, column string) (*sql.Stmt, error) {
	return dbConn.Prepare(fmt.Sprintf(`SELECT id, created_at, updated_at, email, nickname, password_hash FROM users WHERE %s = ?`, column))
}

func NewSQLUserRepository(dbConn *sql.DB) (UserRepository, error) {
	repo := &SQLUserRepository{dbConn: dbConn}

	var err error
	repo.createUserStmt, err = dbConn.Prepare("INSERT INTO users (created_at, updated_at, nickname, email, password_hash) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare createUserStmt: %w", err)
	}

	repo.getUserStmtByID, err = prepareGetUsersStmt(dbConn, "id")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare getUserStmtByID: %w", err)
	}

	repo.getUserStmtByEmail, err = prepareGetUsersStmt(dbConn, "email")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare getUserStmtByEmail: %w", err)
	}

	repo.updateUserStmt, err = dbConn.Prepare("UPDATE users SET updated_at = ?, nickname = ?, email = ?, password_hash = ? WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare updateUserStmt: %w", err)
	}

	return repo, nil
}

func (repo *SQLUserRepository) Insert(user *User) (sql.Result, error) {
	result, err := repo.createUserStmt.Exec(user.CreatedAt, user.UpdatedAt, user.Nickname, user.Email, user.PasswordHash)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(id)
	return result, nil
}

func (repo *SQLUserRepository) getUser(stmt *sql.Stmt, arg interface{}) (*User, error) {
	row := stmt.QueryRow(arg)

	var user User
	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Email, &user.Nickname, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not found
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (repo *SQLUserRepository) GetByID(ID int) (*User, error) {
	return repo.getUser(repo.getUserStmtByID, ID)
}

func (repo *SQLUserRepository) GetByEmail(email string) (*User, error) {
	return repo.getUser(repo.getUserStmtByEmail, email)
}

func (repo *SQLUserRepository) Update(user *User) (sql.Result, error) {
	return repo.updateUserStmt.Exec(user.UpdatedAt, user.Nickname, user.Email, user.PasswordHash, user.ID)
}
