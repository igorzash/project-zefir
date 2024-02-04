package entities

import (
	"database/sql"
	"fmt"

	"github.com/igorzash/project-zefir/web/entities/followpkg"
	"github.com/igorzash/project-zefir/web/entities/userpkg"
)

type Repositories struct {
	UserRepo   *userpkg.UserRepository
	FollowRepo *followpkg.FollowRepository
}

func NewRepositories(dbConn *sql.DB) (*Repositories, error) {
	userRepo, err := userpkg.NewUserRepository(dbConn)
	if err != nil {
		return nil, fmt.Errorf("failed to create user repository: %w", err)
	}

	followRepo, err := followpkg.NewFollowRepository(dbConn)
	if err != nil {
		return nil, fmt.Errorf("failed to create follow repository: %w", err)
	}

	return &Repositories{
		UserRepo:   userRepo,
		FollowRepo: followRepo,
	}, nil
}
