package example

import (
	"github.com/ripx80/log/logger"
)

func SomeCall() {
	// you can use the package global log but when it is not init you have runtime errors
	logger.Log.Errorf("test case")

}

/* Clean solution with explicit logger */
func SomeCallArg(l logger.Logger) {
	// take from arg
	l.Errorf("test with arg case")
}
