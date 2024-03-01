package main

import (
	"fmt"
)

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

type UserOption func(*User)

func WithUsername(value string) UserOption {
	return func(user *User) {
		user.Username = value
	}
}
func WithEmail(value string) UserOption {
	return func(user *User) {
		user.Email = value
	}
}

func WithRole(value string) UserOption {
	return func(user *User) {
		user.Role = value
	}
}

func NewUser(id int, options ...UserOption) *User {
	user := &User{ID: id}
	for _, option := range options {
		option(user)
	}
	return user
}
func main() {
	user := NewUser(1, WithUsername("testuser"), WithEmail("testuser@example.com"), WithRole("admin"))
	fmt.Printf("User: %+v\n", user)
}
