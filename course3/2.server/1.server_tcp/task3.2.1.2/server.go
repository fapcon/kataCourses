// Код сервера
package main

import (
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}

// handleConn обрабатывает входящие сообщения от клиента
func handleConnection(conn net.Conn) {
	buf := []byte("HTTP/1.1 200 OK\r\n<!DOCTYPE html>\n<html>\n<head>\n<title>Webserver</title>\n</head>\n<body>\nhello world\n</body>\n</html>")
	go clientWriter(conn, buf)
}

// clientWriter отправляет сообщения текущему клиенту
func clientWriter(conn net.Conn, buf []byte) {
	_, err := conn.Write(buf)
	if err != nil {
		fmt.Println("Error writing to client:", err)
		return
	}
}
