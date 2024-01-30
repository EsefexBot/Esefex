package util

import (
	"os"
	"os/signal"
	"syscall"
)

func Interrupt() chan os.Signal {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	return stop
}
