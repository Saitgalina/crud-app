package service

type Authorization interface {
}

type Book interface {
}

type Sevice struct {
	Authorization
	Book
}

func NewService() *Sevice {
	return &Sevice{}
}
