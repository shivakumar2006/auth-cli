package session

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Storage interface {
	CreateSession(session *Session) error
	GetSessionByToken(token string) (*Session, error)
	DeleteSession(token string) error
}

var (
	ErrSessionNotFound = errors.New("session not found")
	ErrSessionExpired  = errors.New("session expired")
)

const (
	DefaultSessionDuration = 30 * time.Minute
)

type Manager struct {
	Storage Storage
}

func NewManager(storage Storage) *Manager {
	return &Manager{
		Storage: storage,
	}
}

func (m *Manager) CreateSession(userID int64) (*Session, error) {
	session := &Session{
		UserID:    userID,
		Token:     uuid.NewString(),
		CreatedAt: time.Now().UTC(),
		ExpiresAt: time.Now().UTC().Add(DefaultSessionDuration),
	}

	if err := m.Storage.CreateSession(session); err != nil {
		return nil, fmt.Errorf("failed to create session : %s", err.Error())
	}

	return session, nil
}

func (m *Manager) ValidateSession(token string) (*Session, error) {
	session, err := m.Storage.GetSessionByToken(token)
	if err != nil {
		return nil, err
	}

	if session.IsExpired() {
		_ = m.Storage.DeleteSession(token)
		return nil, ErrSessionExpired
	}

	return session, nil
}

func (m *Manager) DestroySession(token string) error {
	return m.Storage.DeleteSession(token)
}
