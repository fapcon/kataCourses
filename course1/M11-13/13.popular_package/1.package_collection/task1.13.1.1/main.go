package main

import (
	"strconv"
	"strings"
)

//go:generate grizzly generate main.go
//grizzly:generate
type User struct {
	Id   int
	Name string
	Age  int
}

func getUsersByCondition(users []*User, condition string) []*User {
	newUsers := NewUserCollection(users)
	// реализуйте функцию getUsersByCondition

	s := strings.Split(condition, " ")
	if len(s) == 1 {
		youngUsers := newUsers.Filter(func(user *User) bool {
			return user.Name == condition
		})
		return youngUsers.Items
	}
	a, _ := strconv.Atoi(s[len(s)-1])
	switch s[len(s)-2] {
	case ">":
		youngUsers := newUsers.Filter(func(user *User) bool {
			return user.Age > a
		})
		return youngUsers.Items
	case "<":
		youngUsers := newUsers.Filter(func(user *User) bool {
			return user.Age < a
		})
		return youngUsers.Items
	case ">=":
		youngUsers := newUsers.Filter(func(user *User) bool {
			return user.Age >= a
		})
		return youngUsers.Items
	case "<=":
		youngUsers := newUsers.Filter(func(user *User) bool {
			return user.Age <= a
		})
		return youngUsers.Items
	case "=":
		youngUsers := newUsers.Filter(func(user *User) bool {
			return user.Age == a
		})
		return youngUsers.Items
	}
	return nil
}

func getUsersByAge(users []*User, age int) []*User {
	newUsers := NewUserCollection(users)
	youngUsers := newUsers.Filter(func(user *User) bool {
		return user.Age == age
	})
	return youngUsers.Items
}

func getUsersByNickName(users []*User, nickName string) []*User {
	// реализуйте функцию getUsersByNickName
	newUsers := NewUserCollection(users)
	youngUsers := newUsers.Filter(func(user *User) bool {
		return user.Name == nickName
	})
	return youngUsers.Items
}

func getUsersUniqueNickName(users []*User) []*User {
	newUsers := NewUserCollection(users)
	uniqNames := newUsers.UniqByName()
	return uniqNames.Items
}
