package auth

import "time"

// user struct represent the registered user with these feilds in the system
// use *string for MFASecret instead of string.
// That maps cleanly to PostgreSQL NULL values and avoids ambiguity between "2FA not configured" and "empty string"
type User struct {
	ID             int64      `json:"id"`
	Username       string     `json:"username"`
	PasswordHash   string     `json:"-"`
	MFAEnabled     bool       `json:"mfa_enabled"`
	MFASecret      *string    `json:"-"`
	FailedAttempts int        `json:"failed_attempts"`
	LockedUntil    *time.Time `json:"locked_until,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	LastLogin      *time.Time `json:"last_login,omitempty"`
}

// public user is the user version that can be displayed actually
// in cli output without expose any sensitive data
type PublicUser struct {
	Username    string     `json:"username"`
	MFAEnabled  bool       `json:"mfa_enabled"`
	CreatedAt   time.Time  `json:"created_at"`
	LastLogin   *time.Time `json:"last_login,omitempty"`
	LockedUntil *time.Time `json:"locked_until,omitempty"`
}

// this function convert a user struct into a public user struct
func (u *User) ToPublic() *PublicUser {
	return &PublicUser{
		Username:    u.Username,
		MFAEnabled:  u.MFAEnabled,
		CreatedAt:   u.CreatedAt,
		LastLogin:   u.LastLogin,
		LockedUntil: u.LockedUntil,
	}
}
