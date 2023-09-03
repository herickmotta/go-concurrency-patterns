package semaphore

/*
	A semaphore pattern can be useful to limit the maximum concurrent go routines the program can run at once.
	Here we are using the idea of a buffered channel. Each time a goroutine is started, it must try to do an acquire,
	if the buffered channel is full, the goroutine must wait until a release is called.

	This can be used to limit the number of requests made to an external service or manage your database/container resources.
*/

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	semChan chan struct{}
}

func NewSemaphore(maxConcurrency int) Semaphore {
	return &semaphore{
		semChan: make(chan struct{}, maxConcurrency),
	}
}

func (s *semaphore) Acquire() {
	s.semChan <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.semChan
}
