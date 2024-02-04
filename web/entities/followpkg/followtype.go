package followpkg

import "time"

type Follow struct {
	FollowerID int    `db:"follower_id" json:"followerId"`
	FolloweeID int    `db:"followee_id" json:"followeeId"`
	CreatedAt  string `db:"created_at" json:"createdAt"`
	UpdatedAt  string `db:"updated_at" json:"updatedAt"`
}

func NewFollow(followerID int, followeeID int) *Follow {
	currentTime := time.Now().Format(time.RFC3339)
	return &Follow{
		FollowerID: followerID,
		FolloweeID: followeeID,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
	}
}

type FollowState string

const (
	NotFollowing FollowState = "not_following"
	Following    FollowState = "following"
	Mutual       FollowState = "mutual"
)
