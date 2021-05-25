package database

import (
	"context"
	"errors"

	"github.com/sparkymat/trackit/database/definitions"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	ErrNotFound = errors.New("record not found")
)

type User interface {
	Find(ctx context.Context, id uint64) (*definitions.User, error)
}

func NewUser(conn *sqlx.DB) User {
	return &userDB{
		conn: conn,
	}
}

type userDB struct {
	conn *sqlx.DB
}

func (u *userDB) Find(ctx context.Context, id uint64) (*definitions.User, error) {
	rows, err := u.conn.Queryx("SELECT * FROM users WHERE id=$1", id)

	if err != nil {
		return nil, err
	}

	users := []definitions.User{}

	for rows.Next() {
		var user definitions.User
		err = rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, ErrNotFound
	}

	firstUser := users[0]
	return &firstUser, nil
}
