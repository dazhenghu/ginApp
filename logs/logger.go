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

func Debugf(format string, args ...interface{}) {
    log.Debugf(format, args...)
}

func Printf(format string, args ...interface{}) {
    log.Printf(format, args...)
}

func Infof(format string, args ...interface{}) {
    log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
    log.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
    log.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
    log.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
    log.Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
    log.Fatalf(format, args...)
}

func WithFields(fields log.Fields) *log.Entry {
    return log.WithFields(fields)
}

func SetOutput(out io.Writer) {
    log.SetOutput(out)
}

/**
设置日志级别
 */
func SetLevel(lvl Level)  {
    log.SetLevel(log.Level(lvl))
}
