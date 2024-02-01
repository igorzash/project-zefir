package followpkg

type Follow struct {
	FollowerID int    `db:"follower_id" json:"followerId"`
	FolloweeID int    `db:"followee_id" json:"followeeId"`
	CreatedAt  string `db:"created_at" json:"createdAt"`
	UpdatedAt  string `db:"updated_at" json:"updatedAt"`
}

type FollowState string

const (
	NotFollowing FollowState = "not_following"
	Following    FollowState = "following"
	Mutual       FollowState = "mutual"
)
