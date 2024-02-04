package followpkg

import (
	"database/sql"
	"fmt"

	"github.com/igorzash/project-zefir/web/entities/userpkg"
)

type SQLFollowRepository struct {
	dbConn                 *sql.DB
	insertFollowStmt       *sql.Stmt
	getByUsersIDsStmt      *sql.Stmt
	getFollowStateStmt     *sql.Stmt
	getUsersFollowedByStmt *sql.Stmt
	getUserFollowersStmt   *sql.Stmt
	deleteFollowStmt       *sql.Stmt
}

func prepareGetUsersStmt(dbConn *sql.DB, column1 string, column2 string) (*sql.Stmt, error) {
	sqlStr := fmt.Sprintf(`
        SELECT users.id, users.created_at, users.updated_at, users.email, users.nickname
        FROM users
        INNER JOIN follows ON follows.%s = users.id
        WHERE follows.%s = ?
        LIMIT ? OFFSET ?
    `, column1, column2)

	stmt, err := dbConn.Prepare(sqlStr)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare getUsersStmt: %w", err)
	}

	return stmt, nil
}

func NewSQLFollowRepository(dbConn *sql.DB) (FollowRepository, error) {
	repo := &SQLFollowRepository{dbConn: dbConn}

	var err error
	repo.insertFollowStmt, err = dbConn.Prepare("INSERT INTO follows (follower_id, followee_id, created_at, updated_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare insertFollowStmt: %w", err)
	}

	repo.getByUsersIDsStmt, err = dbConn.Prepare(`SELECT follower_id, followee_id, created_at, updated_at FROM follows WHERE follower_id = ? AND followee_id = ?`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare getByUsersIDsStmt: %w", err)
	}

	repo.getFollowStateStmt, err = dbConn.Prepare(`SELECT follower_id FROM follows WHERE (follower_id = ? AND followee_id = ?) OR (follower_id = ? AND followee_id = ?)`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare getFollowStateStmt: %w", err)
	}

	repo.getUsersFollowedByStmt, err = prepareGetUsersStmt(dbConn, "followee_id", "follower_id")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare getUsersStmt: %w", err)
	}

	repo.getUserFollowersStmt, err = prepareGetUsersStmt(dbConn, "follower_id", "followee_id")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare getUsersStmt: %w", err)
	}

	repo.deleteFollowStmt, err = dbConn.Prepare(`DELETE FROM follows WHERE follower_id = ? AND followee_id = ?`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare deleteFollowStmt: %w", err)
	}

	return repo, nil
}

func (repo *SQLFollowRepository) Insert(follow *Follow) (sql.Result, error) {
	return repo.insertFollowStmt.Exec(follow.FollowerID, follow.FolloweeID, follow.CreatedAt, follow.UpdatedAt)
}

func (repo *SQLFollowRepository) GetByUsersIDs(followerID int, followeeID int) (*Follow, error) {
	row := repo.getByUsersIDsStmt.QueryRow(followerID, followeeID)

	var follow Follow
	err := row.Scan(&follow.FollowerID, &follow.FolloweeID, &follow.CreatedAt, &follow.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// Follow not found
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &follow, nil
}

func (repo *SQLFollowRepository) GetFollowState(followerID, followeeID int) (FollowState, error) {
	rows, err := repo.getFollowStateStmt.Query(followerID, followeeID, followeeID, followerID)
	if err != nil {
		return NotFollowing, err
	}
	defer rows.Close()

	var count int
	var lastFollowerID int
	for rows.Next() {
		err := rows.Scan(&lastFollowerID)
		if err != nil {
			return NotFollowing, err
		}
		count++
	}

	switch count {
	case 2:
		return Mutual, nil
	case 1:
		if lastFollowerID == followerID {
			return Following, nil
		}
		fallthrough
	default:
		return NotFollowing, nil
	}
}

func getUsers(stmt *sql.Stmt, userID int, limit int, offset int) ([]*userpkg.User, error) {
	rows, err := stmt.Query(userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*userpkg.User{}
	for rows.Next() {
		user := userpkg.User{}
		err := rows.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Email, &user.Nickname)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (repo *SQLFollowRepository) GetUserFollowers(userID int, limit int, offset int) ([]*userpkg.User, error) {
	return getUsers(repo.getUserFollowersStmt, userID, limit, offset)
}

func (repo *SQLFollowRepository) GetUsersFollowedBy(userID int, limit int, offset int) ([]*userpkg.User, error) {
	return getUsers(repo.getUsersFollowedByStmt, userID, limit, offset)
}

func (repo *SQLFollowRepository) Delete(followerID int, followeeID int) (sql.Result, error) {
	return repo.deleteFollowStmt.Exec(followerID, followeeID)
}
