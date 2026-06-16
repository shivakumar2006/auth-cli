package session

import "time"

// this session struct represents an authenticated user session
type Session struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// this function checks whether the session is expired
func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}
