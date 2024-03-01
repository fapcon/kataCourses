package main

import (
	"fmt"
	"github.com/go-daq/crc8"
	"github.com/howeyc/crc16"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"time"
)

type bucketNode struct {
	key   string
	value interface{}
	next  *bucketNode
}

type bucket struct {
	head *bucketNode
}

type HashMap struct {
	hasher hash.Hash
	arr    []*bucket
}

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

func WithHashCRC64() func(hm *HashMap) {
	return func(hm *HashMap) {
		hm.hasher = crc64.New(crc64.MakeTable(crc64.ECMA))
	}
}

func WithHashCRC32() func(hm *HashMap) {
	return func(hm *HashMap) {
		hm.hasher = crc32.New(crc32.MakeTable(crc32.IEEE))
	}
}

func WithHashCRC16() func(hm *HashMap) {
	return func(hm *HashMap) {
		hm.hasher = crc16.New(crc16.MakeTable(crc16.IBM))
	}
}
func WithHashCRC8() func(hm *HashMap) {
	return func(hm *HashMap) {
		hm.hasher = crc8.New(crc8.MakeTable(crc8.Size))
	}
}

func NewHashMap(size int, opts ...func(hm *HashMap)) *HashMap {
	hm := &HashMap{arr: make([]*bucket, size)}
	for i := 0; i < size; i++ {
		hm.arr[i] = &bucket{}
	}
	for _, opt := range opts {
		opt(hm)
	}
	return hm
}

func MeassureTime(a func()) time.Duration {
	start := time.Now()
	a()
	duration := time.Since(start)
	return duration
}

func (hm *HashMap) Set(key string, value interface{}) {
	hash, err := hm.hasher.Write([]byte(key))
	if err != nil {
		fmt.Errorf("ERR IN SET FUNC")
	}
	index := hash % len(hm.arr)

	newNode := &bucketNode{
		key:   key,
		value: value,
		next:  nil,
	}

	if hm.arr[index].head == nil {
		hm.arr[index].head = newNode
	} else {
		currentNode := hm.arr[index].head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
}

func (hm *HashMap) Get(key string) (interface{}, bool) {
	hash, err := hm.hasher.Write([]byte(key))
	if err != nil {
		fmt.Errorf("ERR IN GET FUNC")
	}
	index := hash % len(hm.arr)

	currentNode := hm.arr[index].head
	for currentNode != nil {
		if currentNode.key == key {
			return currentNode.value, true
		}
		currentNode = currentNode.next
	}

	return nil, false
}

func main() {

	m := NewHashMap(16, WithHashCRC64())
	since := MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = NewHashMap(16, WithHashCRC32())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = NewHashMap(16, WithHashCRC16())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = NewHashMap(16, WithHashCRC8())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)
}
