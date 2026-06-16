package auth

import "time"

// user struct represent the registered user with these feilds in the system
type User struct {
	ID             int64     `json:"id"`
	Username       string    `json:"username"`
	Passwordhash   string    `json:"_"`
	MFAEnabled     bool      `json:"mfa_enabled"`
	MFASecret      *string   `json:"_"`
	FailedAttempts int       `json:"failed_attempts"`
	LockedUntil    time.Time `json:"locked_until,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	LastLogin      time.Time `json:"last_login,omitempty"`
}

// public user is the user version that can be displayed actually
// in cli output without expose any sensitive data
type PublicUser struct {
	Username    string    `json:"username"`
	MFAEnabled  bool      `json:"mfa_enabled"`
	CreatedAt   time.Time `json:"created_at"`
	LastLogin   time.Time `json:"last_login,omitempty"`
	LockedUntil time.Time `json:"last_login,omitempty"`
}

// this function convert a user struct into a public user struct
func (u *User) ToPublic() (*PublicUser, error) {
	return &PublicUser{
		Username:    u.Username,
		MFAEnabled:  u.MFAEnabled,
		CreatedAt:   u.CreatedAt,
		LastLogin:   u.LastLogin,
		LockedUntil: u.LockedUntil,
	}, nil
}
