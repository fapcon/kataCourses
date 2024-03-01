package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем новый тикер с интервалом 1 секунда
	ticker := time.NewTicker(1 * time.Second)

	data := NotifyEvery(ticker, 5*time.Second, "Таймер сработал")

	for v := range data {
		fmt.Println(v)
	}

	fmt.Println("Программа завершена")
}

func NotifyEvery(ticker *time.Ticker, d time.Duration, message string) <-chan string {
	res := make(chan string, 1000000)
	//done := make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:

				res <- message
			}
		}
	}()
	time.Sleep(d)
	ticker.Stop()
	time.Sleep(2 * time.Second)
	close(res)
	return res
}
