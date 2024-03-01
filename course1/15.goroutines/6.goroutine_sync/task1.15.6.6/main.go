package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	ID   int
	Name string
}

type Cache struct {
	users []User
	data  sync.Map
}

func NewCache() *Cache {
	cache := Cache{}
	return &cache
}

func (c *Cache) Set(key string, value interface{}) {
	c.data.Store(key, value)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.data.Load(key)
}

func keyBuilder(keys ...string) string {
	return strings.Join(keys, "-")
}

func GetUser(i interface{}) *User {
	return i.(*User)
}

func main() {

	cache := NewCache()

	for i := 0; i < 100; i++ {
		go cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
			ID:   i,
			Name: fmt.Sprint("user-", i),
		})
	}

	for i := 0; i < 100; i++ {
		go func(i int) {
			raw, _ := cache.Get(keyBuilder("user", strconv.Itoa(i)))
			fmt.Println(GetUser(raw))
		}(i)
	}

}
