/*Package main is a example impl*/
package main

import (
	logger "github.com/ripx80/log"
	lr "github.com/ripx80/log/logrus"

	"github.com/ripx80/log/example"


	"github.com/sirupsen/logrus"
)

func main() {
	// need a init logger
	logrusLog := logrus.New()

	// or use the defaults
	logrusLog = lr.New()

	logrusLog.SetLevel(logrus.DebugLevel)

	//log.SetLogger(logrusLog)

	log := *logger.New(logrusLog)
	logger.Set(log)


	log.Infof("test")

	/*init new logrus with fields*/


	log.Debugf("test")

	/* Full control of all features of the logger */
	logrusLog.SetLevel(logrus.WarnLevel)
	log.Debugf("silent me")

	/* standard error msgs */
	log.InvalidArg("p")
	log.InvalidConfig("unkown-config")
	/* use in libs*/

	log = *logger.New(logrusLog.WithFields(logrus.Fields{"animal": "walrus",}))
	logger.Set(log)

	example.SomeCall()
	example.SomeCallArg(log)
}
