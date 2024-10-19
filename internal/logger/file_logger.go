package logger

import (
	"fmt"
	"os"
	"time"
)

// FileLogger is a logger that writes to a file
type FileLogger struct {
	file *os.File
}

// NewFileLogger method create a new FileLogger, which will write formatted errror messages to a specified file
func NewFileLogger(filename string) (*FileLogger, error) {

	// ! OpenFile(file path, flags, permission):
		// Flags can be compounded using the OR operator `|`
			// os.O_APPEND - IF file exists will append data at the end instead of overwriting it
			// os.O_CREATE	- Will create file if it doesnt exist
			// os.O_WRONLY - Open the file as write only
		// Permission is a number that represents the privledges for the file (readable by everyone, writable only bty owner)
			// 6: read and write permission for the owner
			// 4: read permission for the group
			// 4: read permission for others

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open the log file: %v", err)
	}
	return &FileLogger{file: file}, nil
} 

// Method to create the log in the file
func (l *FileLogger) log(level, message string, args ...interface{}) {
	timestamp := time.Now().Format("02 Jan 06 15:04 -0700")
	formattedMessage := fmt.Sprintf(message, args...)
	logEngry := fmt.Sprintf("[%s] %s: %s\n", timestamp, level, formattedMessage)
	l.file.WriteString(logEngry) // * Write the log string to the file.
}


// Method for logging INFO level
func (l *FileLogger) Info(message string, args ...interface{}) {
	l.log("INFO", message, args...)
}

// Method for logging ERROR level
func (l *FileLogger) Error(message string, args ...interface{}) {
	l.log("ERROR", message, args...)
}
// Method for logging DEBUG level
func (l *FileLogger) Debug(message string, args ...interface{}) {
	l.log("DEBUG", message, args...)
}

// Method for closing the log file
func (l *FileLogger) Close() error {
	return l.file.Close()
}