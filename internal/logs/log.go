package logs

import (
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

type LogLevel int

type Logger struct {
	mu       sync.Mutex
	out      io.Writer
	minLevel LogLevel
	log      *log.Logger
}

func New(out io.Writer, minLevel LogLevel) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
		log:      log.New(out, "", 0),
	}
}

// loger consts

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

var levelColors = map[LogLevel]string{
	DEBUG: "\033[36m",
	INFO:  "\033[32m",
	WARN:  "\033[33m",
	ERROR: "\033[31m",
}

func (l *Logger) logf(level LogLevel, format string, args ...any) {
	if level < l.minLevel {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	color := levelColors[level]
	reset := "\033[0m"

	timeStap := time.Now().Format("2025-01-01 00:00:00")

	levelStr := [...]string{"DEBUG", "INFO", "WARN", "ERROR"}[level]

	msg := fmt.Sprintf(format, args...)
	output := fmt.Sprintf("%s [%s%s%s] %s", timeStap, color, levelStr, reset, msg)

	l.log.Println(output)
}

func (l *Logger) Debug(format string, args ...any) { l.logf(DEBUG, format, args...) }
func (l *Logger) Info(format string, args ...any)  { l.logf(INFO, format, args...) }
func (l *Logger) Warn(format string, args ...any)  { l.logf(WARN, format, args...) }
func (l *Logger) Error(format string, args ...any) { l.logf(ERROR, format, args...) }
