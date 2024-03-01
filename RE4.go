// 1. Задача  Написать функцию, которая принимает канал и число N типа int, после получения N значений из канала, функция должна вернуть срез с этими значениями

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func RandNumbers(length, max int) []int {
	var s []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		s = append(s, rand.Intn(max))
	}
	return s
}

func writeToChan(ch chan<- int) {
	defer close(ch)
	for _, v := range RandNumbers(100, 100) {
		time.Sleep(100 * time.Millisecond)
		ch <- v
	}
}

func chSlice(ch <-chan int, n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		value := <-ch
		s = append(s, value)
	}
	return s
}

func IsClosed(ch <-chan int) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}

func mergeChan(ch ...chan int) <-chan int {

	res := make(chan int, 10)

	go func() {

		var wg sync.WaitGroup
		wg.Add(len(ch))
		for _, c := range ch {
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

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	mergedChan := mergeChan(ch1, ch2, ch3, ch4)
	go writeToChan(ch1)
	go writeToChan(ch2)
	go writeToChan(ch3)
	go writeToChan(ch4)

	N := 10

	mergedSlice := chSlice(mergedChan, N)
	fmt.Println(mergedSlice)

	for {
		value, ok := <-mergedChan
		if !ok {
			break
		}
		fmt.Println(value)
	}
}
