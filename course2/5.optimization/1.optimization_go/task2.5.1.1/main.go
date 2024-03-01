package main

import (
	"fmt"
	"sync"
)

type HashMap struct {
	data   *sync.Map
	casheS []Pair
	casheL *Node
}

type Node struct {
	Key   string
	Value interface{}
	Next  *Node
}

type Pair struct {
	key   string
	value interface{}
}

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

func (h *HashMap) Set(key string, value interface{}) {
	h.data.Store(key, value)
}
func (h *HashMap) Get(key string) (interface{}, bool) {
	v, ok := h.data.Load(key)
	return v, ok
}

func (h *HashMap) CasheSlice() []Pair {
	h.casheS = []Pair{}
	h.data.Range(func(key, value interface{}) bool {
		h.casheS = append(h.casheS, Pair{key: key.(string),
			value: value,
		})
		return true
	})
	return h.casheS
}
func (h *HashMap) GetCahseSliceValue(key string) (interface{}, bool) {
	for _, kv := range h.casheS {
		if kv.key == key {
			return kv.value, true
		}
	}
	return nil, false
}

func (h *HashMap) CasheList() *Node {
	if h.casheL != nil {
		return h.casheL
	}
	var head *Node
	var tail *Node
	h.data.Range(func(key, value interface{}) bool {
		node := &Node{Key: key.(string),
			Value: value}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
		return true
	})
	h.casheL = head
	return h.casheL
}
func (h *HashMap) GetCasheListValue(key string) (interface{}, bool) {
	node := h.casheL
	for node != nil {
		if node.Key == key {
			return node.Value, true
		}
		node = node.Next
	}
	return nil, false
}
func NewHashMap() *HashMap {
	h := &HashMap{
		data: &sync.Map{}}
	return h
}
func main() {
	m := NewHashMap()
	m.Set("key1", "value1")
	m.Set("key2", "value2")
	if value, ok := m.Get("key1"); ok {
		fmt.Println("Key1:", value)
	} else {
		fmt.Println("Key1 not found")
	}
	if value, ok := m.Get("key2"); ok {
		fmt.Println("Key2:", value)
	} else {
		fmt.Println("Key2 not found")
	}
	if value, ok := m.Get("key3"); ok {
		fmt.Println("Key3:", value)
	} else {
		fmt.Println("Key3 not found")
	}
}
