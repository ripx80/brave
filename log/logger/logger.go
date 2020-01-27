// Package logger represents a generic logging interface
package logger

// Logger represent common interface for logging function
type LoggerInt interface {
	Error(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Debug(args ...interface{})
	Panic(args ...interface{})
	WithFields(f map[string]interface{}) LoggerInt // must have, then i can remove the format strings
}

/*Logger holds some funcs */
type Logger struct{ LoggerInt }

type noOpLogger struct{}

/* Only use one format no format string? */
func (n *noOpLogger) Panic(v ...interface{})                        {}
func (n *noOpLogger) Error(v ...interface{})                        {}
func (n *noOpLogger) Warn(v ...interface{})                         {}
func (n *noOpLogger) Info(v ...interface{})                         {}
func (n *noOpLogger) Debug(v ...interface{})                        {}
func (n *noOpLogger) WithFields(f map[string]interface{}) LoggerInt { return &noOpLogger{} }

//func (n *noOpLogger) WithFields(v LogFields) Logger          { return Logger{&noOpLogger{}} } // look at this

/*Log global Log instance, noop logger is a placeholder*/
var log = Logger{&noOpLogger{}} // no magic in libs but you want it

/*New return a new logger instance*/
func New(l LoggerInt) *Logger {
	return &Logger{
		LoggerInt: l,
	}
}

/*Set the package global logger with magic fingers*/
func Set(l LoggerInt) {
	log = Logger{
		LoggerInt: l,
	}
}

// no exit codes.exit is main stuff no logging stuff
const (
	invalidArgMessage      = "Invalid arg: %s"
	invalidArgValueMessage = "Invalid value for argument: %s: %v"
	missingArgMessage      = "Missing arg: %s"
	invalidConfigMessage   = "Invalid config: %s"
	missingConfigValue     = "Missing config value: %s"
)

func (l *Logger) Set(nl LoggerInt) {
	l.LoggerInt = nl
}

/*InvalidArg default log entry*/
func (l *Logger) InvalidArg(argumentName string) {
	l.Error(invalidArgMessage, argumentName)
}

/*InvalidArgValue default log entry*/
func (l *Logger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Error(invalidArgValueMessage, argumentName, argumentValue)
}

/*MissingArg default log entry*/
func (l *Logger) MissingArg(argumentName string) {
	l.Error(missingArgMessage, argumentName)
}

/*InvalidConfig default log entry*/
func (l *Logger) InvalidConfig(argumentName string) {
	l.Error(invalidArgMessage, argumentName)
}

/*Errorf wrapper func*/
func Error(v ...interface{}) {
	log.Error(v[:])
}

/*Info wrapper func*/
func Info(v ...interface{}) {
	log.Info(v[:])
}

/*Info wrapper func*/
func Warn(v ...interface{}) {
	log.Warn(v[:])
}

/*Debugf wrapper func*/
func Debug(format string, v ...interface{}) {
	log.Debug(format, v)
}

/*Panic wrapper func*/
func Panic(v ...interface{}) {
	log.Panic(v)
}
