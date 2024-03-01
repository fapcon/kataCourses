// Код сервера
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

type client struct {
	conn net.Conn
	name string
	ch   chan<- string
}

var (
	// канал для всех входящих клиентов
	entering = make(chan client)
	// канал для сообщения о выходе клиента
	leaving = make(chan client)
	// канал для всех сообщений
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// broadcaster рассылает входящие сообщения всем клиентам
// следит за подключением и отключением клиентов
func broadcaster() {
	// здесь хранятся все подключенные клиенты
	clients := make(map[client]bool)
	for {
		select {
		case cli := <-entering:
			// New client has entered
			clients[cli] = true
		case cli := <-leaving:
			// Client has left
			delete(clients, cli)
			close(cli.ch)
		case msg := <-messages:
			// Broadcast message to all clients
			for cli := range clients {
				cli.ch <- msg
			}
		}
	}
}

// handleConn обрабатывает входящие сообщения от клиента
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{conn, who, ch}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		fmt.Println(who + ": " + input.Text())
		messages <- who + ": " + input.Text()
	}

	//fmt.Println(who + ": " + input.Text())

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

// clientWriter отправляет сообщения текущему клиенту
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		// Write the message to the client
		_, err := conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}
		// Flush the writer to ensure the message is sent immediately
		err = conn.(*net.TCPConn).SetWriteDeadline(time.Now().Add(1 * time.Second))
		if err != nil {
			fmt.Println("Error setting write deadline:", err)
			return
		}
		err = conn.(*net.TCPConn).SetWriteDeadline(time.Time{})
		if err != nil {
			fmt.Println("Error clearing write deadline:", err)
			return
		}
	}
}
