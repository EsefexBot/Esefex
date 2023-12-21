package service

type Service interface {
	Start() <-chan struct{}
	Stop() <-chan struct{}
}
