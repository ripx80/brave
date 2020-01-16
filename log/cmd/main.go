package main

import (
	"github.com/ripx80/log/logger"
	"github.com/ripx80/log/some"

	"github.com/sirupsen/logrus"
)

/*
 // you can use the default logger and your own config
	"go.uber.org/zap"
*/

func main() {
	logrusLog := logrus.New()
	logrusLog.SetLevel(logrus.DebugLevel)

	//log.SetLogger(logrusLog)
	logger.SetLogger(logrusLog)
	logger.Log.Info("test")
	logger.Log.Debug("test")

	/* Full control of all features of the logger */
	logrusLog.SetLevel(logrus.WarnLevel)
	logger.Log.Debug("silent me")

	/* standard error msgs */
	logger.Log.InvalidArg("p")
	logger.Log.InvalidConfig("unkown-config")
	/* use in libs*/
	some.SomeCall()
}
