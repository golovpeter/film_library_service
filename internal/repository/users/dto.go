package users

type UserDataIn struct {
	Username string
	Password string
}

type UserDataOut struct {
	ID           int64  `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
	Role         string `db:"role"`
}
