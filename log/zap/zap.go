package zap

import (
	"github.com/ripx80/brave/log/logger"
	"go.uber.org/zap"
)

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

/*New returns a new logger wrapper*/
func New() logger.LoggerInt {
	logger, _ := zap.NewProduction()
	return &zapLogger{
		sugaredLogger: logger.Sugar(),
	}
}

/*Configured return a new wrapper with the given logger*/
func Configured(l *zap.Logger) logger.LoggerInt {
	return &zapLogger{
		sugaredLogger: l.Sugar(),
	}
}

/*Debug wrapper*/
func (l *zapLogger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

/*Info wrapper*/
func (l *zapLogger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

/*Warn wrapper*/
func (l *zapLogger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

/*Error wrapper*/
func (l *zapLogger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

/*Panic wrapper*/
func (l *zapLogger) Panic(args ...interface{}) {
	l.sugaredLogger.Panic(args...)
}

/*WithFields wrapper, must use wrapper because of this function. */
func (l *zapLogger) WithFields(fields logger.Fields) logger.LoggerInt {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger}
}
