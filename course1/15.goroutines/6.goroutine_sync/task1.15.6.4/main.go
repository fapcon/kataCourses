package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type User struct {
	ID   int
	Name string
}

type Cache struct {
	users []User
	mutex sync.RWMutex
}

func NewCache() *Cache {
	cache := Cache{}
	return &cache
}

func (c *Cache) Set(key string, user *User) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	newUser := User{
		ID:   100,
		Name: key,
	}
	c.users = append(c.users, newUser)
}

func (c *Cache) Get(key string) *User {
	var res User
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	for i := range c.users {
		if c.users[i].Name == key {
			res = c.users[i]
		}
	}
	return &res
}

func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

func main() {
	cache := NewCache()

	for i := 0; i < 100; i++ {
		go cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
			ID:   i,
			Name: fmt.Sprint("user-", i),
		})
	}

	time.Sleep(1 * time.Second)

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(cache.Get(keyBuilder("user", strconv.Itoa(i))))
		}(i)
	}
}
