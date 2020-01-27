package main

import (
	"fmt"

	"github.com/ripx80/brave/log/logger"
	"github.com/ripx80/brave/log/memlog"
)

func main() {
	mem := memlog.New()
	log := *logger.New(mem)

	log.Info("test in mem")
	fmt.Println(mem.Entries)
}
