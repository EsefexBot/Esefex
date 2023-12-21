package main

import (
	"esefexapi/service"
	"time"
)

func main() {
	ts := service.NewExampleService()

	<-ts.Start()
	time.Sleep(5 * time.Second)
	<-ts.Stop()
}
