package definitions

import "time"

type User struct {
	ID                uint64    `db:"id"`
	Username          string    `db:"username"`
	EncryptedPassword string    `db:"encrypted_password"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
	IsAdmin           bool      `db:"is_admin"`
}
