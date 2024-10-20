# Logger System with Factory Method Pattern

This project demonstrates the implementation of a flexible logging system in Go using the Factory Method pattern. It provides a simple way to create and use different types of loggers while maintaining a consistent interface.

![cover](./images/cover.png)

## Features

- Console Logger: Outputs log messages to the console
- File Logger: Writes log messages to a specified file
- Factory Method implementation for creating loggers
- Extensible design for easy addition of new logger types

## Project Structure

- `cmd/main.go`: Main application demonstrating logger usage
- `internal/logger/`: Logger interface and implementations
- `internal/factory/`: Logger factory implementation
- `tests/`: Test files for all components

## Setup

1. Clone the repository:
   ```
   git clone https://github.com/MaterDev/Golang_logger-system.git
   ```

2. Navigate to the project directory:
   ```
   cd Golang_logger-system
   ```

3. Run the main application:
   ```
   go run cmd/main.go
   ```

## Usage

To use the logger system in your own code:

1. Create a logger factory:
   ```go
   loggerFactory := factory.NewConcreteLoggerFactory()
   ```

2. Create a logger of your choice:
   ```go
   logger, err := loggerFactory.CreateLogger("console") // or "file"
   if err != nil {
       // Handle error
   }
   ```

3. Use the logger:
   ```go
   logger.Info("This is an info message")
   logger.Error("This is an error message")
   logger.Debug("This is a debug message")
   ```

## Running Tests

To run the tests for this project:

```bash
go test ./tests
```

## Design Pattern: Factory Method

This project demonstrates the Factory Method pattern, which provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created. 

Key components:

- `Logger` interface: Defines the common methods for all logger types
- Concrete loggers (`ConsoleLogger`, `FileLogger`): Implement the `Logger` interface
- `LoggerFactory` interface: Declares the factory method for creating loggers
- `ConcreteLoggerFactory`: Implements the factory method to create specific logger types

Benefits of this pattern in our logger system:

1. Encapsulation of object creation logic
2. Flexibility to add new logger types without modifying existing code
3. Consistent interface for using different logger types

## Next Steps

Potential enhancements for this project:

1. Implement a Formatter Logger for custom log message formatting
2. Add support for log levels (e.g., INFO, DEBUG, ERROR)
3. Implement log rotation for the File Logger
4. Create a configuration file for logger settings
