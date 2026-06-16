package cli

import (
	"auth-cli/internal/auth"
	"auth-cli/internal/session"
)

type Context struct {
	LoggedIn       bool
	CurrentUser    *auth.User
	CurrentSession *session.Session
}

func NewContext() *Context {
	return &Context{
		LoggedIn:       false,
		CurrentUser:    nil,
		CurrentSession: nil,
	}
}

func (c *Context) Login(user *auth.User, session *session.Session) {
	c.LoggedIn = true
	c.CurrentUser = user
	c.CurrentSession = session
}

func (c *Context) Logout() {
	c.LoggedIn = false
	c.CurrentUser = nil
	c.CurrentSession = nil
}

func (c *Context) IsLoggedIn() bool {
	return c.LoggedIn
}
