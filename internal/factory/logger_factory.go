package factory

import (
	"fmt"
	"golang-logger-system/internal/logger"
)

// Interface for creating loggers
type LoggerFactory interface {
	CreateLogger(loggerType string) (logger.Logger, error)
}

// Concrete implementation of LoggerFactory
type ConcreteLoggerFactory struct{}

// Function to create a new ConcreteLoggerFAtory
func NewConcreteLoggerFactory() *ConcreteLoggerFactory {
	return &ConcreteLoggerFactory{}
}

func (f *ConcreteLoggerFactory) CreateLogger(loggerType string) (logger.Logger, error) {
	switch loggerType {
	case "console":
		return logger.NewConsoleLogger() , nil
	case "file":
		return logger.NewFileLogger("log.txt")
	default:
		return nil, fmt.Errorf("unsupported logger type: %s", loggerType)
	}
}

