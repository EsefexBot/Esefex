package main

import (
	"esefexapi/service"
	"time"
)

func main() {
	ts := service.NewExampleService(0 * time.Second)

	<-ts.Start()
	time.Sleep(1 * time.Second)
	<-ts.Stop()
}
