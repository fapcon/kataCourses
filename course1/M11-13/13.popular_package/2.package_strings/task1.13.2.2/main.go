package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name     string
	Comments []Comment
}

type Comment struct {
	Message string
}

func main() {
	users := []User{
		{
			Name: "Betty",
			Comments: []Comment{
				{Message: "good Comment 1"},
				{Message: "BaD CoMmEnT 2"},
				{Message: "Bad Comment 3"},
				{Message: "Use camelCase please"},
			},
		},
		{
			Name: "Jhon",
			Comments: []Comment{
				{Message: "Good Comment 1"},
				{Message: "Good Comment 2"},
				{Message: "Good Comment 3"},
				{Message: "Bad Comments 4"},
			},
		},
	}
	fmt.Println(GetBadComments(users))
	fmt.Println(GetGoodComments(users))
	users = FilterComments(users)
	fmt.Println(users)
}

func FilterComments(users []User) []User {
	for i := 0; i < len(users); i++ {
		for j := 0; j < len(users[i].Comments); j++ {
			if IsBadComment(users[i].Comments[j].Message) {
				users[i].Comments = append(users[i].Comments[:j], users[i].Comments[j+1:]...)
			}
		}
	}
	for i := 0; i < len(users); i++ {
		for j := 0; j < len(users[i].Comments); j++ {
			if IsBadComment(users[i].Comments[j].Message) {
				users[i].Comments = append(users[i].Comments[:j], users[i].Comments[j+1:]...)
			}
		}
	}
	return users
}

func IsBadComment(comment string) bool {
	r := strings.ToLower(comment)
	if strings.Contains(r, "bad comment") {
		return true
	}
	return false
}

func GetBadComments(users []User) []Comment {
	var res []Comment
	for i := 0; i < len(users); i++ {
		for j := 0; j < len(users[i].Comments); j++ {
			if IsBadComment(users[i].Comments[j].Message) {
				res = append(res, users[i].Comments[j])
			}
		}
	}
	return res
}

func GetGoodComments(users []User) []Comment {
	var res []Comment
	for i := 0; i < len(users); i++ {
		for j := 0; j < len(users[i].Comments); j++ {
			if !IsBadComment(users[i].Comments[j].Message) {
				res = append(res, users[i].Comments[j])
			}
		}
	}
	return res
}
