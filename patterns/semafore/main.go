package main

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrTimeoutExceeded = errors.New("can't catch semafore")
	ErrIllegalRelease  = errors.New("can't release semafore")
)

type SemaforeInterface interface {
	Acquire() error
	Release() error
}

type semafore struct {
	sem     chan struct{}
	timeout time.Duration
}

func (s *semafore) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		// do some specific logic for every acquire (tick)
		return nil
	case <-time.After(s.timeout):
		return ErrTimeoutExceeded
	}
}

func (s *semafore) Release() error {
	select {
	case <-s.sem:
		// do some specific logic on every release (tick)
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}
}

func NewSemafore(tickets int, timeout time.Duration) SemaforeInterface {
	return &semafore{
		sem: make(chan struct{}, tickets),
		timeout: timeout,
	}
}

// TODO : разобраться, зачем это нужно?

func main() {
	tickets := 3

	start := time.Now()
	semafore := NewSemafore(tickets, 3 * time.Second)
	for i := 0; i < tickets; i++ {
		if err := semafore.Acquire(); err != nil {
			panic(err)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("all tickets acquired, time elapsed : %f sec\n", elapsed.Seconds())

	for i := 0; i < tickets; i++ {
		if err := semafore.Release(); err != nil {
			panic(err)
		}
	}

	elapsed = time.Since(start)
	fmt.Printf("all tickets released, time elapsed : %f sec\n", elapsed.Seconds())
	fmt.Println("it is fine")
}
