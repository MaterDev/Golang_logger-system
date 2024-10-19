package tests

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"golang-logger-system/internal/logger"
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

func TestConsoleLogger(t *testing.T) {
	// Redirect stout to capture output
	oldStdout := os.Stdout
	// ? `r` is the read end of the pipe (what will be used to get the captured output)
	// ? `w` is the write end of the pipe. (Anything written to pipe will be written here.) 
	r, w, _ := os.Pipe()
	// ? Redirect stdout to the write end of the pipe
	os.Stdout = w // Will clear Stdoutput by assigning it to the empty value `w` (where things get written)

	// ! Subsequent writes will go to the pipe instead of the terminal.


	// Create and use the logger
	l := logger.NewConsoleLogger()
	l.Info("Test Info message")
	l.Error("Test Error message")
	l.Debug("Test Debug message")

	// Restore stdout
	w.Close() // Close the write pipe (No more data will be written to pipe.)
		// If the written end of the pipe isnt closed, subsequent operations will be waiting for more data.
	os.Stdout = oldStdout // Reconnect stdout to terminal

	// ! Read captured output:
	
	// Make buffer for storing bytes
	var buff bytes.Buffer
	// Copy read end of the pipe to the address of the buffer
	io.Copy(&buff, r)
	// Return contents of buffer as a strong to store in output variable
	output := buff.String()

	expectedMessages := []string{
		"INFO: Test Info message",
		"ERROR: Test Error message",
		"DEBUG: Test Debug message",
	}

	for _, msg := range expectedMessages {
		if !strings.Contains(output, msg) {
			t.Errorf("Expected output to contain '%s', but it didn't", msg)
		}
	}
}

func TestFileLogger(t *testing.T) {
	// Create temporary file for testing
		// in the default testing directory, name beginning with test_log
	tmpfile, err := os.CreateTemp("", "test_log")

	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	// At the end of the scope will remove the temporary file
	defer os.Remove(tmpfile.Name())

	// Create and use logger, giving the temporary file
	l, err := logger.NewFileLogger(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to create FileLogger: %v", err)
	}
	// Close the file at the end
	defer l.Close()

	// Call all the log methods (to write to the temp file)
	l.Info("Test Info message")
	l.Error("Test Error message")
	l.Debug("Test Debug message")

	// Read the log file
	content, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	expectedMessages := []string{
		"INFO: Test Info message",
		"ERROR: Test Error message",
		"DEBUG: Test Debug message",
	}

	for _, msg := range expectedMessages {
		if !strings.Contains(string(content), msg) {
			t.Errorf("Expected output to contain '%s', but it didn't", msg)
		}
	}
}
