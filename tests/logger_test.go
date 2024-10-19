package tests

import (
	"golang-logger-system/internal/logger"
	"testing"
)

// Mock implementation of the Logger interface for testing
type MockLogger struct { // Implements the logger interface, because it has the same methods of Logger (Info, Error, Debug)
	InfoCalled bool
	ErrorCalled bool
	DebugCalled bool
}

func (m *MockLogger) Info(message string, args ...interface{}) {m.InfoCalled = true}
func (m *MockLogger) Error(message string, args ...interface{}) {m.ErrorCalled = true}
func (m *MockLogger) Debug(message string, args ...interface{}) {m.DebugCalled = true}


func TestLoggerInterface(t *testing.T)	{
	// `l` is using the `Logger` interface as its type.
		// Create a new instance of `MockLogger`` and assign its address aof`l`
			// This works because `MockLogger` implements the m methods defined in logger.Logger
		// The fields of `l` from `MockLogger` still exist, but cant be directly accessed by `l`
	// Holds:
		// pointer to underlying concrete value
		// Information about the type of that value

	// use &MockLogger{} because methods are defined on *MockLogger, not MockLogger.
		// Necessary for making thje fields mutable
	var l logger.Logger = &MockLogger{} // Otherwise fields of `MockLogger` wouldnt be accessible, because `l` is a logger type.

	// Call all methods
	l.Info("test")
	l.Error("test")
	l.Debug("test")

	// This allows us to access the implementations (value from the fields of MockLogger)
	mock := l.(*MockLogger) // `.(*MockLogger)` a type assertion to access the concrete type behind an interace.

	if !mock.InfoCalled {
		t.Error("Info method was not called")
	}
	if !mock.ErrorCalled {
		t.Error("Error method was not called")
	}
	if !mock.DebugCalled {
		t.Error("Debug method was not called")
	}
}
