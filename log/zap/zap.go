package zap

import (
	"encoding/json"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

/*
not working at the moment. panics
can be used for setting up defaults
*/

/*New create a new zap logger with defaults*/
func New() (*zap.SugaredLogger, error) {
	var cfg zap.Config
	var zLogger *zap.Logger
	defer zLogger.Sync()
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

	//standard configuration
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		return nil, errors.Wrap(err, "Unmarshal")
	}
	//customize it from configuration file
	sugar := zLogger.Sugar()
	return sugar, nil
}
