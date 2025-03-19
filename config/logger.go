package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	flags := log.Ldate | log.Ltime

	return &Logger{
		debug:   log.New(writer, p+"[DEBUG]: ", flags),
		info:    log.New(writer, p+"[INFO]: ", flags),
		warning: log.New(writer, p+"[WARNING]: ", flags),
		err:     log.New(writer, p+"[ERROR]: ", flags),
		writer:  writer,
	}
}

// Create non-formatted log messages
func (l *Logger) Debug(v ...any) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...any) {
	l.info.Println(v...)
}
func (l *Logger) Warn(v ...any) {
	l.warning.Println(v...)
}
func (l *Logger) Error(v ...any) {
	l.err.Println(v...)
}

// Create formatted log messages
func (l *Logger) Debugf(format string, v ...any) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ...any) {
	l.info.Printf(format, v...)
}
func (l *Logger) Warnf(format string, v ...any) {
	l.warning.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...any) {
	l.err.Printf(format, v...)
}
