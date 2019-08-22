// Package invision provides Logur integration for InVision go-logger interface.
package invision

import (
	"fmt"
	"strings"

	log "github.com/InVisionApp/go-logger"
	"logur.dev/logur"
)

// Logger is a github.com/InVisionApp/go-logger.Logger logger.
type Logger struct {
	logger logur.Logger
}

// New returns a new github.com/InVisionApp/go-logger.Logger logger.
func New(logger logur.Logger) *Logger {
	return &Logger{logger: logger}
}

func (l *Logger) Debug(msg ...interface{}) {
	l.logger.Debug(fmt.Sprint(msg...))
}

func (l *Logger) Info(msg ...interface{}) {
	l.logger.Info(fmt.Sprint(msg...))
}

func (l *Logger) Warn(msg ...interface{}) {
	l.logger.Warn(fmt.Sprint(msg...))
}

func (l *Logger) Error(msg ...interface{}) {
	l.logger.Error(fmt.Sprint(msg...))
}

func (l *Logger) Debugln(msg ...interface{}) {
	l.Debug(strings.TrimSuffix(fmt.Sprintln(msg...), "\n"))
}

func (l *Logger) Infoln(msg ...interface{}) {
	l.Info(strings.TrimSuffix(fmt.Sprintln(msg...), "\n"))
}

func (l *Logger) Warnln(msg ...interface{}) {
	l.Warn(strings.TrimSuffix(fmt.Sprintln(msg...), "\n"))
}

func (l *Logger) Errorln(msg ...interface{}) {
	l.Error(strings.TrimSuffix(fmt.Sprintln(msg...), "\n"))
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

// WithFields returns a new logger with the additional supplied fields.
func (l *Logger) WithFields(fields log.Fields) log.Logger {
	return &Logger{logur.WithFields(l.logger, fields)}
}
