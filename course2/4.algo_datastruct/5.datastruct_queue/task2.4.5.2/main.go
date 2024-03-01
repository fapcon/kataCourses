package main

import (
	"fmt"
)

type BrowserHistory struct {
	stack []string
	//history []string
}

func (h *BrowserHistory) Visit(url string) {
	fmt.Println("Посещение " + url)
	h.stack = append(h.stack, url)
	//h.history = append(h.history, url)

}

func (h *BrowserHistory) Back() {
	fmt.Println("Возврат к " + h.stack[len(h.stack)-1])
	if len(h.stack) == 0 {
		return
	}
	if len(h.stack) > 0 {
		h.stack = h.stack[:len(h.stack)-1]
	}
	//h.history = append(h.history, element)
}

func (h *BrowserHistory) PrintHistory() {
	//fmt.Println(h.history)
	if len(h.stack) == 0 {
		fmt.Println("Нет больше истории для возврата")
	}
	for i := len(h.stack) - 1; i >= 0; i-- {
		fmt.Println(h.stack[i])
	}
}

func main() {
	history := &BrowserHistory{}
	history.Visit("www.google.com")
	history.Visit("www.github.com")
	history.Visit("www.openai.com")
	history.Back()
	history.Back()
	history.Back()
	history.PrintHistory()
}
