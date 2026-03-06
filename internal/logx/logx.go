package logx

import (
	"fmt"
	"log"
	"os"
)

// Logger is a module-prefixed logger enriched with hostname.
type Logger struct {
	module   string
	hostname string
}

// New creates a new logger for a given module.
func New(module string) *Logger {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return &Logger{module: module, hostname: hostname}
}

func (l *Logger) Info(format string, args ...any) {
	log.Printf("[%s] [host=%s] %s", l.module, l.hostname, fmt.Sprintf(format, args...))
}

func (l *Logger) Error(format string, args ...any) {
	log.Printf("[%s] [host=%s] ERROR: %s", l.module, l.hostname, fmt.Sprintf(format, args...))
}
