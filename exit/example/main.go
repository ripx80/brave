package main

import (
	"fmt"
	"github.com/ripx80/brave/exit"
)

func main() {
	defer exit.Safe()
	defer fmt.Println("close db connection")

	fmt.Println("do stuff here")

	/*
		//examples

		// respect the defer function and send a error code to the system
		panic(ExitCode{3})

		// raise a normal panic with a string
		panic("raise a normal panic")

		//raise a normal panic with a int
		panic(3)

		// if you use logrus, dont use the fatal function. Its bad and do a os.Exit().
		logrus.Panic("panic logging")

		// fmt.Errorf("normal Error. then do a safe goExit")
	*/

	exit.Exit(1)
}
