// Package log represents a generic logging interface
package log

// Logger represent common interface for logging function
type logger interface {
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	//Info(args ...interface{})
	Debugf(format string, args ...interface{})
	//Debug(args ...interface{})
	//WFields(keyValues Fields) Logger
}

/*Logger holds some funcs */
type Logger struct{ logger }

/*Fields to log withFields*/
type Fields map[string]interface{}
type noOpLogger struct{}

/* Only use one format no format string? */
func (n *noOpLogger) Errorf(format string, v ...interface{}) {}
func (n *noOpLogger) Infof(format string, v ...interface{})  {}
func (n *noOpLogger) Debugf(format string, v ...interface{}) {}

/*Log global Log instance, noop logger is a placeholder*/
var Log = Logger{&noOpLogger{}}

/*New return a new logger instance*/
func New(log logger) *Logger {
	return &Logger{
		logger: log,
	}
}

/*Set the package global logger*/
func Set(log logger) {
	Log = Logger{
		logger: log,
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

/*InvalidArg default log entry*/
func (l *Logger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage, argumentName)
}

/*InvalidArgValue default log entry*/
func (l *Logger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage, argumentName, argumentValue)
}

/*MissingArg default log entry*/
func (l *Logger) MissingArg(argumentName string) {
	l.Errorf(missingArgMessage, argumentName)
}

/*InvalidConfig default log entry*/
func (l *Logger) InvalidConfig(argumentName string) {
	l.Errorf(invalidArgMessage, argumentName)
}
