package logger

// Logger is the interface that wraps the basic logging methods
type Logger interface {
	Info(message string, args ...interface{})
	Error(message string, args ...interface{})
	Debug(message string, args ...interface{})
}

