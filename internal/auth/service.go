package auth

import (
	"auth-cli/internal/auth/storage"
	"errors"
	"fmt"
	"time"
)

var (
	ErrAccountLocked = errors.New("account is temporarily locked")
)

const (
	MaxFailedAttempts = 5
	LockoutDuration   = 15 * time.Minute
)

type Service struct {
	Storage storage.Storage
}

func NewService(storage storage.Storage) *Service {
	return &Service{
		Storage: storage,
	}
}

func (s *Service) Register(username string, password string) error {
	// validate password strength
	if err := ValidatePassword(password); err != nil {
		return fmt.Errorf("invalid password : %s", err.Error())
	}

	// check if the user is already exist or not
	_, err := s.Storage.GetUserByUsername(username)
	if err == nil {
		return ErrUserExists
	}

	if !errors.Is(err, ErrUserNotFound) {
		return fmt.Errorf("failed to check user existence : %s", err.Error())
	}

	// hash the password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password : %s", err.Error())
	}

	// create the user
	user := &User{
		Username:       username,
		PasswordHash:   hashedPassword,
		MFAEnabled:     false,
		MFASecret:      nil,
		FailedAttempts: 0,
		LockedUntil:    nil,
	}

	if err := s.Storage.CreateUser(user); err != nil {
		return fmt.Errorf("failed to create user : %s", err.Error())
	}

	return nil
}

func (s *Service) Login(username, password string) (*User, error) {
	user, err := s.Storage.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username : %s", err.Error())
	}

	// account locked
	if user.LockedUntil != nil && user.LockedUntil.After(time.Now()) {
		return nil, ErrAccountLocked
	}

	// validate password
	if err := ComparePassword(user.PasswordHash, password); err != nil {
		user.FailedAttempts++

		if user.FailedAttempts >= MaxFailedAttempts {
			lockUntil := time.Now().Add(LockoutDuration)

			user.LockedUntil = &lockUntil
		}

		if err := s.Storage.UpdateUser(user); err != nil {
			return nil, fmt.Errorf("failed to update user %s", err.Error())
		}

		return nil, ErrInvalidPassword
	}

	// reset failed attempts on successful login
	user.FailedAttempts = 0
	user.LockedUntil = nil

	if err := s.Storage.UpdateUser(user); err != nil {
		return nil, fmt.Errorf("failed to update user %s", err.Error())
	}

	// update last login time
	if err := s.Storage.UpdateLastLogin(user.ID); err != nil {
		return nil, fmt.Errorf("failed to update last login %s", err.Error())
	}

	return user, nil
}

// Enable2FA enables MFA for a user.
func (s *Service) Enable2FA(userID int64, secret string) error {
	return s.Storage.EnableMFA(
		userID,
		secret,
	)
}

// Disable2FA disables MFA for a user.
func (s *Service) Disable2FA(userID int64) error {
	return s.Storage.DisableMFA(userID)
}
