package zap

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/ripx80/log/config"
	"github.com/ripx80/log/logger"
	"go.uber.org/zap"
)

func RegisterLog(lc config.LogConfig) error {
	zLogger, err := initLog(lc)
	if err != nil {
		return errors.Wrap(err, "RegisterLog")
	}
	defer zLogger.Sync()
	zSugarlog := zLogger.Sugar()
	//zSugarlog.Info()
	//This is for loggerWrapper implementation
	//appLogger.SetLogger(&loggerWrapper{zaplog})

	logger.SetLogger(zSugarlog)
	return nil
}

// initLog create logger
func initLog(lc config.LogConfig) (zap.Logger, error) {
	rawJSON := []byte(`{
	 "level": "info",
     "Development": true,
      "DisableCaller": false,
	 "encoding": "console",
	 "outputPaths": ["stdout", "../../../demo.log"],
	 "errorOutputPaths": ["stderr"],
	 "encoderConfig": {
		"timeKey":        "ts",
		"levelKey":       "level",
		"messageKey":     "msg",
         "nameKey":        "name",
		"stacktraceKey":  "stacktrace",
         "callerKey":      "caller",
		"lineEnding":     "\n",
        "timeEncoder":     "time",
		"levelEncoder":    "lowercaseLevel",
        "durationEncoder": "stringDuration",
         "callerEncoder":   "shortCaller"
	 }
	}`)

	var cfg zap.Config
	var zLogger *zap.Logger
	//standard configuration
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		return *zLogger, errors.Wrap(err, "Unmarshal")
	}
	//customize it from configuration file
	err := customizeLogFromConfig(&cfg, lc)
	if err != nil {
		return *zLogger, errors.Wrap(err, "cfg.Build()")
	}
	zLogger, err = cfg.Build()
	if err != nil {
		return *zLogger, errors.Wrap(err, "cfg.Build()")
	}

	zLogger.Debug("logger construction succeeded")
	return *zLogger, nil
}

// customizeLogFromConfig customize log based on parameters from configuration file
func customizeLogFromConfig(cfg *zap.Config, lc config.LogConfig) error {
	cfg.DisableCaller = !lc.EnableCaller

	// set log level
	l := zap.NewAtomicLevel().Level()
	err := l.Set(lc.Level)
	if err != nil {
		return errors.Wrap(err, "")
	}
	cfg.Level.SetLevel(l)

	return nil
}
