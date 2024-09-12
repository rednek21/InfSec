package logger

import (
	"fmt"
	"go.uber.org/zap"
)

//const (
//	// LOCAL variable for defining log mode with humanized console output
//	LOCAL = iota
//
//	// PRODUCTION variable for defining log mode with JSON output
//	PRODUCTION
//)

// ILogger interface declares methods for Logger struct
//
//go:generate go run github.com/vektra/mockery/v2@v2.44.1 --name ILogger
type ILogger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

// Logger struct to store logger
type Logger struct {
	logger *zap.Logger
}

const (
	DEV  = "dev"
	PROD = "prod"
)

// NewLogger constructor to create logger with passed env settings
// Gets logType int = 0(equals local), 1(equals production)
func NewLogger(env string) *Logger {
	var logger *zap.Logger
	var err error

	switch env {
	case PROD:
		logger, err = SetupProductionLogger()
	case DEV:
		logger, err = SetupLocalLogger()
	default:
		panic("invalid log type")
	}

	if err != nil {
		panic(err)
	}

	return &Logger{logger: logger}
}

// FormatMessage function for handling multiple args to format log message
func FormatMessage(msg string, args ...interface{}) string {
	if len(args) > 0 {
		return fmt.Sprintf(msg, args...)
	}
	return msg
}

// Debug method implementation of ILogger
func (l *Logger) Debug(msg string, args ...interface{}) {
	l.logger.Debug(FormatMessage(msg, args...))
}

// Info method implementation of ILogger
func (l *Logger) Info(msg string, args ...interface{}) {
	l.logger.Info(FormatMessage(msg, args...))
}

// Warn method implementation of ILogger
func (l *Logger) Warn(msg string, args ...interface{}) {
	l.logger.Warn(FormatMessage(msg, args...))
}

// Error method implementation of ILogger
func (l *Logger) Error(msg string, args ...interface{}) {
	l.logger.Error(FormatMessage(msg, args...))
}

// Fatal method implementation of ILogger // Need to be accurate calling that, bc it stops your program
func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.logger.Fatal(FormatMessage(msg, args...))
}
