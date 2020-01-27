/*Package memlog is a in memory logger for testing*/
package memlog

import (
	"fmt"
	"sync"

	"github.com/ripx80/brave/log/logger"
)

// Logger implements the log.Logger interface by capturing logged
// lines.
type Logger struct {
	mutex *sync.Mutex
	// Entries holds logged entries in submission order.
	Entries []string
}

/*Info wrapper*/
func (l *Logger) Info(v ...interface{}) {
	l.log(fmt.Sprint(v...))
}

/*Warn wrapper*/
func (l *Logger) Warn(v ...interface{}) {
	l.log(fmt.Sprint(v...))
}

/*Debug wrapper*/
func (l *Logger) Debug(v ...interface{}) {
	l.log(fmt.Sprint(v...))
}

/*Error wrapper*/
func (l *Logger) Error(v ...interface{}) {
	l.log(fmt.Sprint(v...))
}

/*Panic wrapper*/
func (l *Logger) Panic(v ...interface{}) {
	l.log(fmt.Sprint(v...))
}

func (l *Logger) log(entry string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.Entries = append(l.Entries, entry)
}

/*New return a new Logger wrapper*/
func New() *Logger {
	return &Logger{
		mutex:   &sync.Mutex{},
		Entries: []string{},
	}
}

/*WithFields not supported im memolg*/
func (l *Logger) WithFields(fields map[string]interface{}) logger.LoggerInt {
	return &Logger{
		mutex:   &sync.Mutex{},
		Entries: []string{},
	}

}
