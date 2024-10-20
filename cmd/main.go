package main

import (
	"fmt"
	
	"golang-logger-system/internal/factory"
)

func main() {
	// Create new logger factory
	loggerFactory := factory.NewConcreteLoggerFactory()

	// Create different types of loggers
	consoleLogger, err := loggerFactory.CreateLogger("console")
	if err != nil {
		fmt.Printf("Error creating console logger: %v\n", err)
	}

	fileLogger, err := loggerFactory.CreateLogger("file")
	if err != nil {
		fmt.Printf("Error creating console logger: %v\n", err)
	}

	// Use the loggers
	consoleLogger.Info("This is an info message from the console logger")
	consoleLogger.Error("This is an error message from the console logger")
	consoleLogger.Debug("This is an debug message from the console logger")

	fileLogger.Info("This is an info message from the file logger")
	fileLogger.Error("This is an error message from the file logger")
	fileLogger.Debug("This is an debug message from the file logger")

	fmt.Println("Logging complete. Check the log file for the file logger output. 📝")
	
}
