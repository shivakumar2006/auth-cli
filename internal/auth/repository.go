package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("username already exists")
)

type Database struct {
	Db *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{
		Db: db,
	}
}

func (d *Database) CreateUser(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := d.Db.QueryRowContext(ctx, `
		INSERT INTO users (username, password_hash, mfa_enabled, mfa_secret, failed_attempts, locked_until)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id, created_at
	`, user.Username,
		user.PasswordHash,
		user.MFAEnabled,
		user.MFASecret,
		user.FailedAttempts,
		user.LockedUntil).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create user : %s", err.Error())
	}
	return nil
}

func (d *Database) GetUserByUsername(username string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User

	err := d.Db.QueryRowContext(ctx, `
		SELECT id, username, password_hash, mfa_enabled, mfa_secret, failed_attempts, locked_until, created_at, last_login
		FROM users
		WHERE username = $1
	`, username).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.MFAEnabled,
		&user.MFASecret,
		&user.FailedAttempts,
		&user.LockedUntil,
		&user.CreatedAt,
		&user.LastLogin)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by username : %s", err.Error())
	}

	return &user, nil
}

func (d *Database) UpdateUser(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := d.Db.ExecContext(ctx, `
		UPDATE users 
		SET 
			mfa_enabled = $1,
			mfa_secret = $2,
			failed_attempts = $3,
			locked_until = $4
		WHERE id = $5
	`, user.MFAEnabled,
		user.MFASecret,
		user.FailedAttempts,
		user.LockedUntil,
		user.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update user : %s", err.Error())
	}
	return nil
}

func (d *Database) UpdateLastLogin(userID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now().UTC()

	_, err := d.Db.ExecContext(ctx, `
		UPDATE users
		SET last_login = $1
		WHERE id = $2
	`, now, userID)

	if err != nil {
		return fmt.Errorf("failed to update last login : %s", err.Error())
	}

	return nil
}
