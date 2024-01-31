package user

type User struct {
	Id           int    `db:"id"`
	Email        string `db:"email" json:"-"`
	CreatedAt    string `db:"created_at" json:"createdAt"`
	UpdatedAt    string `db:"updated_at" json:"updatedAt"`
	Nickname     string `db:"nickname"`
	PasswordHash string `db:"password_hash" json:"-"`
}
