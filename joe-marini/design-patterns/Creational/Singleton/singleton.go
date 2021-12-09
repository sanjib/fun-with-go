package main

// TODO: use the "sync" package for the Once API
import (
	"fmt"
	"sync"
)

// MyLogger is the struct we want to make a singleton
type MyLogger struct {
	loglevel int
}

// Log a message using the logger
func (l *MyLogger) Log(s string) {
	fmt.Println(l.loglevel, ":", s)
}

// SetLogLevel sets the log level of the logger
func (l *MyLogger) SetLogLevel(level int) {
	l.loglevel = level
}

// the logger instance
var logger *MyLogger

// use the sync package to enforce goroutine safety
var once sync.Once

// getLoggerInstance function provides global access to the
// logger class instance
func getLoggerInstance() *MyLogger {
	once.Do(func() {
		fmt.Println("creating logger instance")
		logger = &MyLogger{}
	})

	fmt.Println("returning logger instance")
	return logger
}
