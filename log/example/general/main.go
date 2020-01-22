/*Package main is a example impl*/
package main

import (
	"github.com/ripx80/brave/log/logger"
	lr "github.com/ripx80/brave/log/logrus"

	example "github.com/ripx80/brave/log/example/somelib"

	"github.com/sirupsen/logrus"
)

func main() {
	// need a init logger
	logrusLog := logrus.New()

	// or use the defaults
	logrusLog = lr.New()

	logrusLog.SetLevel(logrus.DebugLevel)

	log := *logger.New(logrusLog)
	logger.Set(log)
	log.Infof("test")

	/*init new logrus with fields*/

	log.Debugf("test")

	/* Full control of all features of the logger */
	logrusLog.SetLevel(logrus.WarnLevel)
	log.Debugf("silent me")

	/* standard error msgs, for testing */
	log.InvalidArg("p")
	log.InvalidConfig("unkown-config")
	/* use in libs*/

	log = *logger.New(logrusLog.WithFields(logrus.Fields{"animal": "walrus"}))
	example.SomeCall()
	// you realy log in libs? no magic!
	example.SomeCallArg(log)
}
