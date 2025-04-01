package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	tickets chan struct{}
}

func NewSemaphore(ticketsNumber int) Semaphore {
	return Semaphore{
		tickets: make(chan struct{}, ticketsNumber),
	}
}

func (s *Semaphore) Acquire() {
	s.tickets <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.tickets
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(6)

	sema := NewSemaphore(5)

	for i := 0; i < 6; i++ {
		sema.Acquire()

		go func() {
			defer func() {
				wg.Done()
				sema.Release()
			}()

			fmt.Println("working...")
			time.Sleep(2 * time.Second)
			fmt.Println("exiting...")
		}()
	}

	wg.Wait()
}
