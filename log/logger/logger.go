// Package logger represents a generic logging interface
package logger

// Logger represent common interface for logging function
type logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	//WithField(key string, value interface{}) Logger
}

/*Logger holds some funcs */
type Logger struct{ logger }

/*
Log global Log instance, i dont like this but i need a cmd init logger in all libs
*/
var Log Logger

type LogError struct {
	id      int
	message string
}

//https://tldp.org/LDP/abs/html/exitcodes.html or internal codes
var (
	invalidArgMessage      = LogError{1, "Invalid arg: %s"}
	invalidArgValueMessage = LogError{2, "Invalid value for argument: %s: %v"}
	missingArgMessage      = LogError{3, "Missing arg: %s"}
	invalidConfigMessage   = LogError{4, "Invalid config: %s"}
	missingConfigValue     = LogError{5, "Missing config value: %s"}
)

func SetLogger(newLogger logger) {
	Log.logger = newLogger
}

func (l *Logger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage.message, argumentName)
}

func (l *Logger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

func (l *Logger) MissingArg(argumentName string) {
	l.Errorf(missingArgMessage.message, argumentName)
}

func (l *Logger) InvalidConfig(argumentName string) {
	l.Errorf(invalidArgMessage.message, argumentName)
}
