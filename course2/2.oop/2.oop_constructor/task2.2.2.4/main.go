package main

import (
	"bufio"
	"os"
)

// Logger interface
type Logger interface {
	Log(message string)
}

type FileLogger struct {
	file *os.File
}
type LogSystem struct {
	logger FileLogger
}

// LogOption functional option type
type LogOption func(*LogSystem)

func (l *LogSystem) Log(m string) {
	writer := bufio.NewWriter(l.logger.file)
	writer.Write([]byte(m))
	writer.Flush()
}

func WithLogger(value FileLogger) LogOption {
	return func(logSystem *LogSystem) {
		logSystem.logger = value
	}
}

func NewLogSystem(option LogOption) *LogSystem {
	logg := &LogSystem{}
	option(logg)
	return logg
}

func main() {
	file, _ := os.Create("log.txt")
	defer file.Close()

	fileLogger := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))

	logSystem.Log("Hello, world!")
}
