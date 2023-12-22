package service

import (
	"log"
	"time"
)

var _ IService = &ExampleService{}

// Service is the interface that all services must implement
type ExampleService struct {
	stop  chan struct{}
	ready chan struct{}
	delay time.Duration
}

func NewExampleService(d time.Duration) *ExampleService {
	return &ExampleService{
		stop:  make(chan struct{}),
		ready: make(chan struct{}),
		delay: d,
	}
}

func (s *ExampleService) run() {
	defer close(s.stop)

	log.Println("ExampleService starting...")
	time.Sleep(s.delay)
	log.Println("ExampleService started")

	close(s.ready)
	<-s.stop

	log.Println("ExampleService stopping...")
	time.Sleep(s.delay)
	log.Println("ExampleService stopped")
}

// Start implements Service.
func (s *ExampleService) Start() <-chan struct{} {
	go s.run()
	return s.ready
}

// Stop implements Service.
func (s *ExampleService) Stop() <-chan struct{} {
	s.stop <- struct{}{}
	return s.stop
}
