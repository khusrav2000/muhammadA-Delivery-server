package model

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()
	return &User{
		Login:    "userkhusrav",
		Password: "password",
	}
}
