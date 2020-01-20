package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/ripx80/brave/exit"
	"github.com/ripx80/brave/work"
)

func someFunc() error {
	fmt.Println("do stuff here")
	return nil
}

func main() {
	/* Add Threat Safe Go exit, handle os signals and errors in goroutines*/
	defer exit.Safe()

	stop := make(chan struct{}) // stop chan for all routines
	done := make(chan struct{}) // stop chan for all routines
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM) // register syscalls
	wg := new(sync.WaitGroup)

	go func() {
		defer wg.Done()
		wg.Add(1)
		// blocking func. routine work
		if err := work.TimeoutFunc(stop, someFunc, 2*time.Second); err != nil {
			fmt.Printf("error: %v\n", err)
			exit.Exit(1)
		}
		// work done exit
		done <- struct{}{}
	}()

	for {
		select {
		//event that main must care about. In this example its a msg string
		case <-signals:
		case <-stop:
		case <-done:
		}

		close(stop)
		work.WaitTimeout(wg, 1*time.Second) // wait for all workers with timeout
		exit.Exit(0)                        // check the return
	}
}
