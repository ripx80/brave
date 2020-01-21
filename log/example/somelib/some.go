package example

import (
	"github.com/ripx80/brave/log/logger"
)

func SomeCall() {
	//you can use the package global log but when it is not init you have runtime errors
	logger.Log.Errorf("test case")

}

/*SomeCallArg clean solution with explicit logger, no magic */
func SomeCallArg(l logger.Logger) {
	// take from arg
	l.Errorf("test with arg case")
}
