package logs

import (
    log "github.com/sirupsen/logrus"
    "io"
)

func Info(args ...interface{})  {
    log.Info(args...)
}

func Debug(args ...interface{})  {
    log.Debug(args...)
}

func Warn(args ...interface{}) {
    log.Warn(args...)
}

func Warning(args ...interface{}) {
    log.Warning(args...)
}

func Error(args ...interface{})  {
    log.Error(args...)
}

func Panic(args ...interface{}) {
    log.Panic(args...)
}

func Fatal(args ...interface{}) {
    log.Fatal(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
    log.Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
    log.Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
    log.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
    log.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
    log.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
    log.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
    log.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
    log.Fatalf(format, args...)
}

func WithFields(fields log.Fields) *log.Entry {
    return log.WithFields(fields)
}

// SetOutput sets the standard logger output.
func SetOutput(out io.Writer) {
    log.SetOutput(out)
}

/**
设置日志级别
 */
func SetLevel(lvl Level)  {
    log.SetLevel(log.Level(lvl))
}
