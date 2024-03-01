package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go clientReader(conn)

	for {
		scanner := bufio.NewReader(os.Stdin)
		s, err := scanner.ReadString('\n')

		_, err = conn.Write([]byte(s))
		if err != nil {
			fmt.Errorf("error sending message to geogrpc")
		}
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
	//wg.Wait()
}

// clientReader выводит на экран все сообщения от сервера
func clientReader(conn net.Conn) {
	for {
		g := make([]byte, 1024)
		//time.Sleep(100 * time.Millisecond)
		_, err := conn.Read(g)
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}
		res := string(g)
		if res != "" {
			fmt.Println(res)
		}
	}
}
