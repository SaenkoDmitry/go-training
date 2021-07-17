package main

import (
	"fmt"
)

type semaphore chan struct{}

func NewSemaphore(n int) semaphore {
	return make(semaphore, n)
}

func (s semaphore) Incr(k int) {
	for i := 0; i < k; i++ {
		s <- struct{}{}
	}
}

func (s semaphore) Decr(k int) {
	for i := 0; i < k; i++ {
		<-s
	}
}

func main() {
	numbers := []int{1,2,3,4,5}
	n := len(numbers)

	sema := NewSemaphore(n)
	for _, num := range numbers {
		go func(n int) {
			fmt.Println(n)
			sema.Incr(1)
		}(num)
	}
	sema.Decr(n)
}
