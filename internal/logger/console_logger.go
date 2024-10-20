package logger

import (
	"fmt"
	"time"
)

// ConsoleLogger is alogger that writes to the console
type ConsoleLogger struct {}

// Function that creates a new ConsoleLogger 
	// (a function being responsible for creating a returning a logger is an example of the factory design pattern)
func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

// Method to perform the logging
	// Create a formatted timestamp for "Now()"
	// Will format a new message composed of an incoming message and whatever other args provided. (variatic function)
	// Will print to console, the final message , which has
		// - timestamp
		// - level of severity
		// - formatted message
func (l *ConsoleLogger) log(level, message string, args ...interface{}) {
	timestamp := time.Now().Format("02 Jan 06 15:04 -0700")
	formattedMessage := fmt.Sprintf(message, args...)
	fmt.Printf("[%v] %s: %s\n", timestamp, level, formattedMessage)
}

func (l *ConsoleLogger) Info(message string, args ...interface{}) {
	l.log("INFO", message, args...)
}
func (l *ConsoleLogger) Error(message string, args ...interface{}) {
	l.log("ERROR", message, args...)
}
func (l *ConsoleLogger) Debug(message string, args ...interface{}) {
	l.log("DEBUG", message, args...)
}