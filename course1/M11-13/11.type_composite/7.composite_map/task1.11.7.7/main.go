package main

import "fmt"

func main() {
	people := []User{
		{Nickname: "Алексей", Age: 25, Email: "1"},
		{Nickname: "Елена", Age: 31, Email: "2"},
		{Nickname: "Алексей", Age: 27, Email: "3"},
		{Nickname: "Иван", Age: 35, Email: "4"},
		{Nickname: "Елена", Age: 30, Email: "5"},
	}

	fmt.Println(getUniqueUsers(people))
}

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUsers(users []User) []User {

	// Твой код для получения уникальных пользователей по никнейму

	var nick, nickunique []string
	for i := 0; i < len(users); i++ {
		nick = append(nick, users[i].Nickname)
	}
	fmt.Println(nick)
	lenmap := make(map[string]int)
	for _, name := range nick {
		lenmap[name]++
		if lenmap[name] == 1 {
			nickunique = append(nickunique, name)
		}
	}
	fmt.Println(nickunique)
	uniqueUsers := make([]User, len(nickunique), len(nickunique))
	idx := 0
	for i := range users {
		for j := idx; j < len(nickunique); j++ {
			if users[i].Nickname == nickunique[j] {
				uniqueUsers[j] = users[i]
				idx++
			}
		}
	}
	return uniqueUsers
}
