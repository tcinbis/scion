package utils

import (
	"flag"
	"fmt"
	"github.com/scionproto/scion/go/lib/log"
	"os"
)

func SetupLogger() {
	logCfg := log.Config{Console: log.ConsoleConfig{Level: "debug"}}
	if err := log.Setup(logCfg); err != nil {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "Error configuring logger. Exiting due to:%s\n", err)
		os.Exit(-1)
	}
}
