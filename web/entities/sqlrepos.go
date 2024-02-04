package entities

import (
	"database/sql"
	"fmt"

	"github.com/igorzash/project-zefir/web/entities/followpkg"
	"github.com/igorzash/project-zefir/web/entities/userpkg"
)

func NewSQLRepositories(dbConn *sql.DB) (*Repositories, error) {
	userRepo, err := userpkg.NewSQLUserRepository(dbConn)
	if err != nil {
		return nil, fmt.Errorf("failed to create user repository: %w", err)
	}

	followRepo, err := followpkg.NewSQLFollowRepository(dbConn)
	if err != nil {
		return nil, fmt.Errorf("failed to create follow repository: %w", err)
	}

	return &Repositories{
		UserRepo:   userRepo,
		FollowRepo: followRepo,
	}, nil
}
