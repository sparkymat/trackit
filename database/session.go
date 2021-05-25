package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sparkymat/trackit/database/definitions"
)

type Session interface {
	Find(ctx context.Context, id uint64) (*definitions.Session, error)
}

func NewSession(conn *sqlx.DB) Session {
	return &sessionDB{
		conn: conn,
	}
}

type sessionDB struct {
	conn *sqlx.DB
}

func (u *sessionDB) Find(ctx context.Context, id uint64) (*definitions.Session, error) {
	rows, err := u.conn.Queryx("SELECT * FROM sessions WHERE id=$1", id)

	if err != nil {
		return nil, err
	}

	sessions := []definitions.Session{}

	for rows.Next() {
		var session definitions.Session
		err = rows.StructScan(&session)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}

	if len(sessions) == 0 {
		return nil, ErrNotFound
	}

	firstSession := sessions[0]
	return &firstSession, nil
}
