package service

type IService interface {
	Start() <-chan struct{}
	Stop() <-chan struct{}
}
