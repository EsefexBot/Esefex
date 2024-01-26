package timer

import (
	"log"
	"sync"
	"time"
)

var rw sync.RWMutex
var start time.Time
var logEnabled bool

func SetStart() {
	rw.Lock()
	defer rw.Unlock()
	start = time.Now()
}

func Elapsed() time.Duration {
	rw.RLock()
	defer rw.RUnlock()
	return time.Since(start)
}

func PrintElapsed() {
	rw.RLock()
	defer rw.RUnlock()

	if logEnabled {
		log.Printf("Elapsed: %v\n", time.Since(start))
	}
}

func MessageElapsed(msg string) {
	rw.RLock()
	defer rw.RUnlock()

	if logEnabled {
		log.Printf("%s: %v\n", msg, time.Since(start))
	}
}

func EnableLog() {
	rw.Lock()
	defer rw.Unlock()
	logEnabled = true
}

func DisableLog() {
	rw.Lock()
	defer rw.Unlock()
	logEnabled = false
}
