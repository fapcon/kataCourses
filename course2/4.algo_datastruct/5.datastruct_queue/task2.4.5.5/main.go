package main

import "fmt"

type CircuitRinger interface {
	Add(val int)
	Get() (int, bool)
}

type RingBuffer struct {
	buffer []int
	size   int
	head   int
	tail   int
	count  int
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buffer: make([]int, size),
		size:   size,
		head:   0,
		tail:   0,
		count:  0,
	}
}

func (rb *RingBuffer) Add(val int) {
	rb.buffer[rb.tail] = val
	rb.tail = (rb.tail + 1) % rb.size
	if rb.count < rb.size {
		rb.count++
	} else {
		rb.head = (rb.head + 1) % rb.size
	}
}

func (rb *RingBuffer) Get() (int, bool) {
	if rb.count > 0 {
		val := rb.buffer[rb.head]
		rb.head = (rb.head + 1) % rb.size
		rb.count--
		return val, true
	}
	return 0, false
}

func main() {
	rb := NewRingBuffer(3)
	rb.Add(1)
	rb.Add(2)
	rb.Add(3)
	rb.Add(4)

	if val, ok := rb.Get(); ok {
		fmt.Println(val) // Выводит: 2
	}
	if val, ok := rb.Get(); ok {
		fmt.Println(val) // Выводит: 3
	}
	if val, ok := rb.Get(); ok {
		fmt.Println(val) // Выводит: 4
	}
	if _, ok := rb.Get(); !ok {
		fmt.Println("Буфер пуст") // Выводит: Буфер пуст
	}
	rb.Add(4)
	rb.Add(4)
	rb.Add(4)
	fmt.Println(rb.buffer)
}
