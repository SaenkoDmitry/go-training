// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func splitChannels[T any](input chan T, n int) []<-chan T {
	outputCh := make([]chan T, n)
	for i := 0; i < n; i++ {
		outputCh[i] = make(chan T)
	}

	go func() {
		idx := 0
		for value := range input {
			outputCh[idx] <- value
			idx = (idx + 1) % n
		}

		for _, ch := range outputCh {
			close(ch)
		}
	}()

	// cannot case []chan T into []<-chan T
	resultCh := make([]<-chan T, n)
	for i := 0; i < n; i++ {
		resultCh[i] = outputCh[i]
	}
	return resultCh
}

func main() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()

	channels := splitChannels(channel, 2)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for value := range channels[0] {
			fmt.Println("ch1:", value)
		}
	}()

	go func() {
		defer wg.Done()
		for value := range channels[1] {
			fmt.Println("ch2:", value)
		}
	}()

	wg.Wait()
}
