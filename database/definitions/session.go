package definitions

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Session struct {
	ID        uint64    `db:"id"`
	UserID    uint64    `db:"user_id"`
	UUID      uuid.UUID `db:"uuid"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	ExpiresAt time.Time `db:"expires_at"`
	User      User
}
