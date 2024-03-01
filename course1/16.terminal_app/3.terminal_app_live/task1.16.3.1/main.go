package main

import (
	"fmt"
	"github.com/gosuri/uilive"
	"time"
)

func main() {
	writer := uilive.New()
	writer.Start()
	for i := 0; i <= 100; i++ {
		now := time.Now()
		_, _ = fmt.Fprintf(writer.Newline(), "Текущее время: %v\n", now.Format("15:04:05"))
		_, _ = fmt.Fprintf(writer.Newline(), "Текущая дата: %v\n", now.Format("2006-01-02"))
		time.Sleep(time.Second * 1)
	}
}
