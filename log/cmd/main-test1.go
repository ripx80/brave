package main

import (
	"github.com/ripx80/log/config"
	"github.com/ripx80/log/logger"
	"github.com/ripx80/log/loggerfactory" // with loggerfactory you can switch between all available loggers
)

func LogMe() {
	logger.Log.Debug("Debug case")
	logger.Log.Info("Info case")
	//logger.Log.Warnf("Warning case") //zap throw a complete error stack
	// logger.Log.Errorf("Error case")
	//logger.Log.Fatal("Fatal case")
}

func main() {

	/*
		go internal log without interface? hijack logrus log?

		add a Format String config option
		need a withFields function, check if zap had one
		logrus.SetFormatter same in zap or other functions
		SetFlags function?
		rewrite the fearlessly code and remove the Panic
		set log of lib to none
		add events and error msg with const
	*/

	lc := config.LogConfig{
		Code:         "logrus",
		Level:        "Debug",
		EnableCaller: true,
	}

	_ = loggerfactory.GetLogFactoryBuilder(lc.Code).Build(&lc)

	// if err != nil {
	// 	return errors.Wrap(err, "")
	// }
	LogMe()

	// We want a silent lib, only realy bad problems rise
	lc = config.LogConfig{
		Code:         "logrus",
		Level:        "Fatal",
		EnableCaller: true,
	}
	loggerfactory.GetLogFactoryBuilder(lc.Code).Build(&lc)
	LogMe()

	//change the logger in realtime

	lc = config.LogConfig{
		Code:         "zap",
		Level:        "Info",
		EnableCaller: true,
	}
	loggerfactory.GetLogFactoryBuilder(lc.Code).Build(&lc)
	LogMe()

	logger.Log.Info("Goodbye")
}
