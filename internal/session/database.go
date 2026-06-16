package session

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Database struct {
	Db *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{
		Db: db,
	}
}

func (d *Database) CreateSession(session *Session) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := d.Db.QueryRowContext(ctx, `
		INSERT INTO sessions (user_id, token, created_at, expires_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id 
	`, session.UserID, session.Token, session.CreatedAt, session.ExpiresAt).Scan(&session.ID)

	if err != nil {
		return fmt.Errorf("failed to create session : %w", err)
	}

	return nil
}

func (d *Database) GetSessionByToken(token string) (*Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var session Session

	err := d.Db.QueryRowContext(ctx, `
		SELECT id, user_id, token, created_at, expires_at
		FROM sessions 
		WHERE token = $1
	`, token).Scan(&session.ID, &session.UserID, &session.Token, &session.CreatedAt, &session.ExpiresAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrSessionNotFound
		}

		return nil, fmt.Errorf("failed to get session : %w", err)
	}

	return &session, nil
}

func (d *Database) DeleteSession(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := d.Db.ExecContext(ctx, `
		DELETE FROM sessions 
		WHERE token = $1
	`, token)

	if err != nil {
		return fmt.Errorf("failed to delete session : %w", err)
	}

	return nil
}
