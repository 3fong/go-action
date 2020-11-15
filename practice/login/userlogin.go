package login

import (
	"errors"
	"fmt"
)

// User user type for login
type User struct {
	Name     string
	Password string
}

// Userlogin userlogin
func (u User) Userlogin() string {
	return fmt.Sprintf("%s---%s", u.Name, u.Password)
}

// UserLogout userLogout
func (u User) UserLogout(token string) error {
	if token == "" {
		return errors.New("token is nil")
	}
	return nil
}
