package storage

import "auth-cli/internal/auth"

type Storage interface {
	CreateUser(user *auth.User) error
	GetUserByUsername(username string) (*auth.User, error)
	UpdateUser(user *auth.User) error
	UpdateLastLogin(userID int64) error

	EnableMFA(userID int64, secret string) error
	DisableMFA(userID int64) error

	IncrementFailedAttempts(userID int64) error
	ResetFailedAttempts(userID int64) error
}
