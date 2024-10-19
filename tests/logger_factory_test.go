package tests

import (
	"reflect"
	"testing"

	"golang-logger-system/internal/logger"
	"golang-logger-system/internal/factory"
)

func TestLoggerFactory(t *testing.T) {
	// Create new factory
	factory := factory.NewConcreteLoggerFactory()

	// Setup a slice of tests to run in a loop
	tests := []struct {
		name 			string
		loggerType		string
		expectedType 	interface{}
		expectError		bool
	}{
		// For the sake of the test to have a type to compare to we make a zero-value instances of the two structs. We are expecting types from reflect on the pointers for both loggers
			// ! So, reflect.TypeOf((*logger.ConsoleLogger)(nil)) returns the type *logger.ConsoleLogger without actually creating a ConsoleLogger instance
		{"Console Logger", "console", reflect.TypeOf((*logger.ConsoleLogger)(nil)), false},
		{"File Logger", "file", reflect.TypeOf((*logger.FileLogger)(nil)), false},
		{"Unsupported Logger", "unsupported", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := factory.CreateLogger(tt.loggerType)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
			} else {
				if err !=nil {
					t.Errorf("Unexpected error %v", err)
				}

				if logger == nil {
					t.Errorf("Expected a logger, but got nil")
				} else if reflect.TypeOf(logger) != tt.expectedType {
					t.Errorf("Expected logger of type %T, but got %T", tt.expectedType, logger)
				}
			}
		})
	}
}