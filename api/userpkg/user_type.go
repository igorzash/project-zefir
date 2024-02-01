package userpkg

type User struct {
	ID           int    `db:"id" json:"id" fake:"-"`
	Email        string `db:"email" json:"-" fake:"{email}"`
	CreatedAt    string `db:"created_at" json:"createdAt" fake:"-"`
	UpdatedAt    string `db:"updated_at" json:"updatedAt" fake:"-"`
	Nickname     string `db:"nickname" json:"nickname" fake:"{username}"`
	PasswordHash string `db:"password_hash" json:"-" fake:"-"`
}
