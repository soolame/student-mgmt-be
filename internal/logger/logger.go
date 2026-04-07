package logger

import (
	"log"
	"os"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

type Logger struct {
	level  Level
	logger *log.Logger
}

var defaultLogger *Logger

func Init(level Level) {
	defaultLogger = &Logger{
		level:  level,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func SetLevel(level Level) {
	defaultLogger.level = level
}

func (l *Logger) log(level Level, prefix string, msg string, args ...any) {
	if level < l.level {
		return
	}
	l.logger.Printf(prefix+" "+msg, args...)
}

func Debug(msg string, args ...any) {
	defaultLogger.log(DEBUG, "[DEBUG]", msg, args...)
}

func Info(msg string, args ...any) {
	defaultLogger.log(INFO, "[INFO]", msg, args...)
}

func Warn(msg string, args ...any) {
	defaultLogger.log(WARN, "[WARN]", msg, args...)
}

func Error(msg string, args ...any) {
	defaultLogger.log(ERROR, "[ERROR]", msg, args...)
}
