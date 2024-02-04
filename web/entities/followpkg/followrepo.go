package followpkg

import (
	"database/sql"

	"github.com/igorzash/project-zefir/web/entities/userpkg"
)

type FollowRepository interface {
	Insert(follow *Follow) (sql.Result, error)
	GetByUsersIDs(followerID int, followeeID int) (*Follow, error)
	GetFollowState(followerID, followeeID int) (FollowState, error)
	GetUserFollowers(userID int, limit int, offset int) ([]*userpkg.User, error)
	GetUsersFollowedBy(userID int, limit int, offset int) ([]*userpkg.User, error)
	Delete(followerID int, followeeID int) (sql.Result, error)
}
