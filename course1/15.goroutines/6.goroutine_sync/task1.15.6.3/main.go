package main

import (
	"sync"
)

type Counter struct {
	count int64
	mutex sync.RWMutex
}

func main() {
	concurrentSafeCounter()
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
}
func (c *Counter) GetCount() int64 {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.count
}

func concurrentSafeCounter() {
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
}
