package util

import (
	"os"
	"os/signal"
)

func OsInterrupt() chan os.Signal {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	return stop
}
