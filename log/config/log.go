package config

type LogConfig struct {
	// log library name
	Code string `yaml:"code"`
	// log level
	Level string `yaml:"level"`
	// show caller in log message
	EnableCaller bool `yaml:"enableCaller"`
}

/* improved config */
// type Configuration struct {
// 	EnableConsole     bool
// 	ConsoleJSONFormat bool
// 	ConsoleLevel      string
// 	EnableFile        bool
// 	FileJSONFormat    bool
// 	FileLevel         string
// 	FileLocation      string
// }

// const (
// 	LOGRUS string = "logrus"
// 	ZAP    string = "zap"
// )
