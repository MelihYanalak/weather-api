package logger

import (
	"fmt"
	"os"
	"sync"
)

// LogLevel represents different logging levels
type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
)

// Logger is the interface for logging
type Logger interface {
	Debug(message string)
	Info(message string)
	Error(message string)
	Warning(message string)
}

// BasicFileLogger is a basic file-based implementation of Logger
type FileLogger struct {
	level LogLevel
	file  *os.File
	mu    sync.Mutex
}

func NewFileLogger(level LogLevel, filename string) (*FileLogger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &FileLogger{
		level: level,
		file:  file,
	}, nil
}

func (l *FileLogger) Debug(message string) {
	l.log("[DEBUG]", message)
}

func (l *FileLogger) Info(message string) {
	l.log("[INFO]", message)
}

func (l *FileLogger) Error(message string) {
	l.log("[ERROR]", message)
}

func (l *FileLogger) Warning(message string) {
	l.log("[WARNING]", message)
}

func (l *FileLogger) log(level string, message string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level <= DebugLevel || l.level <= InfoLevel || l.level <= WarningLevel || l.level <= ErrorLevel {
		fmt.Fprintf(l.file, "%s %s\n", level, message)
	}
}

func (l *FileLogger) Close() error {
	return l.file.Close()
}
