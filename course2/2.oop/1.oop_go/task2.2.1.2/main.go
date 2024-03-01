package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Logger interface {
	Log(a string) error
}

type ConsoleLogger struct {
	Writer io.Writer
}

type FileLogger struct {
	fileName string
}

func (f *FileLogger) Log(s string) error {
	var file *os.File
	_, err := os.Stat("logging.txt")
	if os.IsExist(err) {
		file, err = os.Create("logging.txt")
		if err != nil {
			fmt.Errorf("oshibka", err)
		}
	}
	writer := bufio.NewWriter(file)
	writer.Write([]byte(s))
	writer.Flush()
	return err
}

func (c *ConsoleLogger) Log(s string) error {
	_, err := c.Writer.Write([]byte(s))
	return err
}

func LogAll(loggers []Logger, message string) {
	for _, logger := range loggers {
		err := logger.Log(message)
		if err != nil {
			log.Println("Failed to log message:", err)
		}
	}
}

func main() {
	consoleLogger := &ConsoleLogger{Writer: os.Stdout}
	fileLogger := &FileLogger{fileName: "logging.txt"}

	loggers := []Logger{consoleLogger, fileLogger}
	LogAll(loggers, "This is a test log message.")
}
