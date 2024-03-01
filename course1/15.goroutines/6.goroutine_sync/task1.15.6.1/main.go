package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func waitGroupExample(goroutines ...func() string) string {
	var result string
	for i := 0; i < len(goroutines); i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			result = result + goroutines[n]() + "\n"
		}(i)
	}
	wg.Wait()
	return result
}

func main() {
	count := 1000
	goroutines := make([]func() string, count)

	for i := 0; i < count; i++ {
		j := i
		goroutines[j] = func() string {
			return fmt.Sprintf("goroutine %d", j)
		}
	}
	fmt.Println(waitGroupExample(goroutines...))
}
