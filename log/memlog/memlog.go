/*Package memlog is a in memory logger for testing*/
package memlog

import (
	"fmt"
	"sync"
)

// Logger implements the log.Logger interface by capturing logged
// lines.
type Logger struct {
	mutex *sync.Mutex
	// Entries holds logged entries in submission order.
	Entries []string
}

func (logger *Logger) Infof(format string, v ...interface{}) {
	logger.log(fmt.Sprintf(format, v...))
}

func (logger *Logger) Debugf(format string, v ...interface{}) {
	logger.log(fmt.Sprintf(format, v...))
}

func (logger *Logger) Errorf(format string, v ...interface{}) {
	logger.log(fmt.Sprintf(format, v...))
}

func (logger *Logger) log(entry string) {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.Entries = append(logger.Entries, entry)
}

func New() *Logger {
	return &Logger{
		mutex:   &sync.Mutex{},
		Entries: []string{},
	}
}
