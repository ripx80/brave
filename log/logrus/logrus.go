package logrus

import (
	"github.com/ripx80/brave/log/logger"
	"github.com/sirupsen/logrus"
)

type logrusLogEntry struct {
	entry *logrus.Entry
}

type logrusLogger struct {
	logger *logrus.Logger
}

func getFormatter(isJSON bool) logrus.Formatter {
	if isJSON {
		return &logrus.JSONFormatter{}
	}
	return &logrus.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	}
}

/*New return a new logger wrapper*/
func New() logger.LoggerInt {
	return &logrusLogger{
		logger: logrus.New(),
	}
}

/*Configured return a new wrapper with the given logger*/
func Configured(l *logrus.Logger) logger.LoggerInt {
	return &logrusLogger{
		logger: l,
	}
}

/*Debug wrapper*/
func (l *logrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

/*Info wrapper*/
func (l *logrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

/*Warn wrapper*/
func (l *logrusLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

/*Error wrapper*/
func (l *logrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

/*Panic wrapper*/
func (l *logrusLogger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

/*Debug wrapper*/
func (l *logrusLogEntry) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

/*Info wrapper*/
func (l *logrusLogEntry) Info(args ...interface{}) {
	l.entry.Info(args...)
}

/*Warn wrapper*/
func (l *logrusLogEntry) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

/*Error wrapper*/
func (l *logrusLogEntry) Error(args ...interface{}) {
	l.entry.Error(args...)
}

/*Panic wrapper*/
func (l *logrusLogEntry) Panic(args ...interface{}) {
	l.entry.Panic(args...)
}

/*WithFields wrapper, must use wrapper because of this function. */
func (l *logrusLogger) WithFields(fields logger.Fields) logger.LoggerInt {
	return &logrusLogEntry{
		entry: l.logger.WithFields(convertToLogrusFields(fields)),
	}

}

/*WithFields wrapper*/
func (l *logrusLogEntry) WithFields(fields logger.Fields) logger.LoggerInt {
	return &logrusLogEntry{
		entry: l.entry.WithFields(convertToLogrusFields(fields)),
	}
}

func convertToLogrusFields(fields logger.Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}
