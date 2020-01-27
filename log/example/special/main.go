package main

import (
	"github.com/ripx80/brave/log/logger"
	blogrus "github.com/ripx80/brave/log/logrus"
	bzap "github.com/ripx80/brave/log/zap"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type fields map[string]interface{}

func main() {

	// for advanced config and then set the configured to package
	l := logrus.New()
	// you can set a formatter
	l.SetFormatter(&logrus.JSONFormatter{})

	//use own instance
	log := logger.New(blogrus.Configured(l))

	/*
		// or normal
		log:=logger.New()
	*/

	// no fields
	log.Info("test")
	// only fields for this entry
	log.WithFields(fields{"user": 1234, " err": nil}).Info("test2")

	// set global fields
	log.Set(log.WithFields(fields{"user": "peter", " err": nil}))
	log.Info("test4")

	/*
		log.Set(bzap.New())
	*/

	zaplog, _ := zap.NewProduction()
	zl := bzap.Configured(zaplog)
	log.Set(zl)

	log.Info("zap says huhu")
	log.WithFields(fields{"user": "peter", " err": nil}).Info("bam")

}
