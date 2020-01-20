/*Package logrus handles creating logrus logger with default flags*/
package logrus

import (
	"github.com/sirupsen/logrus"
)

/*New return a customized/default logrus logger*/
func New() *logrus.Logger {
	//standard configuration
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})
	return log
}
