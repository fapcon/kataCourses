package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://httpbin.org/get"
	parallelRequest := 5
	requestCount := 50
	result := benchRequest(url, parallelRequest, requestCount)

	for i := 0; i < requestCount; i++ {
		statusCode := <-result
		if statusCode != 200 {
			panic(fmt.Sprintf("Ошибка при отправке запроса: %d", statusCode))
		}
	}
	fmt.Println("Все горутины завершили работу")
}

func benchRequest(url string, parallelRequest int, requestCount int) <-chan int {
	ch := make(chan struct{}, parallelRequest)
	res := make(chan int, requestCount)
	for i := 1; i <= requestCount; i++ {
		ch <- struct{}{} // Занятие места в канале
		r, _ := httpRequest(url)
		res <- r
		<-ch
	}
	close(ch)
	close(res)
	return res
}

func httpRequest(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
