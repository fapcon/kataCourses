package main

import (
	"sync"
)

func main() {

}

func generateChan(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func mergeChan(mergeTo chan int, from ...chan int) {
	/*
		go func() {
			mergeTo = generateChan(1)
		}()

		go func() {
			for i := range from {
				from[i] = generateChan(i)
			}
		}()
		for i := range from {
			go func() {
				for j := range from[i] {
					mergeTo <- j
				}
			}()
			time.Sleep(time.Second)
		}

	*/
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(from))
		for _, c := range from {
			go func(c chan int) {
				for j := range c {
					mergeTo <- j
				}
				wg.Done()
			}(c)
		}
		wg.Wait()
		close(mergeTo)
	}()
}

func mergeChan2(chans ...chan int) chan int {
	res := make(chan int)
	/*
		go func() {
			res = generateChan(1)
		}()

		go func() {
			for i := range chans {
				chans[i] = generateChan(i)
			}
		}()

	*/
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(chans))
		for _, c := range chans {
			go func(c chan int) {
				for j := range c {
					res <- j
				}
				wg.Done()
			}(c)
		}
		wg.Wait()
		close(res)
	}()
	return res
}
