package main

import (
	"fmt"
	"github.com/ripx80/log/logger"
	"github.com/ripx80/log/memlog"
)

func main() {
	mem := memlog.New()
	log := *logger.New(mem)
	logger.Set(log)

	log.Infof("test in mem")
	fmt.Println(mem.Entries)
}
