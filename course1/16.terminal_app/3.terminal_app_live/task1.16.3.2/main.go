package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/gosuri/uilive"
)

func main() {
	writer := uilive.New()
	writer.Start()
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	_, _ = fmt.Fprintf(writer.Newline(), "My menu:\n")
	_, _ = fmt.Fprintf(writer.Newline(), "1. Submenu 1\n")
	_, _ = fmt.Fprintf(writer.Newline(), "2. Submenu 2\n")
	_, _ = fmt.Fprintf(writer.Newline(), "Press q to quit\n")
	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		//fmt.Printf("You pressed: rune %q, key %X\r\n", event.Rune, event.Key)
		if event.Rune == '1' {
			_, _ = fmt.Fprintf(writer.Newline(), "Submenu 1\n")
			_, _ = fmt.Fprintf(writer.Newline(), "Content submenu 1\n")
		}
		if event.Rune == '2' {
			_, _ = fmt.Fprintf(writer.Newline(), "Submenu 2\n")
			_, _ = fmt.Fprintf(writer.Newline(), "Content submenu 2\n")
		}
		if event.Rune == '\x00' {
			_, _ = fmt.Fprintf(writer.Newline(), "My menu:\n")
			_, _ = fmt.Fprintf(writer.Newline(), "1. Submenu 1\n")
			_, _ = fmt.Fprintf(writer.Newline(), "2. Submenu 2\n")
			_, _ = fmt.Fprintf(writer.Newline(), "Press q to quit\n")
		}
		if event.Rune == 'q' {
			break
		}
	}
}

//fmt.Printf("You pressed: %q\r\n", char)

/*
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	writer := uilive.New()
	writer.Start()
	ch := make(chan rune, 2)
	for i := 0; i <= 100; i++ {
		go func() {
			ch <- char
		}()
*/
/*
        case ch <- '2':
			_, _ = fmt.Fprintf(writer.Newline(), "Submenu 2\n")
			_, _ = fmt.Fprintf(writer.Newline(), "Content submenu 2\n")
		case ch <- '\x00':
			_, _ = fmt.Fprintf(writer.Newline(), "My menu:\n")
			_, _ = fmt.Fprintf(writer.Newline(), "1. Submenu 1\n")
			_, _ = fmt.Fprintf(writer.Newline(), "2. Submenu 2\n")
			_, _ = fmt.Fprintf(writer.Newline(), "Press q to quit\n")
		case ch <- 'q':
			break
	time.Sleep(time.Millisecond * 200)
	if char == '1' {
		writer.Stop()
		writer.Start()
		_, _ = fmt.Fprintf(writer.Newline(), "Submenu 1\n")
		_, _ = fmt.Fprintf(writer.Newline(), "Content submenu 1\n")
	}
	if char == '2' {
		writer.Stop()
		writer.Start()
		_, _ = fmt.Fprintf(writer.Newline(), "Submenu 2\n")
		_, _ = fmt.Fprintf(writer.Newline(), "Content submenu 2\n")
	}
*/
