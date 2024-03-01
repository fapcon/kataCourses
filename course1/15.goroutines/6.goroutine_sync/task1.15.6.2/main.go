package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mutex sync.Mutex
}

func main() {
	fmt.Println(concurrentSafeCounter())
}

func (c *Counter) Increment() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
	return c.value
}

func concurrentSafeCounter() int {
	counter := Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	return counter.value
}
