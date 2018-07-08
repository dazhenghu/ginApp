package logs

import log "github.com/sirupsen/logrus"

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

func WithFields(fields log.Fields) *log.Entry {
    return log.WithFields(fields)
}

/**
设置日志级别
 */
func SetLevel(lvl Level)  {
    log.SetLevel(log.Level(lvl))
}
